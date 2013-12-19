package service

import (
	"file"
	"fmt"
	"models"
	"os"
	"strings"
	"utilities"
)

type DZOperation struct {
	running bool
}

func (d *DZOperation) Key() string {
	return string("operation")
}

func (d *DZOperation) Run() error {
	d.running = true
	return nil
}

func (d *DZOperation) IsRunning() bool {
	return d.running
}

func (d *DZOperation) PrepareForReuse() error {
	if d.running {
		return utilities.NewError(-1, string("asdfasd"))
	}
	return nil
}

type DZOpeationQueue struct {
	opMap map[string]*DZMultiPartWriteOperation
}

func (q *DZOpeationQueue) AddOperation(op *DZMultiPartWriteOperation) {
	key := op.Key()
	_, exist := q.opMap[key]
	if !exist {
		q.opMap[key] = op
	}
}

func (d *DZOpeationQueue) RemoveOperation(op *DZMultiPartWriteOperation) {
	key := op.Key()
	_, exist := d.opMap[key]
	if exist {
		delete(d.opMap, op.Key())
	}
}

func (d *DZOpeationQueue) ExistOpeationWithKey(key string) bool {
	_, exist := d.opMap[key]
	return exist
}

func (d *DZOpeationQueue) OperationByKey(key string) (*DZMultiPartWriteOperation, error) {
	if d.ExistOpeationWithKey(key) {
		op := d.opMap[key]
		return op, nil
	} else {
		err := utilities.NewError(-1, "not find")
		return nil, err
	}
}

func (d *DZOpeationQueue) handleFinishNotification(event *utilities.Event) {
	fmt.Println("end *****")
	if len(event.Params) > 0 {
		op := event.Params[utilities.DZParamsKeyOperation]
		a := op.(*DZMultiPartWriteOperation)
		d.RemoveOperation(a)
		fmt.Println("remove write operation %s", a.Key())
	}
}

func (d *DZOpeationQueue) Init() {
	notificenter := utilities.ShareNotificationCenter()
	var f utilities.EventCallback = d.handleFinishNotification
	notificenter.AddEventListener(utilities.DZMessageOperationFinish, &f)
	d.opMap = make(map[string]*DZMultiPartWriteOperation)
}

type DZMultiPartWriteOperation struct {
	DZOperation
	Count      int
	FileKey    string
	PartSN     int
	Md5        string
	DataChan   chan models.DZDataPart
	file       *os.File
	filePath   string
	finish     bool
	dataStatus []bool
	dataBuffer map[int][]byte
}

func NewDZMultiPartWriteOperation(data models.DZDataPart) *DZMultiPartWriteOperation {
	op := new(DZMultiPartWriteOperation)
	op.FileKey = data.Guid
	op.Count = data.PartCount
	op.PartSN = 0
	op.Md5 = data.SumMD5
	op.dataStatus = make([]bool, data.PartCount, data.PartCount)
	op.dataBuffer = make(map[int][]byte)
	return op
}

func (d *DZMultiPartWriteOperation) Run() error {
	err := d.DZOperation.Run()
	if err != nil {
		return err
	}
	path := file.JoinPath(d.FileKey)
	d.filePath = path
	d.file, err = os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		d.file.Close()
		return err
	}
	d.finish = false
	return nil
}

func (d *DZMultiPartWriteOperation) predataNotLocalied(index int) []int {
	var nldata []int = make([]int, 0, d.Count)
	for i := 0; i < index; i++ {
		if !d.dataStatus[i] {
			nldata = append(nldata, i)
		}
	}
	return nldata
}

func (d *DZMultiPartWriteOperation) writeDataAtIndex(data []byte, index int) error {
	predatas := d.predataNotLocalied(index)
	if len(predatas) == 0 {
		fmt.Println("write data %d", index)
		_, err := d.file.Write(data)
		if err != nil {
			return err
		}
		_, e := d.dataBuffer[index]
		if e {
			delete(d.dataBuffer, index)
		}
		d.dataStatus[index] = true
		for i, bufferData := range d.dataBuffer {
			d.writeDataAtIndex(bufferData, i)
		}
	} else {
		lastIndex := predatas[len(predatas)-1]
		fmt.Println("pre index is %d", lastIndex)
		predata, e := d.dataBuffer[lastIndex]
		if e {
			err := d.writeDataAtIndex(predata, lastIndex)
			if err != nil {
				return err
			}
		} else {
			fmt.Println("buffer data index %d", index)
			d.dataBuffer[index] = data
		}
	}

	return nil
}

func (d *DZMultiPartWriteOperation) handleData(data models.DZDataPart) error {
	if data.PartCount != d.Count {
		return utilities.NewError(utilities.DZErrorCodeDataDirty, "count not match")
	}
	if data.PartSN < 0 || data.PartSN >= d.Count {
		return utilities.NewError(utilities.DZErrorCodeDataDirty, "part count out index")
	}
	md5 := utilities.MD5Bytes(data.Data)
	if strings.ToLower(md5) != strings.ToLower(data.MD5) {
		return utilities.NewError(utilities.DZErrorCodeMD5NotMactch, "data md5 error!")
	}
	err := d.writeDataAtIndex(data.Data, data.PartSN)
	if err != nil {
		return err
	}
	for i, e := range d.dataStatus {
		fmt.Println("data %d exist is %d", i, e)
	}
	var allReviced bool = true
	for _, v := range d.dataStatus {
		allReviced = allReviced && v
	}
	fmt.Println("allRecivied %d", allReviced)
	if allReviced {
		d.finish = true
		params := make(map[string]interface{})
		params[utilities.DZParamsKeyOperation] = d
		utilities.ShareNotificationCenter().PosetEventWithNameAndInfos(utilities.DZMessageOperationFinish, params)

		if data.FileType == models.DZFileTypeImage {
			image := models.NewImage()
			image.Guid = data.Guid
			image.Md5 = data.SumMD5
			image.Version = 0
			image.LocalUrl = d.filePath
			params := make(map[string]interface{})
			params[utilities.DZParamsKeyDZImage] = image
			utilities.ShareNotificationCenter().PosetEventWithNameAndInfos(utilities.DZMessageImageDidGetNewData, params)
		}
	}
	return nil
}

func (d *DZMultiPartWriteOperation) Key() string {
	return d.FileKey
}

var _shareFileWriteQueue *DZOpeationQueue

func ShareFileWriteQueue() *DZOpeationQueue {
	if _shareFileWriteQueue == nil {
		_shareFileWriteQueue = new(DZOpeationQueue)
		_shareFileWriteQueue.Init()
	}
	return _shareFileWriteQueue
}

func AddFilePartData(data models.DZDataPart) error {
	op, _ := ShareFileWriteQueue().OperationByKey(data.Guid)
	if op != nil {
		if op.IsRunning() {
			return op.handleData(data)
		} else {
			return utilities.NewError(utilities.DZErrorCodeOperation, "operation has finished and not remove by queue")
		}
	} else {
		op := NewDZMultiPartWriteOperation(data)
		ShareFileWriteQueue().AddOperation(op)
		err := op.Run()
		if err != nil {
			return err
		}
		return op.handleData(data)
	}
}
