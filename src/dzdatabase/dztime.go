package dzdatabase

import (
	"fmt"
	"labix.org/v2/mgo/bson"
	"models"
)

func DZTimeByGuid(guid string) (*models.DZTime, error) {
	s := ShareDBSessionPool().OneSession()
	defer ShareDBSessionPool().EndUseSession(s)
	var dt models.DZTime
	err := s.CollectionTimes().Find(bson.M{"dzobject.guid": guid}).One(&dt)
	if err != nil {
		return nil, nil
	}
	fmt.Println(dt)
	return &dt, nil
}
func UpdateDZTime(dt *models.DZTime) error {
	s := ShareDBSessionPool().OneSession()
	defer ShareDBSessionPool().EndUseSession(s)
	dbtime, err := DZTimeByGuid(dt.Guid)
	if err != nil {
		return err
	}
	dt.Version, err = IncreaseTimesVersion(dt.UserGUID)
	if err != nil {
		return err
	}
	if dbtime != nil {
		fmt.Println("update")
		err := s.CollectionTimes().Update(bson.M{"dzobject.guid": dbtime.Guid},
			bson.M{"$set": bson.M{models.DZObjectKeyVersion: dt.Version,
				models.DZObjectKeyDateBegin:  dt.DateBegin,
				models.DZObjectKeyDateEnd:    dt.DateEnd,
				models.DZObjectKeyTypeGuid:   dt.TypeGUID,
				models.DZObjectKeyDetail:     dt.Detail,
				models.DZObjectKeyUserGuid:   dt.UserGUID,
				models.DZObjectKeyDeviceGUID: dt.DeviceGUID}})
		return err
	} else {
		fmt.Println("insert")
		return s.CollectionTimes().Insert(dt)
	}
}
