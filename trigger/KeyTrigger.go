package trigger

import (
	"errors"
	"fmt"
	"goexamer/io"
)

// 键盘输入触发器
type KeyTrigger struct {}

func NewConsoleTrigger() *KeyTrigger {
	return &KeyTrigger{}
}

// 判断是否错误
func (*KeyTrigger) Wait() {
	io.Wait()
}

// 判断是否错误
func (*KeyTrigger) Judge(callback func(r bool)) {
	io.ReadAndCompare("y", func(string){
		callback(true)
	}, func(string){
		callback(false)
	})
}

func (*KeyTrigger) Init() {
	setInput("y")
	testInput("y")
}

func setInput(str string) {
	fmt.Print("input key " + str + "-> ")
	io.SetInputString(str)
}

func testInput(str string) {
	fmt.Print("test key " + str + "-> ")
	if io.ReadInput() == io.GetInputString(str) {
		fmt.Println("Success!")
	} else {
		panic(errors.New("test failed"))
	}
}