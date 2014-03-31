package models

import (
	"github.com/bitly/go-simplejson"
	"utilities"
)

const (
	DZAppJKGuid         = "guid"
	DZAppJKOwnerGuid    = "owner_guid"
	DZAppJKName         = "name"
	DZAppJKReferenceURL = "reference_url"
	DZAppJKDescription  = "description"
	DZAppJKVersion      = "version"
	DZAppJKDetail       = "detail"
	DZAppJKType         = "type"
	DZAppJKSubType      = "sub_type"
	DZAppJKPlatfomat    = "platfomat"
	DZAppJKTags         = "tags"
)

type DZApp struct {
	DZObject
	OwnerGuid    string
	Name         string
	ReferenceURL string
	Description  string
	Version      string
	Detail       string
	Type         string
	SubType      string
	Platfomat    string
	Tags         string
	IsVaild      bool
}

func (d *DZApp) DecodeFromJSONObject(json *simplejson.Json) error {
	err := d.DZObject.DecodeFromJSONOBject(json)
	if err != nil {
		return err
	}
	d.Name, err = json.Get(DZAppJKName).String()
	if err != nil {
		return utilities.NewError(utilities.DZErrorCodePaser, "name guid error")
	}

	d.Type, err = json.Get(DZAppJKName).String()
	if err != nil {
		return utilities.NewError(utilities.DZErrorCodePaser, "type guid error")
	}

	d.SubType, err = json.Get(DZAppJKSubType).String()
	if err != nil {
		return utilities.NewError(utilities.DZErrorCodePaser, "subtype guid error")
	}

	d.OwnerGuid, err = json.Get(DZAppJKOwnerGuid).String()
	if err != nil {
		return utilities.NewError(utilities.DZErrorCodePaser, "owner guid error")
	}

	d.Version, _ = json.Get(DZAppJKVersion).String()
	d.ReferenceURL, _ = json.Get(DZAppJKReferenceURL).String()
	d.Platfomat, _ = json.Get(DZAppJKPlatfomat).String()
	d.Tags, _ = json.Get(DZAppJKTags).String()
	d.Description, _ = json.Get(DZAppJKDescription).String()
	d.Detail, _ = json.Get(DZAppJKDetail).String()
	return nil
}

func NewAppFromJSON(json *simplejson.Json) (*DZApp, error) {
	app := &DZApp{}
	err := app.DecodeFromJSONOBject(json)
	return app, err
}
