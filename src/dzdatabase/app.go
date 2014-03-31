package dzdatabase

import (
	"labix.org/v2/mgo/bson"
	"models"
	"utilities"
)

const (
	DZAppKeyGuid         = "guid"
	DZAppKeyOwnerGuid    = "ownerguid"
	DZAppKeyName         = "name"
	DZAppKeyReferenceURL = "referenceurl"
	DZAppKeyDescription  = "description"
	DZAppKeyVersion      = "version"
	DZAppKeyType         = "type"
	DZAppKeySubType      = "subtype"
	DZAppKeyPlatfomat    = "paltfomat"
	DZAppKeyTags         = "tags"
	DZAppKeyDetail       = "detail"
	DZAppKeyIsVaild      = "isvaild"
)

func IsExistAppWithGuid(guid string) (bool, error) {
	s := ShareDBSessionPool().OneSession()
	defer ShareDBSessionPool().EndUseSession(s)
	count, err := s.CollectionApps().Find(bson.M{"dzobject.guid": guid}).Count()
	if err != nil {
		return false, utilities.NewError(DZDBErrorCodeCantFindObject, "app not exist")
	}
	return count > 0, nil
}

func DZAppWithGuid(guid string) (*models.DZApp, error) {
	s := ShareDBSessionPool().OneSession()
	defer ShareDBSessionPool().EndUseSession(s)
	var app models.DZApp
	err := s.CollectionApps().Find(bson.M{"dzobject.guid": guid}).One(&app)
	return &app, err
}

func UpdateDZApp(app *models.DZApp) error {
	s := ShareDBSessionPool().OneSession()
	defer ShareDBSessionPool().EndUseSession(s)
	exist, err := IsExistAppWithGuid(app.Guid)
	if err != nil {
		return err
	}
	if exist {
		return s.CollectionApps().Update(bson.M{"dzobject.guid": app.Guid},
			bson.M{"$set": bson.M{
				DZAppKeyGuid:         app.Guid,
				DZAppKeyDescription:  app.Description,
				DZAppKeyName:         app.Name,
				DZAppKeyOwnerGuid:    app.OwnerGuid,
				DZAppKeyReferenceURL: app.ReferenceURL,
				DZAppKeyVersion:      app.Version,
				DZAppKeyDetail:       app.Detail,
				DZAppKeyType:         app.Type,
				DZAppKeySubType:      app.SubType,
				DZAppKeyPlatfomat:    app.Platfomat,
				DZAppKeyTags:         app.Tags,
				DZAppKeyIsVaild:      app.IsVaild}})
	} else {
		return s.CollectionApps().Insert(&app)
	}
}

func CheckUserAuthApp(userGuid string, appkey string) (bool, error) {
	s := ShareDBSessionPool().OneSession()
	defer ShareDBSessionPool().EndUseSession(s)
	count, err := s.CollectionApps().Find(bson.M{"dzobject.guid": appkey, DZAppKeyOwnerGuid: userGuid}).Count()
	if err != nil {
		return false, utilities.NewError(DZDBErrorCodeCantFindObject, "app not exist")
	}
	return count > 0, nil
}
