package trigger

type Msg struct {
	Flag int
	Ctx string
}

type Trigger interface {
	ReadInput(actionCallBack func(msg *Msg, exit *bool))	// bool为false退出
	Init()	// 初始化配置
}
