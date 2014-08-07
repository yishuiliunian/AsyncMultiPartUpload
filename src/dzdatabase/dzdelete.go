package dzdatabase

import (
	"labix.org/v2/mgo/bson"
	"models"
)

func DZDeletedObjectByGuid(guid string)(*models.DZDeleteObject, error) {
	s := ShareDBSessionPool().OneSession()
	defer ShareDBSessionPool().EndUseSession(s)
	var dt models.DZDeleteObject
	err := s.CollectionDeletedObjects().Find(bson.M{"dzobject.guid":guid}).One(&dt)
	if err != nil {
		return nil, err
	}
	return &dt, err
}

func UpdateDZDeletedObject(dt* models.DZDeleteObject) error {
 	s := ShareDBSessionPool().OneSession()
	defer ShareDBSessionPool().EndUseSession(s)
	dbDeleted, _ := DZDeletedObjectByGuid(dt.Guid)
	var err error
	dt.Version, err = IncreaseDeleteObjectsVersion(dt.UserGUID)
	if err != nil {
		return err
	}
	if dbDeleted != nil {
		err := s.CollectionDeletedObjects().Update(bson.M{"dzobject.guid": dbDeleted.Guid},
			bson.M{"$set":bson.M{models.JDK_DeletedVersion : dt.Version,
				models.JDK_DeletedType: dt.Type,
				models.JDK_DeletedTime: dt.DeletedTime}})
		return err
	} else {
		return s.CollectionDeletedObjects().Insert(dt)
	}
}


func GetDeletedObjectsOfUserWithVersionSpace(userguid string, startVersion int64, endVersion int64) ([]models.DZDeleteObject, error) {
	s := ShareDBSessionPool().OneSession()
	defer ShareDBSessionPool().EndUseSession(s)
	var times []models.DZDeleteObject
	err := s.CollectionDeletedObjects().Find(bson.M{models.DZObjectKeyVersion: bson.M{MongoMethodGreatThan: startVersion,
		MongoMethodLittleThan: endVersion}, models.DZObjectKeyUserGuid: userguid}).Sort(models.DZObjectKeyVersion).All(&times)
	return times, err
}