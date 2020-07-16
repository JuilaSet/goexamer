package main

import (
	"fmt"
	"goexamer/controller"
	"goexamer/views"
	"os"
	"runtime"
)

func main(){
	// 错误日志
	defer func(){
		if err := recover(); err != nil {
			runtime.Gosched()
			fmt.Println(err)
		}
	}()
	// 启动状态机
	go func(){
		views.Wait()
		controller.Init()
		controller.ReadFile()
		controller.Exam()
		os.Exit(0)
	}()
	// 启动视图进程
	views.Index()
}