package config

import (
	"goexamer/io"
	"goexamer/trigger"
)

var ioTrigger trigger.Trigger // 触发器
var output io.OutPutter // 输出器

func IoTrigger() trigger.Trigger {
	return ioTrigger
}

func OutPutter() io.OutPutter {
	return output
}

func init() {
	//ioTrigger = trigger.NewConsoleTrigger() // 默认控制台触发器
	//output = io.NewConsoleOutPutter()	// 默认控制台输出
	ioTrigger = trigger.NewGUITrigger()
	output = io.NewGUIOutPutter()
}