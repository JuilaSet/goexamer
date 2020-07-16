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

func (*GUITrigger) Init() {}

