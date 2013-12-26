package dzdatabase

import (
	"fmt"
	"labix.org/v2/mgo/bson"
	"models"
)

const (
	DBKeyTypeGuid       = "guid"
	DBKeyTypeName       = "name"
	DBKeyTypeUserGuid   = "userGuid"
	DBKeyTypeDetail     = "detail"
	DBKeyTypeOtherInfos = "otherinfos"
	DZAppKeyTypeVersion = "version"
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
		err := s.CollectionTimes().Update(bson.M{"dzobject.guid": dbType.Guid},
			bson.M{"$set": bson.M{models.DZObjectKeyVersion: dt.Version,
				DBKeyTypeDetail:     dt.Detail,
				DBKeyTypeUserGuid:   dt.UserGuid,
				DBKeyTypeName:       dt.Name,
				DBKeyTypeOtherInfos: dt.OtherInfos}})
		return err
	} else {
		fmt.Println("insert")
		return s.CollectionTimes().Insert(dt)
	}
}
