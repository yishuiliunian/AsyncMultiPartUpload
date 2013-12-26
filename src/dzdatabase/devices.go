package dzdatabase

import (
	"fmt"
	"labix.org/v2/mgo/bson"
	"models"
)

func IsExistDeviceWithGuid(guid string) (bool, error) {
	s := ShareDBSessionPool().OneSession()
	defer ShareDBSessionPool().EndUseSession(s)
	count, err := s.CollectionUsers().Find(bson.M{"dzobject.guid": guid}).Count()
	return count > 0, err
}

func DZDeviceByGuid(guid string) (*models.DZDevice, error) {
	s := ShareDBSessionPool().OneSession()
	defer ShareDBSessionPool().EndUseSession(s)
	var dv *models.DZDevice
	err := s.CollectionDiveces().Find(bson.M{"dzobject.guid": guid}).One(dv)
	return dv, err
}

func UpdateDZDevice(device *models.DZDevice) error {
	s := ShareDBSessionPool().OneSession()
	defer ShareDBSessionPool().EndUseSession(s)
	e, _ := IsExistDeviceWithGuid(device.Guid)
	if e {
		dv, err := DZDeviceByGuid(device.Guid)
		if err != nil {
			return err
		}
		fmt.Println("update device %s", device.Guid)

		return s.CollectionDiveces().Update(bson.M{"dzobject.guid": device.Guid},
			bson.M{"$set": bson.M{
				models.DZObjectKeyUserGuid:   dv.UserGUID,
				models.DZObjectKeyOtherInfos: dv.OtherInfos,
				models.DZObjectKeyDetail:     dv.Detail,
				models.DZObjectKeyName:       dv.Name,
				"activedevices":              dv.ActiveDevices}})
	} else {
		fmt.Println("insert device %s", device.Guid)
		return s.CollectionDiveces().Insert(device)
	}
	return nil
}
