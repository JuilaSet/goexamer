package io

import (
	"fmt"
	"goexamer/views"
	"os"
	"strconv"
)

type OutPutter interface {
	Clear()
	Close()
	SetTitle(title... string)
	Print(...interface{})
	Println(...interface{})
}

// 文件输出
type FileOutPutter struct {
	fileName string
	file *os.File
}

func NewFileOutPutter(fileName string) *FileOutPutter {
	fd, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil{
		panic(err)
	}
	return &FileOutPutter{fileName:fileName, file:fd}
}

func (o *FileOutPutter) Clear(){
	Write(o.fileName, "")
}

func (o *FileOutPutter) SetTitle(title... string) {
	titles := "title:"
	for _, tLine := range title {
		titles += tLine + "\n"
	}
	o.file.Write([]byte(titles))
}

func (o *FileOutPutter) Print(str ...interface{}){
	var msg string
	for _, s := range str {
		switch s.(type) {
		case int:
			msg += strconv.Itoa(s.(int))
		case string:
			msg += string(s.(string))
		}
	}
	o.file.Write([]byte(msg))
}

func (o *FileOutPutter) Println(str ...interface{}){
	var msg string
	for _, s := range str {
		switch s.(type) {
		case int:
			msg += strconv.Itoa(s.(int))
		case string:
			msg += string(s.(string))
		}
	}
	o.file.Write([]byte(msg + "\n"))
}

func (o *FileOutPutter) Close(){
	o.file.Close()
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

func (*ConsoleOutPutter) Close(){
	// 暂未实现
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

func (*GUIOutPutter) Close(){
	// 暂未实现
}