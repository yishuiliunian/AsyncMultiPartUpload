package utilities

import (
	"encoding/base64"
	"fmt"
	"os"
)

const (
	base64Table = "sdfsdfsdfsdf123QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912"
)

func GUID() string {
	f, _ := os.OpenFile("/dev/urandom", os.O_RDONLY, 0)
	b := make([]byte, 16)
	f.Read(b)
	f.Close()
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

var coder = base64.NewEncoding(base64Table)

func Base64Encode(src []byte) []byte {
	return []byte(coder.EncodeToString(src))
}
func Base64Decode(src []byte) ([]byte, error) {
	return coder.DecodeString(string(src))
}
