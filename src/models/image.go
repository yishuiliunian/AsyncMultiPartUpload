package models

import (
	"github.com/bitly/go-simplejson"
)

type DZImage struct {
	Guid     string
	Version  uint64
	Md5      string
	LocalUrl string
}

func (d *DZImage) ToJSONObject() (*simplejson.Json, error) {
	json, err := simplejson.NewJson([]byte("{}"))
	if err != nil {
		return json, err
	}
	json.Set(DZObjectKeyGUID, d.Guid)
	json.Set(DZObjectKeyMD5, d.Md5)
	json.Set(DZObjectKeyVersion, d.Version)
	json.Set(DZObjectKeyLocalUrl, d.LocalUrl)
	return json, nil
}

func NewImage() *DZImage {
	return &DZImage{}
}
