package controller

import (
	"goexamer/config"
	"goexamer/io"
	"goexamer/router"
	"goexamer/service"
	"goexamer/store"
	"goexamer/trigger"
)

var ioTrigger trigger.Trigger // 触发器
var output io.OutPutter // 输出器

func init(){
	ioTrigger = config.IoTrigger()
	output = config.OutPutter()
}

// 进行测试
func Exam() {
	// process
	pLogo := router.BuildRouteFunc(service.Logo)
	pExam := router.BuildRouteFunc(service.Exam)
	pReview := router.BuildRouteFunc(func() string {
		service.Review()
		return ""
	})

	// router
	var nextLogo, nextExam, nextReview router.NextFunc
	nextLogo = func(lastStr string){
		next := pLogo()
		next(nextExam)
	}
	nextExam = func(string){
		next := pExam()
		// 检测是否还有错题
		for _, n := range service.GetBatch().GetAllScore() {
			if n > 0 {
				output.Print("Review? (y/N)-> ")
				ioTrigger.Judge(func(b bool) {
					output.Clear()
					if b {
						next(nextReview)
					} else {
						next(nextLogo)
					}
				})
				return
			}
		}
	}
	nextReview = func(string){
		next := pReview()
		// 检测是否还有错题
		for _, n := range service.GetBatch().GetAllScore() {
			if n > 0 {
				output.Print("Continue? (y/N)-> ")
				ioTrigger.Judge(func(b bool) {
					output.Clear()
					if b {
						next(nextReview)
					}
				})
				return
			}
		}
	}
	// start
	for {
		for _, batch := range store.GetAllBatch() {
			service.Init(batch.Name)
			nextLogo("")
		}
		output.Print("Again? (y/N)-> ")
		var bb bool
		ioTrigger.Judge(func(b bool) {
			output.Clear()
			bb = b
		})
		if !bb {
			break
		}
	}
}