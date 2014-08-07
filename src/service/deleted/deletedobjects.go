package deleted

import (
	"dzdatabase"
	"fmt"
	"github.com/bitly/go-simplejson"
	"models"
	"restfulbase"
	"utilities"
	"errors"
)

func HandleGetDeletedObjectsRequest(json* simplejson.Json, userGuid string) ([]byte, error) {
	startV, err := json.Get(restfulbase.KGetStartVersion).Int64()
	if err != nil {
		return nil, utilities.NewError(utilities.DZErrorCodePaser, "parser start version error")
	}
	count, err := json.Get(restfulbase.KGeRequestCount).Int64()
	if err != nil {
		return nil, utilities.NewError(utilities.DZErrorCodePaser, "parser requst count")
	}

	//
	deletedObjects , err :=  dzdatabase.GetDeletedObjectsOfUserWithVersionSpace(userGuid, startV, startV+count)
	if err != nil {
		return nil, err
	}

	rj, err := simplejson.NewJson([]byte("{}"))
		if err != nil {
		return nil, utilities.NewError(utilities.DZErrorCodePaser, "parser json error")
	}
	rj.Set("objects", deletedObjects)
	return rj.MarshalJSON()
}


func HandleUpdateDeletedObject(dt *simplejson.Json, userGuid string) ([]byte, error) {
	dtdata, err := models.NewDZDeleteObjectFromJSON(dt)
	if err != nil {
		return nil, err
	}
	if dtdata == nil {
		fmt.Println("nil")
	}
	err = dzdatabase.UpdateDZDeletedObject(dtdata)
	if err != nil {
		return nil, err
	}

	switch dtdata.Type {
		case models.JDK_DeletedType: {
			err = dzdatabase.RemoveDZTimeType(dtdata.Guid)
		}
		case models.JDK_DeletedTime : {
			err = dzdatabase.RemoveDZTime(dtdata.Guid)
		}
		default: {
			err =  errors.New("type error")
		}
	}
	if err != nil {
		return nil , err
	}
	return utilities.DZServerSucceedResponseData(), nil
}