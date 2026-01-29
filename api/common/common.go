package common

import "strings"

// Descrption: 字符串切割成数组
func SmartSplit(str string, sep string) []string {
	//避免空字符串返回[""]
	if str == "" {
		return []string{}
	}
	res := strings.Split(str, sep)
	return res
}
