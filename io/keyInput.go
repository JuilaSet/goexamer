package io

import (
	"bufio"
	"os"
	"strings"
)


// process
var reader *bufio.Reader
func init(){
	reader = bufio.NewReader(os.Stdin)
	strMap = make(map[string]string)
}

// 输入联调
var strMap map[string]string
func SetInputString(str string) {
	strMap[str] = ReadInput()
}

func GetInputString(str string) string {
	return strMap[str]
}

func Wait() {
	reader.ReadString('\n')
	//text, _ := reader.ReadString('\n')
}
func ReadInput() (text string) {
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	text = strings.Replace(text, "\n", "", -1)
	return
}

func ReadAndCompare(key string, trueFunc func(input string), falseFunc func(input string)) (b bool) {
	r := ReadInput()
	b = GetInputString(key) == r
	if b {
		trueFunc(r)
	} else {
		falseFunc(r)
	}
	return
}

func Count(r, key string) int {
	str := GetInputString(key)
	if r == str {
		return 1
	} else {
		return strings.Count(r, str)
	}
}
