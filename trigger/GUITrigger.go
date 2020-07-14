package trigger

import (
	"goexamer/views"
)

// GUI事件输入触发器
type GUITrigger struct {}

func NewGUITrigger() *GUITrigger {
	return &GUITrigger{}
}

// 判断是否错误
func (*GUITrigger) Wait() {
	<-views.GetYesOrNo()
}

// 判断是否错误
func (*GUITrigger) Judge(callback func(bool)) {
	views.SetText("\n")
	callback(<-views.GetYesOrNo())
}

func (*GUITrigger) Init() {}

