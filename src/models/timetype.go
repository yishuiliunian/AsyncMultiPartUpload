package models

import (
	"github.com/bitly/go-simplejson"
	"utilities"
)

const (
	JSONKeyTypeGuid       = "guid"
	JSONKeyTypeName       = "name"
	JSONKeyTypeDetail     = "detail"
	JSONKeyTypeOtherInfos = "other_infos"
	JSONKeyTypeUserGuid   = "user_guid"
)

type DZTimeType struct {
	DZObject
	UserGuid   string
	Name       string
	Detail     string
	OtherInfos string
	Version    int64
}

func (d *DZTimeType) DECodeFromJSONObject(json *simplejson.Json) error {
	var err error
	err = d.DZObject.DecodeFromJSONOBject(json)
	if err != nil {
		return err
	}
	d.UserGuid, err = json.Get(JSONKeyTypeUserGuid).String()
	if err != nil {
		return utilities.NewError(utilities.DZErrorCodePaser, "parse user guid error")
	}
	d.Name, err = json.Get(JSONKeyTypeName).String()
	if err != nil {
		return utilities.NewError(utilities.DZErrorCodePaser, "parse name error")
	}
	d.Detail, err = json.Get(JSONKeyTypeDetail).String()
	if err != nil {
		return utilities.NewError(utilities.DZErrorCodePaser, "parse detail error")
	}
	d.OtherInfos, err = json.Get(JSONKeyTypeOtherInfos).String()
	if err != nil {
		return utilities.NewError(utilities.DZErrorCodePaser, "parse otherinfo error")
	}
	return nil
}

func NewDZTimeTypeFromJSON(json *simplejson.Json) (*DZTimeType, error) {
	tp := &DZTimeType{}
	err := tp.DECodeFromJSONObject(json)
	return tp, err
}
