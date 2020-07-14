package trigger

type Trigger interface {
	Wait() // 停止并等待输入
	Judge(func(bool)) // 判断输入是否为确定
	Init()	// 初始化配置
}
