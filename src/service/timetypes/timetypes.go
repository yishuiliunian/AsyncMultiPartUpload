package timetypes

import (
	"dzdatabase"
	"github.com/bitly/go-simplejson"
	"models"
)

func HandleUpdateTimeTypes(json *simplejson.Json) ([]byte, error) {
	timetype, err := models.NewDZTimeTypeFromJSON(json)
	if err != nil {
		return err
	}
	err = dzdatabase.UpdateDZTimeType(timetype)
	if err != nil {
		return err
	}
	return nil
}
