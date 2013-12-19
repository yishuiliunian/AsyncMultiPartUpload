package utilities

import (
	"time"
)

const (
	DZDefaultDateLayout = "20060102T15:04:05+800"
)

func ParseTimeString(str string) (time.Time, error) {
	return time.Parse(DZDefaultDateLayout, str)
}

func EncodeTime(t *time.Time) string {
	return t.Format(DZDefaultDateLayout)
}
