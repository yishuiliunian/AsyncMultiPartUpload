package authorization

import (
	"authorization"
	"github.com/bitly/go-simplejson"
	"utilities"
)

func HandleUpdateToken(json *simplejson.Json) ([]byte, error) {
	token, err := json.Get("token").String()
	if err != nil {
		return nil, utilities.NewError(utilities.DZErrorCodePaser, "parse token error")
	}
	err = authorization.LengthenDeadlineForToken(token)
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
