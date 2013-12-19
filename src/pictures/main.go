package main

import (
	"catchtime"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"log"
	"models"
	"net/http"
	"networks"
	"restfulbase"
	"service"
	"service/authorization"
	"service/users"
	"utilities"
)

func InitApp() {
	models.ShareImageUpdateMonitor()
}

var dataMap = map[string]models.DZImage{}

func getImageData(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		body := req.Body
		defer body.Close()
		bodystr, _ := ioutil.ReadAll(body)
		var dataError error
		json, err := simplejson.NewJson(bodystr)
		if err != nil {
			dataError = err
		} else {
			part, err := models.NewDataPartWithJson(json)
			if err != nil {
				dataError = err
			} else {
				err := service.AddFilePartData(*part)
				if err != nil {
					dataError = err
				}
			}
		}
		if dataError != nil {
			json := utilities.EncodeError(dataError)
			fmt.Println(json)
			data, _ := json.Encode()
			fmt.Println(data)
			rw.Write(data)
			return
		}
		rw.Write([]byte("nil"))
	}
}

func routeToMethod(json *simplejson.Json, method string) ([]byte, error) {
	switch method {
	case restfulbase.DZRestMethodTimeUpdate:
		{
			return []byte("{ok}"), catchtime.HandleUpdateTime(json)
		}
	case restfulbase.DZRestMethodUserRegister:
		{
			return users.HandleRegisterUser(json)
		}
	case restfulbase.DZRestMethodTimeLogin:
		{
			return users.HandleLoginUser(json)
		}
	case restfulbase.DZRestMethodTokenUpdate:
		{
			return authorization.HandleUpdateToken(json)
		}
	default:
		{
			return nil, utilities.NewError(utilities.DZErrorCodeTokenUnSupoort, "not support")
		}
	}
}

func handleJsonRequst(rw http.ResponseWriter, req *http.Request) {
	if req.Method == networks.NetworkMethodPost {

		reqData, err := networks.DecodeHttpRequest(req)
		if err != nil {
			errjson := utilities.EncodeError(err)
			str, err := errjson.MarshalJSON()
			if err == nil {
				rw.Write(str)
			} else {
				rw.Write([]byte("code error error!"))
			}
		} else {
			data, err := routeToMethod(reqData.BodyJson, reqData.Method)
			if err != nil {
				errjson := utilities.EncodeError(err)
				str, err := errjson.MarshalJSON()
				if err == nil {
					rw.Write(str)
				} else {
					rw.Write([]byte("code error error!"))
				}
			} else {
				rw.Write(data)
			}

		}
	} else {
		rw.Write([]byte("only use with json restful"))
	}
}

func main() {
	InitApp()
	http.HandleFunc("/", getImageData)
	http.HandleFunc("/json", handleJsonRequst)
	fmt.Println("lisent at localhost:9091")
	err := http.ListenAndServe(":9091", nil)

	if err != nil {
		log.Fatal(err)
	}
}
