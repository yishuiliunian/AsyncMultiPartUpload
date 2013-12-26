package service

import (
	"github.com/bitly/go-simplejson"
)

func ResponseSuccessJSON(message string) *simplejson.Json {
	json, _ := simplejson.NewJson([]byte("{}"))
	json.Set("code", "200")
	json.Set("message", "success")
	return json
}
