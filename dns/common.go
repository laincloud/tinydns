package dns

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/laincloud/lainlet/client"
)

const (
	dataDir        = "/etc/ndjbdns"
	dataFile       = "data"
	lainletTimeout = time.Second * 3
)

var lainletClient = client.New(fmt.Sprintf("lainlet.lain:%s", getEnvWithDefault("LAINLET_PORT", "9001")))

type Depends map[string]map[string]map[string]PortalInfo
type CoreInfo map[string]AppInfo
type Fqdns map[string]string

type PortalInfo struct {
	Annotation string            `json:"Annotation"`
	Containers []PortalContainer `json:"Containers"`
}

type AnnotationInfo struct {
	ServiceName string `json:"service_name"`
}

type PortalContainer struct {
	NodeIP      string `json:"NodeIP"`
	ContainerIP string `json:"IP"`
}

type Container struct {
	NodeIP      string `json:"NodeIp"`
	ContainerIP string `json:"ContainerIp"`
}

type AppInfo struct {
	PodInfos []PodInfo `json:"PodInfos"`
}

type PodInfo struct {
	InstanceNo int         `json:"InstanceNo"`
	Containers []Container `json:"ContainerInfos"`
}

func getLainConfig(key string) (string, error) {
	result := make(map[string]string)
	var err error
	var data []byte
	if data, err = lainletClient.Get("/v2/configwatcher?target="+key, lainletTimeout); err == nil {
		err = json.Unmarshal(data, &result)
	}
	return result[key], err
}

func getEnvWithDefault(key, defaultVal string) string {
	var val string
	if val = os.Getenv(key); val == "" {
		val = defaultVal
	}
	return val
}
