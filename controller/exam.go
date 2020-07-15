package controller

import (
	"goexamer/config"
	"goexamer/io"
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
	// router
	var nextExam, nextReview func()
	nextExam = func(){
		output.Clear()
		service.Exam()
		// 检测是否还有错题
		for _, n := range service.GetBatch().GetAllScore() {
			if n > 0 {
				output.Clear()
				if output.Print("Review? (y/N)-> "); ioTrigger.Judge() {
					nextReview()
				}
				return
			}
		}
	}
	nextReview = func(){
		output.Clear()
		service.Review()
		// 检测是否还有错题
		for _, n := range service.GetBatch().GetAllScore() {
			if n > 0 {
				output.Clear()
				output.Print("Continue? (y/N)-> ")
				if ioTrigger.Judge() {
					nextReview()
				}
				return
			}
		}
	}
	// start
	for {
		for _, batch := range store.GetAllBatch() {
			service.Init(batch.Name)
			nextExam()
		}
		output.Clear()
		if output.Print("Again? (y/N)-> "); !ioTrigger.Judge() {
			break
		}
	}
}