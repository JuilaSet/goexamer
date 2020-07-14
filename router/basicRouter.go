package router

type RouteFunc func()func(next func(string))
type NextFunc func(string)

func BuildRouteFunc(process func() string) RouteFunc {
	return func()func(next func(string)) {
		// 处理过程继续接收输入: 先跳转
		lastStr := process()
		return func(next func(string)){
			// 进入下一个状态: 再执行
			next(lastStr)
		}
	}
}

// example
//func Handler(){
//	// 过程
//	P1 := BuildRouteFunc(func() {
//		fmt.Println("P1 process")
//	})
//	P2 := BuildRouteFunc(func() {
//		fmt.Println("R2 process")
//	})
//
//	// 路由
//	var r1Next, r2Next NextFunc
//	r1Next = func(path string){
//		P1(func()string{ return io.InputReader() })(r2Next)
//	}
//	r2Next = func(path string){
//		P2(func()string{ return io.InputReader() })(r1Next)
//	}
//
//	// 启动
//	r1Next("")
//}
