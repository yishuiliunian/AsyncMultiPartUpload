package networks

import (
	"authorization"
	// "fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"utilities"
)

type DZRequstData struct {
	BodyJson   *simplejson.Json
	Token      string
	TokenVaild bool
	DeviceKey  string
	Method     string
	ClientType string
}

func DecodeHttpRequest(req *http.Request) (*DZRequstData, error) {
	body := req.Body
	defer body.Close()
	requstData := &DZRequstData{}
	bodystr, _ := ioutil.ReadAll(body)
	json, err := simplejson.NewJson(bodystr)
	if err != nil {
		return requstData, err
	}
	isVaild := false
	token, err := json.Get(DZProtocolKeyToken).String()
	if err != nil {
		isVaild = false
	} else {
		requstData.Token = token
		isVaild, err = authorization.CheckTokenIsVaild(token)
	}
	requstData.TokenVaild = isVaild
	//
	//
	method, err := json.Get(DZProtocolKeyMethod).String()
	if err != nil {
		return requstData, utilities.NewError(utilities.DZErrorCodePaser, "method paser error!")
	}
	requstData.Method = method
	//
	datas := json.Get(DZProtocolKeyDatas)
	if err != nil {
		return requstData, utilities.NewError(utilities.DZErrorCodePaser, "datas paser error!")
	}
	requstData.BodyJson = datas
	//
	clientType, err := json.Get(DZProtocolKeyCilentType).String()
	requstData.ClientType = clientType
	//
	deviecekey, err := json.Get(DZProtocolKeyDeviceKey).String()
	requstData.DeviceKey = deviecekey
	return requstData, err
}
