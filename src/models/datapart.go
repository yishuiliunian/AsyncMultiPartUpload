package models

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"utilities"
)

type DZDataPart struct {
	DZObject
	MD5       string
	SumMD5    string
	Data      []byte
	PartSN    int
	PartCount int
	FileType  string
}

func (d *DZDataPart) ToJSONObject() (*simplejson.Json, error) {
	json, err := d.DZObject.ToJSONObject()
	if err != nil {
		return json, err
	}
	json.Set(DZObjectKeyMD5, d.MD5)
	json.Set(DZObjectKeyData, d.Data)
	json.Set(DZObjectKeyPartCount, d.PartCount)
	json.Set(DZObjectKeyPartSN, d.PartSN)
	json.Set(DZObjectKeySumMD5, d.SumMD5)
	json.Set(DZObjectKeyFileType, d.FileType)
	return json, nil
}

func (d *DZDataPart) DecodeFromJSONOBject(json *simplejson.Json) error {
	var err error
	err = d.DZObject.DecodeFromJSONOBject(json)
	if err != nil {
		fmt.Println("decode super error")
		return err
	}
	d.MD5, err = json.Get(DZObjectKeyMD5).String()
	if err != nil {
		return utilities.NewError(-4, string("parse md5 error"))
	}
	d.Data, err = json.Get(DZObjectKeyData).Bytes()
	if err != nil {
		return utilities.NewError(-4, string("parse data error"))
	}
	d.PartSN, err = json.Get(DZObjectKeyPartSN).Int()
	if err != nil {
		return utilities.NewError(-4, string("parse part sn error"))
	}
	d.PartCount, err = json.Get(DZObjectKeyPartCount).Int()
	if err != nil {
		return utilities.NewError(-4, string("parse part count error"))
	}
	d.SumMD5, err = json.Get(DZObjectKeySumMD5).String()
	if err != nil {
		return utilities.NewError(-4, string("parse summd5 error"))
	}
	d.FileType, err = json.Get(DZObjectKeyFileType).String()
	if err != nil {
		return utilities.NewError(-4, string("parse filetype error"))
	}
	return err
}

func NewDataPart() *DZDataPart {
	dataPart := new(DZDataPart)
	return dataPart
}

func NewDataPartWithJson(json *simplejson.Json) (*DZDataPart, error) {
	dataPart := NewDataPart()
	err := dataPart.DecodeFromJSONOBject(json)
	if err != nil {
		return nil, err
	}
	return dataPart, nil
}
