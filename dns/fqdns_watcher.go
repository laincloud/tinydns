package dns

import (
	"time"

	"encoding/json"

	"github.com/mijia/sweb/log"
	"golang.org/x/net/context"
)

// FqdnsWatcher
type FqdnsWatcher struct {
	notify chan interface{}
}

func NewFqdnsWatcher(notify chan interface{}) *FqdnsWatcher {
	return &FqdnsWatcher{
		notify: notify,
	}
}

func (w *FqdnsWatcher) Watch() {
	log.Infof("FqdnsWatcher thread started")
	ctx, _ := context.WithTimeout(context.Background(), lainletTimeout)
	for {
		if ch, err := lainletClient.Watch("/v2/configwatcher?target=tinydns_fqdns", ctx); err != nil {
			log.Errorf("FqdnsWatcher connect to lainlet failed. Retry in 3 seconds")
		} else {
			for resp := range ch {
				if resp.Event == "init" || resp.Event == "update" || resp.Event == "delete" {
					newData := make(Fqdns)
					if err := json.Unmarshal(resp.Data, &newData); err != nil {
						log.Errorf("FqdnsWatcher unmarshall data failed: %s", err.Error())
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
