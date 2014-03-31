package utilities

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"reflect"
)

const (
	DZJSONObjectKeyErrorCode    = "errorcode"
	DZJSONObjectKeyErrorMessage = "errormessage"
)
const (
	DZErrorCodeNetWork        = -6001
	DZErrorCodePaser          = -6002
	DZErrorCode               = -6003
	DZErrorCodeWriteFile      = -6004
	DZErrorCodeOperation      = -6005
	DZErrorCodeMD5NotMactch   = -6006
	DZErrorCodeDataDirty      = -6007
	DZErrorCodeTokenInvaild   = -6008
	DZErrorCodeTokenUnSupoort = -6009
	DZErrorCodeUserNotExist   = -6010
)

func GetStructValue(value reflect.Value, key string) (rs reflect.Value) {
	if !value.IsValid() {
		return
	}
	if value.Type().Kind() == reflect.Ptr {
		//value.Elem()可以得到指针所指向的对象
		if value.Elem().Kind() != reflect.Struct {
			return
		}
	} else if value.Type().Kind() == reflect.Struct {
		return
	}

	//好了,来取Struct的Field吧!

	//首先,我们把*T还原为T
	//如果本来就是Struct,那么只是简单返回而已
	//指针类型是不能获取Field的
	v := reflect.Indirect(value)
	field := v.FieldByName(key)
	if field.IsValid() { //字段存在时返回true
		rs = field
		return
	}

	//接下来,看看有米有对应的Method
	//注意,如果是*T,那么全部方法都能拿到
	//如果是T,那么只能获取那些非指针的方法哦
	//我也很纠结这个,尝试突破但没有成功
	t := value.Type()
	method, ok := t.MethodByName(key)
	if !ok { //没找到
		return
	}

	//输入的参数必须为1,也就是当前value,当然,如果你知道其他参数,也可以是传参的,也就一个数组嘛
	//输出的参数不为0就好了,我们只需要取第一个
	if method.Func.Type().NumIn() != 1 || method.Func.Type().NumOut() == 0 {
		return
	}
	//调用之
	rs = method.Func.Call([]reflect.Value{value})[0] //最后的[0]就是取第一个返回值
	return
}

type DZError struct {
	Code    int
	Message string
}

func (d *DZError) String() string {
	return d.Message
}

func (d *DZError) Error() string {
	return d.Message
}

func NewError(code int, message string) *DZError {
	return &DZError{code, message}
}

func EncodeError(err error) *simplejson.Json {
	json, _ := simplejson.NewJson([]byte("{}"))
	json.Set(DZJSONObjectKeyErrorMessage, err.Error())
	code := GetStructValue(reflect.ValueOf(err), "Code")
	var ec int64
	if code.IsValid() {
		ec = code.Int()
	} else {
		ec = -5000
	}
	fmt.Println(code)
	json.Set(DZJSONObjectKeyErrorCode, ec)
	return json
}
