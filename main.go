package main

import (
	"fmt"
	"goexamer/config"
	"goexamer/controller"
	"goexamer/params"
	"goexamer/service"
	"goexamer/views"
	"os"
	"runtime"
)

func main(){
	// 错误日志
	defer func(){
		if err := recover(); err != nil {
			runtime.Gosched()
			config.OutPutter().Println(err)
		}
	}()
	// 启动状态机
	go func(){
		views.Wait()
		controller.Init()
		if fileName := params.GetInputFileName(); fileName != "" {
			fmt.Println("-i", fileName)
			controller.ReadFile(fileName)
			service.Title()
			service.HelpBatchMsg()
		}
		controller.Exam()
		os.Exit(0)
	}()
	// 启动视图进程
	views.Index()
}