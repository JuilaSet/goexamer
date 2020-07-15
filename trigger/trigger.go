package trigger

type Msg struct {
	Flag int
	Ctx string
}

type Trigger interface {
	ReadInput(actionCallBack func(msg *Msg, exit *bool))	// bool为false退出

	// deprecated
	Wait() // 停止并等待输入
	Judge() bool // 判断输入是否为确定
	Init()	// 初始化配置
}
