package timetypes

import (
	"dzdatabase"
	"fmt"
	"github.com/bitly/go-simplejson"
	"models"
	"restfulbase"
	"utilities"
)

func HandleUpdateTimeTypes(json *simplejson.Json, userGuid string) ([]byte, error) {
	timetype, err := models.NewDZTimeTypeFromJSON(json)
	if err != nil {
		return nil, err
	}
	fmt.Println("update types")
	err = dzdatabase.UpdateDZTimeType(timetype)
	if err != nil {
		return nil, err
	}
	return utilities.DZServerSucceedResponseData(), nil
}

func HandleGetTimeTypes(json *simplejson.Json, userGuid string) ([]byte, error) {
	startV, err := json.Get(restfulbase.KGetStartVersion).Int64()
	if err != nil {
		return nil, utilities.NewError(utilities.DZErrorCodePaser, "parser start version error")
	}
	count, err := json.Get(restfulbase.KGeRequestCount).Int64()
	if err != nil {
		return nil, utilities.NewError(utilities.DZErrorCodePaser, "parser requst count")
	}
	times, err := dzdatabase.GetTimeTypesOfUserWithVersionSpace(userGuid, startV, startV+count)
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
