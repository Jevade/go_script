package itemspy

import "strings"

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
func RemoveSpace(str string) string {
	str = strings.Replace(str, " ", "", -1)
	// 去除换行符
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\t", "", -1)
	return str
}
