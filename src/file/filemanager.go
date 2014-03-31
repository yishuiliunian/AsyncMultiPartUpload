package file

const (
	DPUserDataPath string = "/Users/stonedong/godatas"
)

func JoinPath(name string) string {
	return DPUserDataPath + string('/') + name
}
