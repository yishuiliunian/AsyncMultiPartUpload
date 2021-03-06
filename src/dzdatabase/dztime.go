package dzdatabase

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"models"
)

func DZTimeByGuid(guid string) (*models.DZTime, error) {
	s := ShareDBSessionPool().OneSession()
	defer ShareDBSessionPool().EndUseSession(s)
	var dt models.DZTime
	err := s.CollectionTimes().Find(bson.M{"dzobject.guid": guid}).One(&dt)
	if err != nil {
		return nil, err
	}
	return &dt, nil
}
func UpdateDZTime(dt *models.DZTime) error {
	s := ShareDBSessionPool().OneSession()
	defer ShareDBSessionPool().EndUseSession(s)
	dbtime, _ := DZTimeByGuid(dt.Guid)
	var err error
	dt.Version, err = IncreaseTimesVersion(dt.UserGUID)
	if err != nil {
		return err
	}
	if dbtime != nil {
		err := s.CollectionTimes().Update(bson.M{"dzobject.guid": dbtime.Guid},
			bson.M{"$set": bson.M{models.JDK_TimeVersion: dt.Version,
				models.JDK_TimeDateBegin:  dt.DateBegin,
				models.JDK_TimeDateEnd:    dt.DateEnd,
				models.JDK_TimeTypeGUID:   dt.TypeGUID,
				models.JDK_TimeDetail:     dt.Detail,
				models.JDK_TimeUserGUID:   dt.UserGUID,
				models.JDK_TimeDeviceGUID: dt.DeviceGUID}})
		return err
	} else {
		return s.CollectionTimes().Insert(dt)
	}
}
func RemoveDZTime(guid string) error {
		s := ShareDBSessionPool().OneSession()
	defer ShareDBSessionPool().EndUseSession(s)
	err := s.CollectionTimes().Remove(bson.M{"dzobject.guid": guid})
	if err != nil && err != mgo.ErrNotFound {
		return	err;
	}
	return nil
}

func GetTimesOfUserWithVersionSpace(userguid string, startVersion int64, endVersion int64) ([]models.DZTime, error) {
	s := ShareDBSessionPool().OneSession()
	defer ShareDBSessionPool().EndUseSession(s)
	var times []models.DZTime
	err := s.CollectionTimes().Find(bson.M{models.DZObjectKeyVersion: bson.M{MongoMethodGreatThan: startVersion,
		MongoMethodLittleThan: endVersion}, models.DZObjectKeyUserGuid: userguid}).Sort(models.DZObjectKeyVersion).All(&times)
	return times, err
}



