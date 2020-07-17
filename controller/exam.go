package controller

import (
	"goexamer/config"
	"goexamer/params"
	"goexamer/router"
	"goexamer/service"
	"goexamer/store"
	"goexamer/trigger"
	"goexamer/views"
)

var ioTrigger trigger.Trigger // 触发器

func init(){
	ioTrigger = config.IoTrigger()
}

func pNo() {
	service.SelectNo()
}

func pYes() {
	service.SelectYes()
}


// 进行测试
func Exam() {
	var pItemQus, pItemAns, pStart, pFinish, pSelectFile *router.State
	var pSelectBatch func()
	var curState *router.State
	var lastFileName string

	ioTrigger.ReadInput(func(msg *trigger.Msg, exit *bool) {
		pSelectFile = router.NewState(func() {
			if msg.Ctx != "" {
				fileName := msg.Ctx
				ReadFile(fileName)
				lastFileName = fileName
				service.Title()
				service.HelpBatchMsg()
			} else if lastFileName != ""{
				ReadFile(lastFileName)
				service.HelpBatchMsg()
			}
		}, func(input interface{}) {
			switch input.(int) {
			case views.SelectBatch:
				pSelectBatch()
				curState = pItemQus
			}
		})

		pSelectBatch = func() {
			service.Start(service.NewSelector(store.GetBatch(msg.Ctx)))
		}

		pFinish = router.NewState(func() {
			service.FinishMsg()
		}, func(input interface{}) {
			switch input.(int) {
			case views.SelectYes:
				service.Init()
				curState = pItemQus
			case views.SelectNo:
				curState = pStart
			case views.SelectBatch:
				pSelectBatch()
				curState = pItemQus
			case views.SelectFile:
				curState = pSelectFile
			}
		})

		pStart = router.NewState(func() {
			service.HelpFileMsg()
		}, func(input interface{}) {
			switch input.(int) {
			case views.SelectFile:
				curState = pSelectFile
			}
		})

		pItemAns = router.NewState(func() {
			service.ItemAns()
		}, func(input interface{}) {
			switch input.(int) {
			case views.SelectYes:
				pYes()
				if service.IsEnd() {
					curState = pFinish
				} else {
					curState = pItemQus
				}
			case views.SelectNo:
				pNo()
				curState = pItemQus
			case views.SelectBatch:
				pSelectBatch()
				curState = pItemQus
			case views.SelectFile:
				curState = pSelectFile
			}
		})

		pItemQus = router.NewState(func() {
			service.ItemQus()
		}, func(input interface{}) {
			switch input.(int) {
			case views.SelectBatch:
				pSelectBatch()
				curState = pItemQus
			case views.SelectFile:
				curState = pSelectFile
			default:
				curState = pItemAns
			}
		})

		if curState == nil {
			if fileName := params.GetInputFileName(); fileName != "" {
				lastFileName = fileName
				curState = pSelectFile
			} else {
				curState = pStart
				service.HelpFileMsg()
			}
		}
		curState.ChangeState(msg.Flag)
		curState.Todo()
	})
}