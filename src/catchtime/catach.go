package catchtime

import (
	"dzdatabase"
	"fmt"
	"github.com/bitly/go-simplejson"
	"models"
)

func HandleRemoveTime(dt *simplejson.Json) error {
	return nil

}

func HandleUpdateTime(dt *simplejson.Json) error {
	dtdata, err := models.NewDZTimeFromJSON(dt)
	if err != nil {
		return err
	}
	if dtdata == nil {
		fmt.Println("nil")
	}
	err = dzdatabase.UpdateDZTime(dtdata)
	if err != nil {
		return err
	}
	return nil
}
