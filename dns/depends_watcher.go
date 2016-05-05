package dns

import (
	"time"

	"encoding/json"

	"github.com/mijia/sweb/log"
	"golang.org/x/net/context"
)

//DependsWatcher has data which is the map as follows: (portal DEPLOYD_POD_NAME)=>(node HOSTNAME=>(APP name using service=>PortalInfo))
type DependsWatcher struct {
	notify chan interface{}
}

func NewDependsWatcher(notify chan interface{}) *DependsWatcher {
	return &DependsWatcher{
		notify: notify,
	}
}

func (w *DependsWatcher) Watch() {
	log.Infof("DependsWatcher thread started")
	for {
		if ch, err := lainletClient.Watch("/v2/depends", context.Background()); err != nil {
			log.Errorf("DependsWatcher connect to lainlet failed. Retry in 3 seconds")
		} else {
			for resp := range ch {
				if resp.Event == "init" || resp.Event == "update" || resp.Event == "delete" {
					newData := make(Depends)
					if err := json.Unmarshal(resp.Data, &newData); err != nil {
						log.Errorf("DependsWatcher unmarshall data failed: %s", err.Error())
					} else {
						w.notify <- newData
					}
				}
				time.Sleep(3 * time.Second)
			}
		}
		time.Sleep(3 * time.Second)
	}
}
