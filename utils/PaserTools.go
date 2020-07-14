package utils

import (
	"errors"
	"strings"
)

const (
	ActionPrefix = "@"
	ItemPrefix = "#"
	CurQusPrefix = "~" // ~表示当前项名称
)

func GetPair(str string, check func(string)bool) func(func(strS [2]string)) {
	return func(calc func(strS [2]string)){
		var res [2]string
		if check(str) {
			arr := strings.Split(str, ":")
			if len(res) != 2 {
				panic(errors.New("syntax error: " + str))
			}
			res[0], res[1] = strings.TrimLeft(arr[0], ItemPrefix), arr[1]
			calc(res)
		}
	}
}

// 解析 "@action:value" 字符串形式
func GetActionStr(line string) (action string, value string) {
	if strings.HasPrefix(line, "@") {
		params := strings.Split(strings.Split(line, "@")[1], ":")
		if len(params) != 2 {
			return params[0], ""
		}
		return params[0], strings.TrimSpace(params[1])
	}
	panic(errors.New("format error: " + line))
}

// 解析title
func GetTitle(line string) string {
	if strings.HasPrefix(line, "title:") {
		return strings.Split(line, ":")[1]
	}
	panic(errors.New("format error: " + line))
}
