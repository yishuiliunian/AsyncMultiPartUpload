package dzdatabase

import (
	"fmt"
	"labix.org/v2/mgo/bson"
	"models"
	"utilities"
)

func IsExistUserByEmail(email string) (bool, error) {
	session := ShareDBSessionPool().OneSession()
	defer ShareDBSessionPool().EndUseSession(session)
	count, err := session.CollectionUsers().Find(bson.M{models.DZObjectKeyEmail: email}).Count()
	fmt.Println(count)
	if err != nil {
		return false, err
	}
	return (count > 0), nil
}

func DZUserByEmail(email string) (*models.DZUser, error) {
	session := ShareDBSessionPool().OneSession()
	defer ShareDBSessionPool().EndUseSession(session)
	var user models.DZUser
	err := session.CollectionUsers().Find(bson.M{models.DZObjectKeyEmail: email}).One(&user)
	return &user, err
}

func RegisterUser(user *models.DZUser) error {
	eixst, err := IsExistUserByEmail(user.Email)
	if err != nil {
		return err
	}
	if eixst {
		return utilities.NewError(utilities.DZErrorCodeOperation, "user exist!")
	}
	session := ShareDBSessionPool().OneSession()
	defer ShareDBSessionPool().EndUseSession(session)
	c := session.CollectionUsers()
	c.Insert(user)
	return nil
}
