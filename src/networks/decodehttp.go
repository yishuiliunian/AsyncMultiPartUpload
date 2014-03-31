package networks

import (
	"service/authorization"
	// "fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"utilities"
)

type DZRequstData struct {
	BodyJson   *simplejson.Json
	Token      string
	DeviceKey  string
	Method     string
	ClientType string
	AppKey     string
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
	requstData.Token, _ = json.Get(DZProtocolKeyToken).String()
	//
	requstData.AppKey, err = json.Get(DZProtocolKeyAppkey).String()
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

func CheckRequestDataAcessVaild(reqData *DZRequstData) (bool, string, error) {
	token := reqData.Token
	if token == "" {
		return false, "", utilities.NewError(utilities.DZErrorCodePaser, "token is nil")
	}
	deviceKey := reqData.DeviceKey
	if deviceKey == "" {
		return false, "", utilities.NewError(utilities.DZErrorCodePaser, "device key is nil")
	}
	vaild, userguid, err := authorization.CheckTokenIsVaild(token, deviceKey)
	if err != nil || !vaild {
		return false, "", utilities.NewError(utilities.DZErrorCodeTokenInvaild, "token is invaild")
	}
	return true, userguid, nil
}
