package apps

import (
	"dzdatabase"
	"models"
	"networks"
	"service"
)

func HanldeRegisterApp(requestData *networks.DZRequstData) ([]byte, error) {
	app, err := models.NewAppFromJSON(requestData.BodyJson)
	if err != nil {
		return nil, err
	}
	err = dzdatabase.UpdateDZApp(app)
	if err != nil {
		return nil, err
	}
	json := service.ResponseSuccessJSON("update app ok")
	return json.Encode()
}
