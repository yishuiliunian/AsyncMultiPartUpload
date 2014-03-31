package models

import (
	"github.com/bitly/go-simplejson"
	"utilities"
)

const (
	DZUserVIPNone = "free"
	DZUserVIP1    = "vip1"
)

type DZUser struct {
	DZObject
	Name        string
	Email       string
	Password    string
	PhoneNumber string
	NickName    string
	Avatar      string
	DetailInfos string
	VIPType     string
}

func (d *DZUser) DecodeFromJSONOBject(json *simplejson.Json) error {
	var err error
	d.DZObject.DecodeFromJSONOBject(json)
	d.Email, err = json.Get(DZObjectKeyEmail).String()
	if err != nil {
		return utilities.NewError(utilities.DZErrorCodePaser, "parse email error")
	}
	d.Password, err = json.Get(DZObjectKeyPassword).String()
	if err != nil {
		return utilities.NewError(utilities.DZErrorCodePaser, "parse password error")
	}
	d.PhoneNumber, _ = json.Get(DZObjectKeyPhoneNumber).String()
	d.NickName, _ = json.Get(DZObjectKeyNickName).String()
	d.Avatar, _ = json.Get(DZObjectKeyAvatar).String()
	d.DetailInfos, _ = json.Get(DZObjectKeyDetailInfos).String()
	device := json.Get("device")
	if device != nil {
		de := &DZDevice{}
		de.DecodeFromJSONOBject(device)
	}
	return nil
}

func NewDZUserFromJson(json *simplejson.Json) (*DZUser, error) {
	duser := &DZUser{}
	err := duser.DecodeFromJSONOBject(json)
	return duser, err
}
