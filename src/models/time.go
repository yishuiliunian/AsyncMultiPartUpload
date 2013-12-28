package models

import (
	"github.com/bitly/go-simplejson"
	"time"
	"utilities"
)

const (
	JDK_TimeName       = "Name"
	JDK_TimeDateBegin  = "DateBegin"
	JDK_TimeDateEnd    = "DateEnd"
	JDK_TimeDetail     = "Detail"
	JDK_TimeDeviceGUID = "DeviceGUID"
	JDK_TimeGuid       = "Guid"
	JDK_TimeTypeGUID   = "TypeGUID"
	JDK_TimeUserGUID   = "UserGUID"
	JDK_TimeVersion    = DZObjectKeyVersion
)

type DZTime struct {
	DZObject
	DateBegin  time.Time
	DateEnd    time.Time
	TypeGUID   string
	Detail     string
	Version    int64
	DeviceGUID string
	UserGUID   string
}

func (d *DZTime) ToJSONObject() (*simplejson.Json, error) {
	json, err := d.DZObject.ToJSONObject()
	if err != nil {
		return json, err
	}
	json.Set(JDK_TimeDateBegin, d.DateBegin)
	json.Set(JDK_TimeDateEnd, d.DateEnd)
	json.Set(JDK_TimeVersion, d.Version)
	json.Set(JDK_TimeDetail, d.Detail)
	json.Set(JDK_TimeTypeGUID, d.TypeGUID)
	json.Set(JDK_TimeDeviceGUID, d.DeviceGUID)
	return json, nil
}

func (d *DZTime) DecodeFromJSONOBject(json *simplejson.Json) error {
	var err error
	err = d.DZObject.DecodeFromJSONOBject(json)
	if err != nil {
		return err
	}
	d.UserGUID, err = json.Get(JDK_TimeUserGUID).String()
	if err != nil {
		return utilities.NewError(utilities.DZErrorCodePaser, "parser user guid error")
	}
	datebegin, err := json.Get(JDK_TimeDateBegin).String()
	if err != nil {
		return utilities.NewError(utilities.DZErrorCodePaser, "parser date begin error")
	}
	d.DateBegin, err = utilities.ParseTimeString(datebegin)
	if err != nil {
		return utilities.NewError(utilities.DZErrorCodePaser, "date time format is invaild")
	}
	//
	dateend, err := json.Get(JDK_TimeDateEnd).String()
	if err != nil {
		return utilities.NewError(utilities.DZErrorCodePaser, "parser date end error")
	}

	d.DateEnd, err = utilities.ParseTimeString(dateend)
	if err != nil {
		return utilities.NewError(utilities.DZErrorCodePaser, "end date format is invaild")
	}

	d.TypeGUID, err = json.Get(JDK_TimeTypeGUID).String()
	if err != nil {
		return utilities.NewError(utilities.DZErrorCodePaser, "parser type guid  error")
	}

	d.Detail, err = json.Get(JDK_TimeDetail).String()

	d.Version, err = json.Get(JDK_TimeVersion).Int64()
	d.DeviceGUID, err = json.Get(JDK_TimeDeviceGUID).String()
	if err != nil {
		return utilities.NewError(utilities.DZErrorCodePaser, "parse DeviceGUID error")
	}
	return nil
}

func NewDZTimeFromJSON(json *simplejson.Json) (*DZTime, error) {
	dt := &DZTime{}
	err := dt.DecodeFromJSONOBject(json)
	return dt, err
}
