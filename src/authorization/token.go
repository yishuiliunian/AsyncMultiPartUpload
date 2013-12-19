package authorization

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/garyburd/redigo/redis"
	"time"
	"utilities"
)

const (
	KeyToken = "token"
)

const (
	REDISCommandExists string = "EXISTS"
	REDISCommandGET    string = "GET"
	REDISCommandSET    string = "SET"
	REDISCommandEXPIRE string = "EXPIRE"
)

const (
	DZTokenDefaultDeadlineDuration = 60 * 60
)

const (
	_ObjectKeyIdentify string = "identify"
)

type DZToken struct {
	Identify string
	UserGUID string
}

func NewToken() *DZToken {
	return &DZToken{utilities.GUID(), ""}
}

func (d *DZToken) TOJSONObjectString() []byte {
	json, err := simplejson.NewJson([]byte("{}"))
	if err != nil {
		return nil
	}
	json.Set(_ObjectKeyIdentify, d.Identify)
	str, err := json.MarshalJSON()

	if err != nil {
		return nil
	}
	return []byte(str)
}

func DZTokenFromJsonString(jsonStr []byte) *DZToken {
	token := &DZToken{}
	json, err := simplejson.NewJson(jsonStr)
	if err != nil {
		return nil
	}
	identify, err := json.Get(_ObjectKeyIdentify).String()
	if err != nil {
		return nil
	}
	token.Identify = identify
	return token
}

func getTokenDataByIdentify(identify string, conn redis.Conn) *DZToken {

	str, err := redis.Bytes(conn.Do(REDISCommandGET, identify))
	fmt.Println(str)
	fmt.Println(err)
	if err != nil {
		return nil
	}
	return DZTokenFromJsonString(str)
}

func addToken(token *DZToken, conn redis.Conn) error {
	str := token.TOJSONObjectString()
	a, err := conn.Do(REDISCommandSET, token.Identify, str)
	fmt.Println("****************")
	fmt.Println(a)
	fmt.Println(err)
	return err
}

func checkExistToken(identify string, conn redis.Conn) (bool, error) {
	return redis.Bool(conn.Do(REDISCommandExists, identify))
}

func CheckTokenIsVaild(token string) (bool, error) {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		return false, err
	}
	defer conn.Close()
	exist, err := checkExistToken(token, conn)
	if !exist {
		return false, nil
	}
	_, err = conn.Do(REDISCommandGET, token)
	if err != nil {
		return false, nil
	}
	return true, nil
}

func VaildToken(token *DZToken) error {
	conn, err := redis.Dial("tcp", ":6379")
	defer conn.Close()
	if err != nil {
		return err
	}
	token.Deadline = token.Deadline.Add(DZTokenDefaultDeadlineDuration)
	return addToken(token, conn)
}
func LengthenDeadlineForToken(token string) error {
	conn, err := redis.Dial("tcp", ":6379")
	defer conn.Close()
	fmt.Println(token)
	e, _ := checkExistToken(token, conn)
	if !e {
		return utilities.NewError(utilities.DZErrorCodeTokenInvaild, "token invaild not exist!")
	}
	if err != nil {
		return err
	}
	t := getTokenDataByIdentify(token, conn)
	t.Deadline = t.Deadline.Add(DZTokenDefaultDeadlineDuration)
	return addToken(t, conn)
}

func ApplyAnVaildToken() *DZToken {
	token := NewToken()
	VaildToken(token)
	return token
}
