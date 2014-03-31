package authorization

import (
	"dzdatabase"
	"github.com/bitly/go-simplejson"
	"strings"
	"utilities"
)

const (
	KeyToken = "token"
)

const (
	DZAppErrorCodeNotVaild = -5100
)

const (
	_ObjectKeyIdentify string = "identify"
	_DZSplitGUIDKey    string = "|"
)

func joinUserGUIDAndDeviceGUID(ug string, dg string) string {
	s := []string{ug, dg}
	return strings.Join(s, _DZSplitGUIDKey)
}

func splitGetUserGUIDAndDeviceGuid(key string) (string, string) {
	s := strings.Split(key, _DZSplitGUIDKey)
	if len(s) != 2 {
		return "", ""
	}
	return s[0], s[1]
}

func ApplyAnVaildToken(userguid string, deviceGuid string) (string, error) {
	token := utilities.GUID()
	value := joinUserGUIDAndDeviceGUID(userguid, deviceGuid)
	err := dzdatabase.AddExpireKeyValueToReids(token, value)
	return token, err
}

func UpdateTokenExpireTime(token string) error {
	return dzdatabase.UpdateExpireByKey(token)
}

func CheckTokenIsVaild(token string, deviceKey string) (bool, string, error) {

	exist, err := dzdatabase.CheckExistKey(token)
	if err != nil || !exist {
		return false, "", utilities.NewError(utilities.DZErrorCodeTokenInvaild, "token invalid")
	}
	value, err := dzdatabase.RedisGetValueByKey(token)
	if err != nil {
		return false, "", utilities.NewError(utilities.DZErrorCodeTokenInvaild, "get token data is error")
	}
	userGuid, dg := splitGetUserGUIDAndDeviceGuid(value)
	if dg != deviceKey {
		return false, "", utilities.NewError(utilities.DZErrorCodeTokenInvaild, "device is not auth")
	}
	return true, userGuid, nil
}

func ChackAppKeyIsVaild(userGuid string, appkey string) (bool, error) {

	// app, err := dzdatabase.CheckUserAuthApp(userGuid, appkey)
	// if err != nil {
	// 	return false, err
	// }
	// if !app. {
	// 	return false, utilities.NewError(DZAppErrorCodeNotVaild, "app key is not vaild for this user")
	// }
	return true, nil
}

func HandleUpdateToken(json *simplejson.Json, devicekey string) ([]byte, error) {
	token, err := json.Get("token").String()
	if err != nil {
		return nil, utilities.NewError(utilities.DZErrorCodePaser, "parse token error")
	}
	err = UpdateTokenExpireTime(token)
	if err != nil {
		return nil, utilities.NewError(utilities.DZErrorCodePaser, "length token deadline error")
	}
	rj, err := simplejson.NewJson([]byte("{}"))
	if err != nil {
		return nil, utilities.NewError(utilities.DZErrorCodePaser, "new json error!")
	}
	rj.Set("token", token)
	return rj.Encode()
}
