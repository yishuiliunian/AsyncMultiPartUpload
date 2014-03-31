package dztimes

import (
	"dzdatabase"
	"fmt"
	"github.com/bitly/go-simplejson"
	"models"
	"restfulbase"
	"utilities"
)

func HandleGetTimesRequest(json *simplejson.Json, userGuid string) ([]byte, error) {
	startV, err := json.Get(restfulbase.KGetStartVersion).Int64()
	if err != nil {
		return nil, utilities.NewError(utilities.DZErrorCodePaser, "parser start version error")
	}
	count, err := json.Get(restfulbase.KGeRequestCount).Int64()
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
	rj.Set("objects", times)
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
