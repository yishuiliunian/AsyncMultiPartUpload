package dzdatabase

import (
	"github.com/garyburd/redigo/redis"
	"utilities"
)

const (
	REDISErrorCodeSet         = -6000
	REDISErrorCodeExpire      = -6001
	REDISErrorCodeGET         = -6002
	REDISErrorCodeDialConnect = -6003
	REDISErrorCodeNotExistKey = -6004
	REDISErrorCodeOpen        = -6005
)

const (
	REDISCommandExists string = "EXISTS"
	REDISCommandGET    string = "GET"
	REDISCommandSET    string = "SET"
	REDISCommandEXPIRE string = "EXPIRE"
	REDISHttpMethod    string = "tcp"
	REDISPort          string = ":6379"
)

const (
	DZTokenDefaultDeadlineDuration = 60 * 60
)

func AddExpireKeyValueToReids(key string, value string) error {
	conn, err := redis.Dial(REDISHttpMethod, REDISPort)
	defer conn.Close()
	if err != nil {
		return utilities.NewError(REDISErrorCodeDialConnect, "dial redis error")
	}
	_, err = conn.Do(REDISCommandSET, key, value)
	if err != nil {
		return utilities.NewError(REDISErrorCodeSet, key)
	}
	_, err = conn.Do(REDISCommandEXPIRE, key, DZTokenDefaultDeadlineDuration)
	if err != nil {
		return utilities.NewError(REDISErrorCodeExpire, "redis get key error")
	}
	return nil
}

func RedisGetValueByKey(key string) (string, error) {
	conn, err := redis.Dial(REDISHttpMethod, REDISPort)
	defer conn.Close()
	if err != nil {
		return "", utilities.NewError(REDISErrorCodeDialConnect, "dial redis error")
	}
	return redis.String(conn.Do(REDISCommandGET, key))
}

func CheckExistKey(key string) (bool, error) {
	conn, err := redis.Dial(REDISHttpMethod, REDISPort)
	defer conn.Close()

	if err != nil {
		return false, utilities.NewError(REDISErrorCodeOpen, "open redis error!")
	}
	return redis.Bool(conn.Do(REDISCommandExists, key))
}

func UpdateExpireByKey(key string) error {
	e, err := CheckExistKey(key)
	if err != nil {
		return utilities.NewError(REDISErrorCodeGET, "check exist key error")
	}
	if !e {
		return utilities.NewError(REDISErrorCodeNotExistKey, "the key your find does not exist")
	}
	conn, err := redis.Dial(REDISHttpMethod, REDISPort)
	defer conn.Close()
	if err != nil {
		return utilities.NewError(REDISErrorCodeOpen, "open redis error!")
	}
	_, err = conn.Do(REDISCommandEXPIRE, key, DZTokenDefaultDeadlineDuration)
	if err != nil {
		return utilities.NewError(REDISErrorCodeExpire, "redis get key error")
	}
	return nil
}
