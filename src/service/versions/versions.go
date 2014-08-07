package versions

import (
	"dzdatabase"
	"fmt"
	"github.com/bitly/go-simplejson"
)

func HangleGetAllVersionsRequest(json *simplejson.Json, userGuid string) ([]byte, error) {
	fmt.Println("get method get versions")
	timesVersion, err := dzdatabase.GetTimeVersionWithUserGuid(userGuid)
	if err != nil {
		return nil, err
	}
	fmt.Println("time version %d", timesVersion)
	typesVersion, err := dzdatabase.GetTimeTypesVersionWithUserGuid(userGuid)
	if err != nil {
		return nil, err
	}
	fmt.Println("type version %d", typesVersion)

	deletedVersion, err := dzdatabase.GetTimeDeletedVersionWithUserGuid(userGuid)
	if err != nil {
		return nil, err
	}

	rj, _ := simplejson.NewJson([]byte("{}"))
	rj.Set("times", timesVersion)
	rj.Set("types", typesVersion)
	rj.Set("deletedobjects", deletedVersion)
	return rj.MarshalJSON()
}
