package models

import (
	"github.com/bitly/go-simplejson"
)

type DZDevice struct {
	DZObject
	UserGUID   string
	Name       string
	Detail     string
	OtherInfos string
}

func (d *DZDevice) DecodeFromJSONOBject(json *simplejson.Json) error {
	var err error
	err = d.DZObject.DecodeFromJSONOBject(json)
	if err != nil {
		return err
	}
	d.UserGUID, err = json.Get(DZObjectKeyUserGuid).String()
	d.Name, err = json.Get(DZObjectKeyName).String()
	d.Detail, err = json.Get(DZObjectKeyAvatar).String()
	d.OtherInfos, err = json.Get(DZObjectKeyOtherInfos).String()
	return nil
}

func NewDeviceWithJson(json *simplejson.Json) (*DZDevice, error) {
	dv := &DZDevice{}
	err := dv.DecodeFromJSONOBject(json)
	return dv, err
}
