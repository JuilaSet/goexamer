package io

import (
	"fmt"
	"goexamer/views"
	"strconv"
)

type OutPutter interface {
	Clear()
	SetTitle(title... string)
	Print(...interface{})
	Println(...interface{})
}

// 控制台输出
type ConsoleOutPutter struct {}

func NewConsoleOutPutter() *ConsoleOutPutter {
	return &ConsoleOutPutter{}
}

func (*ConsoleOutPutter) Clear(){
	// 暂未实现
}

func (*ConsoleOutPutter) SetTitle(title... string) {
	fmt.Print(title)
}

func (*ConsoleOutPutter) Print(str ...interface{}){
	fmt.Print(str...)
}

func (*ConsoleOutPutter) Println(str ...interface{}){
	fmt.Println(str...)
}

// GUI输出
type GUIOutPutter struct {}

func NewGUIOutPutter() *GUIOutPutter {
	return &GUIOutPutter{}
}

func (*GUIOutPutter) Clear() {
	views.Clear()
}

func (*GUIOutPutter) SetTitle(title... string) {
	views.SetTitle(title...)
}

func (*GUIOutPutter) Print(str ...interface{}){
	var msg string
	for _, s := range str {
		switch s.(type) {
		case int:
			msg += strconv.Itoa(s.(int))
		case string:
			msg += string(s.(string))
		}
	}
	views.SetText(msg)
}

func (*GUIOutPutter) Println(str ...interface{}){
	var msg string
	for _, s := range str {
		switch s.(type) {
		case int:
			msg += strconv.Itoa(s.(int)) + " "
		case string:
			msg += string(s.(string)) + " "
		}
	}
	views.SetText(msg + "\n")
}
