package v2

import (
	"encoding/json"
	"fmt"
	"github.com/laincloud/lainlet/api"
	"github.com/laincloud/lainlet/auth"
	"github.com/laincloud/lainlet/watcher"
	"net/http"
	"reflect"
)

// Config API
type GeneralConfig struct {
	Data map[string]string // data type return by configwatcher
}

func (gc *GeneralConfig) Decode(r []byte) error {
	return json.Unmarshal(r, &gc.Data)
}

func (gc *GeneralConfig) Encode() ([]byte, error) {
	return json.Marshal(gc.Data)
}

func (gc *GeneralConfig) URI() string {
	return "/configwatcher"
}

func (gc *GeneralConfig) WatcherName() string {
	return watcher.CONFIG
}

func (gc *GeneralConfig) Make(conf map[string]interface{}) (api.API, bool, error) {
	ret := &GeneralConfig{
		Data: make(map[string]string),
	}
	for k, v := range conf {
		ret.Data[k], _ = v.(string)
	}

	return ret, !reflect.DeepEqual(gc.Data, ret.Data), nil
}

func (gc *GeneralConfig) Key(r *http.Request) (string, error) {
	if !auth.IsSuper(r.RemoteAddr) {
		return "", fmt.Errorf("authorize failed, super required")
	}
	target := api.GetString(r, "target", "*")
	return target, nil
}
