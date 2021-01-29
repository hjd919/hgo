package qq

import (
	"encoding/json"
	"fmt"

	"github.com/hjd919/gom"
)

var prefixHost = "https://apis.map.qq.com/ws"

type Handle struct {
	Key string
}

// 逆地址
func (t *Handle) GeocodeRegeo(lat, lng float64) (address string, err error) {
	url := prefixHost + fmt.Sprintf("/geocoder/v1/?location=%f,%f&key=%s", lat, lng, t.Key)
	buf, err := gom.HTTPGet(url)
	if err != nil {
		return
	}
	entity := GeocodeRegeo{}
	json.Unmarshal(buf, &entity)
	if entity.Status != 0 {
		err = fmt.Errorf(entity.Message)
		return
	}
	address = entity.Result.Address
	return
}
