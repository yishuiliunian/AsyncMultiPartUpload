package models

import (
	"github.com/bitly/go-simplejson"
	"time"
	"utilities"
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
	json.Set(DZObjectKeyDateBegin, d.DateBegin)
	json.Set(DZObjectKeyDateEnd, d.DateEnd)
	json.Set(DZObjectKeyVersion, d.Version)
	json.Set(DZObjectKeyDetail, d.Detail)
	json.Set(DZObjectKeyTypeGuid, d.TypeGUID)
	json.Set(DZObjectKeyDeviceGUID, d.DeviceGUID)
	return json, nil
}

func (d *DZTime) DecodeFromJSONOBject(json *simplejson.Json) error {
	var err error
	err = d.DZObject.DecodeFromJSONOBject(json)
	if err != nil {
		return err
	}
	d.UserGUID, err = json.Get("userGuid").String()
	if err != nil {
		return utilities.NewError(utilities.DZErrorCodePaser, "parser user guid error")
	}
	datebegin, err := json.Get(DZObjectKeyDateBegin).String()
	if err != nil {
		return utilities.NewError(utilities.DZErrorCodePaser, "parser date begin error")
	}
	d.DateBegin, err = utilities.ParseTimeString(datebegin)
	if err != nil {
		return utilities.NewError(utilities.DZErrorCodePaser, "date time format is invaild")
	}
	//
	dateend, err := json.Get(DZObjectKeyDateEnd).String()
	if err != nil {
		return utilities.NewError(utilities.DZErrorCodePaser, "parser date end error")
	}

	d.DateEnd, err = utilities.ParseTimeString(dateend)
	if err != nil {
		return utilities.NewError(utilities.DZErrorCodePaser, "end date format is invaild")
	}

	d.TypeGUID, err = json.Get(DZObjectKeyTypeGuid).String()
	if err != nil {
		return utilities.NewError(utilities.DZErrorCodePaser, "parser type guid  error")
	}

	d.Detail, err = json.Get(DZObjectKeyDetail).String()

	d.Version, err = json.Get(DZObjectKeyVersion).Int64()
	d.DeviceGUID, err = json.Get(DZObjectKeyDeviceGUID).String()
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
