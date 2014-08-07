package models

import (
	"github.com/bitly/go-simplejson"
	"time"
	"utilities"
)

const (
	JDK_DeletedType = "Type"
	JDK_DeletedTime = "Time"
	JDK_DeletedVersion = "Version"
	JDK_DeletedUserGuid = "UserGUID"
)

type DZDeleteObject struct {
	DZObject
	Type string
	DeletedTime time.Time
	Version int64
	UserGUID string
}

func (d *DZDeleteObject) ToJSONObject() (*simplejson.Json, error) {
	json, err := d.DZObject.ToJSONObject()
	if err != nil {
		return json, err
	}

	json.Set(JDK_DeletedType, d.Type)
	json.Set(JDK_DeletedTime, d.DeletedTime)
	json.Set(JDK_DeletedUserGuid, d.UserGUID)
	return json, nil
}


func (d *DZDeleteObject) DecodeFromJSONOBject(json* simplejson.Json) error {
	var err error
	err = d.DZObject.DecodeFromJSONOBject(json)
	if err != nil {
		return err
	}

	d.Type, err = json.Get(JDK_DeletedType).String()
	if err != nil {
		return utilities.NewError(utilities.DZErrorCodePaser, "parser type error")
	}
	//
	d.UserGUID, err = json.Get(JDK_DeletedUserGuid).String()
	if err != nil {
		return utilities.NewError(utilities.DZErrorCodePaser, "parser user guid error")
	}
	date, err := json.Get(JDK_DeletedTime).String()
	if err != nil {
		return utilities.NewError(utilities.DZErrorCodePaser, "parser time error")
	}
	d.DeletedTime, err = utilities.ParseTimeString(date)
	if err != nil {
		return utilities.NewError(utilities.DZErrorCodePaser, "end date format is invaild")
	}

	return nil
}

func NewDZDeleteObjectFromJSON(json* simplejson.Json) (*DZDeleteObject, error) {
	dt := &DZDeleteObject{}
	err := dt.DecodeFromJSONOBject(json)
	return dt, err
}