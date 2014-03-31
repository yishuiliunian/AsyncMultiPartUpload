package models

import (
	"fmt"
	"github.com/bitly/go-simplejson"
)

const (
	DZObjectKeyGUID        string = "Guid"
	DZObjectKeyVersion     string = "version"
	DZObjectKeyMD5         string = "md5"
	DZObjectBaseVersion    uint64 = 0
	DZObjectKeyData        string = "data"
	DZObjectKeyPartSN      string = "partsn"
	DZObjectKeyPartCount   string = "partcount"
	DZObjectKeySumMD5      string = "summd5"
	DZObjectKeyLocalUrl    string = "localurl"
	DZObjectKeyFileType    string = "filetype"
	DZObjectKeyName        string = "name"
	DZObjectKeyDetail      string = "detail"
	DZObjectKeyDateBegin   string = "dateBegin"
	DZObjectKeyDateEnd     string = "dateEnd"
	DZObjectKeyTypeGuid    string = "typeGuid"
	DZObjectKeyEmail       string = "email"
	DZObjectKeyPassword    string = "password"
	DZObjectKeyDeviceGUID  string = "deviceGuid"
	DZObjectKeyPhoneNumber string = "phonenumber"
	DZObjectKeyNickName    string = "nickname"
	DZObjectKeyAvatar      string = "avatar"
	DZObjectKeyDetailInfos string = "detailinfos"
	DZObjectKeyUserGuid    string = "userguid"
	DZObjectKeyOtherInfos  string = "otherinfos"
)

const (
	DZFileTypeImage string = "image"
)

type DZObject struct {
	Guid string
}

type DZJSONObject interface {
	ToJSONObject() (*simplejson.Json, error)
	DecodeFromJSONOBject(json *simplejson.Json) error
}

func (d *DZObject) ToJSONObject() (*simplejson.Json, error) {
	json, err := simplejson.NewJson([]byte("{}"))
	if err != nil {
		return json, err
	}
	json.Set(DZObjectKeyGUID, d.Guid)
	return json, nil
}

func (d *DZObject) DecodeFromJSONOBject(json *simplejson.Json) error {
	fmt.Println(json)
	str, err := json.Get(DZObjectKeyGUID).String()
	if err != nil {
		fmt.Println("decode guid error")
		return err
	}
	d.Guid = str
	fmt.Println("guid is %d", str)
	return nil
}
