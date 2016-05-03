package dns

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/mijia/sweb/log"
)

// Creator creates dns data file and calls reloading of tinydns
type Creator struct {
	notify   chan interface{}
	dataDir  string
	filePath string
	depends  Depends
	coreInfo CoreInfo
	fqdns    Fqdns
}

func NewCreator(notify chan interface{}) *Creator {
	return &Creator{
		notify:   notify,
		dataDir:  dataDir,
		filePath: filepath.Join(dataDir, dataFile),
		depends:  make(Depends),
		coreInfo: make(CoreInfo),
		fqdns:    make(Fqdns),
	}
}

func (c *Creator) CreateDNS() {
	var needUpdate bool
	var err error
	for data := range c.notify {
		needUpdate = false
		switch data.(type) {
		case Depends:
			newDepends := data.(Depends)
			if !reflect.DeepEqual(c.depends, newDepends) {
				c.depends = newDepends
				needUpdate = true
			}
		case CoreInfo:
			newCoreInfo := data.(CoreInfo)
			if !reflect.DeepEqual(c.coreInfo, newCoreInfo) {
				c.coreInfo = newCoreInfo
				needUpdate = true
			}
		case Fqdns:
			newFqdns := data.(Fqdns)
			if !reflect.DeepEqual(c.coreInfo, newFqdns) {
				c.fqdns = newFqdns
				needUpdate = true
			}
		default:
			log.Error("Creator get notify event error")
		}
		if needUpdate {
			// We must ensure dns-data is updated successfully
			for {
				if err = c.createDNS(); err == nil {
					break
				}
				log.Errorf("Create dns-data failed: %s", err.Error())
				time.Sleep(1 * time.Second)
			}
		}
		time.Sleep(2 * time.Second)
	}
}

// createDNS returns true if DNS data file is created and reloaded successfully
func (c *Creator) createDNS() error {
	var (
		err error
		fd  *os.File
	)
	if fd, err = os.Create(c.filePath); err != nil {
		return err
	}
	defer fd.Close()

	if _, err := fd.WriteString(fmt.Sprintln(strings.Join(c.prepareData(), "\n"))); err != nil {
		return err
	}
	return c.reloadFile()
}

func (c *Creator) prepareData() []string {
	var lines []string
	apps := make(map[string]int)
	// Write appInfos
	for key, val := range c.coreInfo {
		keyParts := strings.Split(key, ".")

		appName, procName := getAppProcName(keyParts)
		if appName == "" {
			continue
		}
		if appName != "webrouter" {
			apps[appName] = 1
		}
		for _, podInfo := range val.PodInfos {
			if podInfo.InstanceNo > 0 {
				lines = append(lines,
					fmt.Sprintf("=%s-%d.%s.lain:%s:300",
						procName,
						podInfo.InstanceNo,
						appName,
						podInfo.Containers[0].ContainerIP))
			}
		}
	}

	// Write portals
	var annotationInfo AnnotationInfo
	for portalName, portals := range c.depends {
		for nodeName, portalInNodes := range portals {
			for appName, portalInfo := range portalInNodes {
				if err := json.Unmarshal([]byte(portalInfo.Annotation), &annotationInfo); err != nil {
					log.Errorf("Unmarshal annotation of portal %s in node %s failed: %s\n", portalName, nodeName, err.Error())
					continue
				}
				if len(portalInfo.Containers) > 0 {
					lines = append(lines,
						fmt.Sprintf("=%s.%s.lain:%s:300::%s",
							annotationInfo.ServiceName,
							appName,
							portalInfo.Containers[0].ContainerIP,
							portalInfo.Containers[0].NodeIP))
				}
			}
		}
	}

	var webrouterIp string
	for key, val := range c.fqdns {
		var domains []string
		if err := json.Unmarshal([]byte(val), &domains); err != nil {
			log.Errorf("Unmarshal %s failed: %s\n", key, err.Error())
			continue
		}
		for _, line := range domains {
			if key == "tinydns_fqdns/webrouter.lain" {
				if strings.HasPrefix(line, "+webrouter.lain:") {
					fields := strings.Split(line, ":")
					if len(fields) >= 1 {
						webrouterIp = fields[1]
					} else {
						log.Errorf("Cannot get webrouter ip: %s\n", line)
					}
				}
			}
			lines = append(lines, line)
		}
	}

	if webrouterIp != "" {
		for app, _ := range apps {
			lines = append(lines, fmt.Sprintf("+%s.lain:%s:300::", app, webrouterIp))
		}
	}
	return lines
}

func (c *Creator) reloadFile() error {
	log.Debug("Rebuild data file")
	var err error
	if err = os.Chdir(c.dataDir); err == nil {
		err = exec.Command("tinydns-data").Run()
	}
	return err
}

// getAppProcName returns appName and procName
func getAppProcName(key []string) (string, string) {
	var procName string
	if len(key) > 0 {
		procName = key[len(key)-1]
	}
	var tmp []string
	for i := len(key) - 3; i >= 0; i-- {
		tmp = append(tmp, key[i])
	}
	return strings.Join(tmp, "."), procName
}
