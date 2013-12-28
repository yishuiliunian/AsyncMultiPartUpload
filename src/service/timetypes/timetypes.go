package timetypes

import (
	"dzdatabase"
	"github.com/bitly/go-simplejson"
	"models"
	"utilities"
)

func HandleUpdateTimeTypes(json *simplejson.Json, userGuid string) ([]byte, error) {
	timetype, err := models.NewDZTimeTypeFromJSON(json)
	if err != nil {
		return nil, err
	}
	err = dzdatabase.UpdateDZTimeType(timetype)
	if err != nil {
		return nil, err
	}
	return utilities.DZServerSucceedResponseData(), nil
}
