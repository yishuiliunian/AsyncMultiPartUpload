package authorization

import (
	"dzdatabase"
	"strings"
	"utilities"
)

const (
	KeyToken = "token"
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

func CheckTokenIsVaild(token string) (bool, error) {
	return dzdatabase.CheckExistKey(token)
}
