package dzdatabase

import (
	"fmt"
	"labix.org/v2/mgo/bson"
	"models"
	"labix.org/v2/mgo"
)

const (
	DBKeyTypeGuid       = "guid"
	DBKeyTypeName       = "name"
	DBKeyTypeUserGuid   = "userguid"
	DBKeyTypeDetail     = "detail"
	DBKeyTypeOtherInfos = "otherinfos"
	DZAppKeyTypeVersion = "version"
	DBKeyTypeFinished	= "finished"
)

func DZTimeTypeByGuid(guid string) (*models.DZTimeType, error) {
	s := ShareDBSessionPool().OneSession()
	defer ShareDBSessionPool().EndUseSession(s)
	var dt models.DZTimeType
	err := s.CollectionTimeTypes().Find(bson.M{"dzobject.guid": guid}).One(&dt)
	if err != nil {
		return nil, nil
	}
	return &dt, nil
}

func RemoveDZTimeType(guid string) error {
	s := ShareDBSessionPool().OneSession()
	defer ShareDBSessionPool().EndUseSession(s)
	err := s.CollectionTimeTypes().Remove(bson.M{"dzobject.guid": guid})
	if err != nil && err != mgo.ErrNotFound{
		return err
	}
	return nil
}

func UpdateDZTimeType(dt *models.DZTimeType) error {
	s := ShareDBSessionPool().OneSession()
	defer ShareDBSessionPool().EndUseSession(s)
	dbType, err := DZTimeTypeByGuid(dt.Guid)
	if err != nil {
		return err
	}
	dt.Version, err = IncreaseTimeTypesVersion(dt.UserGuid)
	if err != nil {
		return err
	}
	
	if dbType != nil {
		fmt.Println("update")
		fmt.Println(dt)
		err := s.CollectionTimeTypes().Update(bson.M{"dzobject.guid": dbType.Guid},
			bson.M{"$set": bson.M{models.DZObjectKeyVersion: dt.Version,
				DBKeyTypeDetail:     dt.Detail,
				DBKeyTypeUserGuid:   dt.UserGuid,
				DBKeyTypeName:       dt.Name,
				DBKeyTypeOtherInfos: dt.OtherInfos,
				DBKeyTypeFinished : dt.Finished}})
		return err
	} else {
		fmt.Println("insert")
		return s.CollectionTimeTypes().Insert(dt)
	}
}

func GetTimeTypesOfUserWithVersionSpace(userguid string, startVersion int64, endVersion int64) ([]models.DZTimeType, error) {
	s := ShareDBSessionPool().OneSession()
	defer ShareDBSessionPool().EndUseSession(s)
	var times []models.DZTimeType
	err := s.CollectionTimeTypes().Find(bson.M{models.DZObjectKeyVersion: bson.M{MongoMethodGreatThan: startVersion,
		MongoMethodLittleThan: endVersion}, models.DZObjectKeyUserGuid: userguid}).Sort(models.DZObjectKeyVersion).All(&times)
	return times, err
}
