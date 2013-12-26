package dzdatabase

import (
	"labix.org/v2/mgo/bson"
	"models"
	"time"
)

func IncreaseTimesVersion(userguid string) (int64, error) {
	return IncreaseShareVersionWith(DZDatabaseColletionTimes, userguid)
}

func IncreaseDevicessVersion(userguid string) (int64, error) {
	return IncreaseShareVersionWith(DZDataBaseColletionDevices, userguid)
}

func IncreaseTimeTypesVersion(userguid string) (int64, error) {
	return IncreaseShareVersionWith(DZDataBaseColletionTimeTypes, userguid)
}

func IncreaseUsersVersion(userguid string) (int64, error) {
	return IncreaseShareVersionWith(DZDataBaseColletionUsers, userguid)
}

func isExistVersion(keytype string, userguid string) (bool, error) {
	s := ShareDBSessionPool().OneSession()
	defer ShareDBSessionPool().EndUseSession(s)
	count, err := s.CollectionVersions().Find(bson.M{"userguid": userguid, "keytype": keytype}).Count()
	return count > 0, err
}

func getVersionByKey(key string, userguid string) (int64, error) {
	s := ShareDBSessionPool().OneSession()
	defer ShareDBSessionPool().EndUseSession(s)
	var v models.DZVersion
	err := s.CollectionVersions().Find(bson.M{"userguid": userguid, "keytype": key}).One(&v)
	if err != nil {
		return 0, nil
	}
	return v.Version, nil
}

func GetTimeVersionWithUserGuid(userguid string) (int64, error) {
	return getVersionByKey(DZDatabaseColletionTimes, userguid)
}

func GetTimeTypesVersionWithUserGuid(userguid string) (int64, error) {
	return getVersionByKey(DZDataBaseColletionTimeTypes, userguid)
}

func IncreaseShareVersionWith(keytype string, userguid string) (int64, error) {
	s := ShareDBSessionPool().OneSession()
	defer ShareDBSessionPool().EndUseSession(s)
	exist, _ := isExistVersion(keytype, userguid)

	if !exist {
		var version models.DZVersion
		version.KeyType = keytype
		version.UserGuid = userguid
		version.LastEditDate = "aa"
		version.Version = 0
		s.CollectionVersions().Insert(&version)
		return 0, nil
	}

	var version models.DZVersion
	err := s.CollectionVersions().Find(bson.M{"userguid": userguid, "keytype": keytype}).One(&version)
	if err != nil {
		version.KeyType = keytype
		version.UserGuid = userguid
		version.Version = 0
		version.LastEditDate = time.Now().String()
		err := s.CollectionVersions().Insert(version)
		return 0, err
	} else {
		err := s.CollectionVersions().Update(bson.M{"userguid": userguid, "keytype": keytype}, bson.M{"$set": bson.M{"version": version.Version + 1}})
		return version.Version + 1, err
	}
}
