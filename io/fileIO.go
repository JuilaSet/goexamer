package io

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"goexamer/utils"
	"io/ioutil"
	"os"
)

type todoFunc func(line string, count int, fileLinesMark rune)

func EachLine(file *os.File) func(todoFunc) {
	scanner := bufio.NewScanner(file)
	var (
		count int
		errMsg string
		noFailed bool
		fileLinesMark rune
		check = utils.CheckLinesHandler()
	)
	return func(todo todoFunc){
		count = 0
		for scanner.Scan() {
			count++
			// 读取当前行内容
			line := scanner.Text()
			// 语法检测, 如果错误就忽视
			errMsg, noFailed, fileLinesMark = check(line, count)
			todo(line, count, fileLinesMark)
		}
		if !noFailed {
			fmt.Println("Warning:")
			fmt.Println(errors.New(errMsg))
		}
	}
}

func Write(filename string, data string) {
	ioutil.WriteFile(filename, []byte(data), 0664)
}