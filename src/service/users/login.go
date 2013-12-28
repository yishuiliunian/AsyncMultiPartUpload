package users

import (
	"dzdatabase"
	"fmt"
	"github.com/bitly/go-simplejson"
	"models"
	"networks"
	"service/authorization"
	"strings"
	"utilities"
)

func HandleRegisterUser(json *simplejson.Json) ([]byte, error) {
	duser, err := models.NewDZUserFromJson(json)
	if err != nil {
		return nil, err
	}
	duser.Guid = utilities.GUID()
	err = dzdatabase.RegisterUser(duser)
	if err != nil {
		return nil, err
	}
	fmt.Println(json)
	device := json.Get("device")
	if device != nil {
		fmt.Println("*****")
		de, err := models.NewDeviceWithJson(device)
		if err != nil {
			fmt.Println("decode device error %s", err)
		}
		dzdatabase.UpdateDZDevice(de)
		fmt.Println(de)
	}
	j, _ := simplejson.NewJson([]byte("{}"))
	j.Set("returcode", 200)
	j.Set("userguid", duser.Guid)
	bs, _ := j.MarshalJSON()
	return bs, nil
}

func HandleLoginUser(json *simplejson.Json, devicekey string) ([]byte, error) {

	fmt.Println("handle login")
	duser, err := models.NewDZUserFromJson(json)
	if err != nil {
		return nil, err
	}
	exsit, err := dzdatabase.IsExistUserByEmail(duser.Email)
	if err != nil {
		return nil, err
	}
	if !exsit {
		return nil, utilities.NewError(utilities.DZErrorCodeUserNotExist, "user not exist!")
	}
	user, err := dzdatabase.DZUserByEmail(duser.Email)
	if err != nil {
		return nil, err
	}
	if !strings.EqualFold(duser.Password, user.Password) {
		return nil, utilities.NewError(utilities.DZErrorCodeTokenUnSupoort, "password error")
	}
	token, err := authorization.ApplyAnVaildToken(user.Guid, devicekey)
	if err != nil {
		return nil, err
	}
	rejson, err := simplejson.NewJson([]byte("{}"))
	if err != nil {
		return nil, utilities.NewError(utilities.DZErrorCodeTokenUnSupoort, "new json error")
	}
	rejson.Set(networks.DZProtocolKeyToken, token)
	rejson.Set("userGuid", user.Guid)
	return rejson.MarshalJSON()
}
