package models

import (
	"fmt"
	"utilities"
)

type DZImageUpdateMonitor struct {
}

func (d *DZImageUpdateMonitor) handleImageUpdate(event *utilities.Event) {
	image := event.Params[utilities.DZParamsKeyDZImage]
	fmt.Println(image)
	fmt.Println("get image")
	// session := dzdatabase.ShareDBSessionPool().OneSession()
	// session.CollectionPictures().Insert(image.(*DZImage))
	// n, err := session.CollectionPictures().Count()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("image count %d", n)
	// }
	fmt.Println("asdfasd")
}

func (d *DZImageUpdateMonitor) Init() {
	var callback utilities.EventCallback = d.handleImageUpdate
	utilities.ShareNotificationCenter().AddEventListener(utilities.DZMessageImageDidGetNewData, &callback)
}

var _shareInstance *DZImageUpdateMonitor

func ShareImageUpdateMonitor() *DZImageUpdateMonitor {
	if _shareInstance == nil {
		_shareInstance = &DZImageUpdateMonitor{}
		_shareInstance.Init()
	}
	return _shareInstance
}
