package trigger

import (
	"goexamer/views"
)

// GUI事件输入触发器
type GUITrigger struct {}

func NewGUITrigger() *GUITrigger {
	return &GUITrigger{}
}

func (*GUITrigger) ReadInput(actionCallBack func(msg *Msg, exit *bool)) {
	msg := new(Msg)
	exit := false
	for !exit {
		msg.Flag, msg.Ctx = views.GetCommunicator().Receive()
		actionCallBack(msg, &exit)
	}
}

// 判断是否错误
func (*GUITrigger) Wait() {
	views.GetCommunicator().Receive()
}

// 判断是否错误
func (*GUITrigger) Judge() (b bool) {
	views.SetText("\n")
	flag, _ := views.GetCommunicator().Receive()
	switch flag {
	case views.SelectYes:
		b = true
	case views.SelectNo:
		b = false
	}
	return
}

func (*GUITrigger) Init() {}

