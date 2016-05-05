package dns

import (
	"time"

	"encoding/json"

	"github.com/mijia/sweb/log"
	"golang.org/x/net/context"
)

type CoreInfoWatcher struct {
	notify chan interface{}
}

func NewCoreInfoWatcher(notify chan interface{}) *CoreInfoWatcher {
	return &CoreInfoWatcher{
		notify: notify,
	}
}

func (w *CoreInfoWatcher) Watch() {
	log.Infof("CoreInfoWatcher thread started")
	for {
		if ch, err := lainletClient.Watch("/v2/coreinfowatcher", context.Background()); err != nil {
			log.Errorf("CoreInfoWatcher connect to lainlet failed. Retry in 3 seconds")
		} else {
			for resp := range ch {
				if resp.Event == "init" || resp.Event == "update" || resp.Event == "delete" {
					newData := make(CoreInfo)
					if err := json.Unmarshal(resp.Data, &newData); err != nil {
						log.Errorf("CoreInfoWatcher unmarshall data failed: %s", err.Error())
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
