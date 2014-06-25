package models

import (
	"github.com/bitly/go-simplejson"
	"time"
	"utilities"
	"fmt"
)

const (
	JSONKeyTypeGuid       = "Guid"
	JSONKeyTypeName       = "Name"
	JSONKeyTypeDetail     = "Detail"
	JSONKeyTypeOtherInfos = "OtherInfos"
	JSONKeyTypeUserGuid   = "UserGuid"
	JSONKeyTypeIsFinished = "Finished"
	JSONKeyTypeCreateDate = "CreateDate"
)

type DZTimeType struct {
	DZObject
	UserGuid   string
	Name       string
	Detail     string
	OtherInfos string
	CreateDate time.Time
	Finished   bool
	Version    int64
}

func (d *DZTimeType) DECodeFromJSONObject(json *simplejson.Json) error {
	var err error
	fmt.Println(json)
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
	cdate, err := json.Get(JSONKeyTypeCreateDate).String()
	d.CreateDate, err = utilities.ParseTimeString(cdate)
	if err != nil {
		return utilities.NewError(utilities.DZErrorCodePaser, "parse create date error")
	}
	d.Finished, err = json.Get(JSONKeyTypeIsFinished).Bool()
	if err != nil {
		return utilities.NewError(utilities.DZErrorCodePaser, "parse finished error")
	}
	return nil
}

func NewDZTimeTypeFromJSON(json *simplejson.Json) (*DZTimeType, error) {
	tp := &DZTimeType{}
	err := tp.DECodeFromJSONObject(json)
	return tp, err
}
