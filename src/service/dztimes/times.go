package dztimes

import (
	"dzdatabase"
	"fmt"
	"github.com/bitly/go-simplejson"
	"models"
	"utilities"
)

const (
	kGetTimeStartVersion = "start_version"
	kGetTimeRequestCount = "request_cout"
)

func HandleGetTimesRequest(json *simplejson.Json, userGuid string) ([]byte, error) {
	startV, err := json.Get(kGetTimeStartVersion).Int64()
	if err != nil {
		return nil, utilities.NewError(utilities.DZErrorCodePaser, "parser start version error")
	}
	count, err := json.Get(kGetTimeRequestCount).Int64()
	if err != nil {
		return nil, utilities.NewError(utilities.DZErrorCodePaser, "parser requst count")
	}
	times, err := dzdatabase.GetTimesOfUserWithVersionSpace(userGuid, startV, startV+count)
	if err != nil {
		return nil, err
	}
	rj, err := simplejson.NewJson([]byte("{}"))
	if err != nil {
		return nil, utilities.NewError(utilities.DZErrorCodePaser, "parser json error")
	}
	rj.Set("times", times)
	return rj.MarshalJSON()
}

func HandleUpdateTime(dt *simplejson.Json, userGuid string) ([]byte, error) {
	dtdata, err := models.NewDZTimeFromJSON(dt)
	if err != nil {
		return nil, err
	}
	if dtdata == nil {
		fmt.Println("nil")
	}
	err = dzdatabase.UpdateDZTime(dtdata)
	if err != nil {
		return nil, err
	}
	return utilities.DZServerSucceedResponseData(), nil
}
