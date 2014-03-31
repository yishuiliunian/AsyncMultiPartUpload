package utilities

import (
	"github.com/bitly/go-simplejson"
)

var _succedData []byte = nil

func DZServerSucceedResponseData() []byte {
	if _succedData == nil {
		json, _ := simplejson.NewJson([]byte("{}"))
		json.Set("code", 200)
		json.Set("message", "ok")
		_succedData, _ = json.MarshalJSON()
	}
	return _succedData
}
