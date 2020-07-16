package controller

import (
	"goexamer/config"
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

// 进行测试
func Exam() {
	var pItemQus, pItemAns, pStart, pFinish *router.State
	var pSelectBatch, pNo, pYes func()
	var curState *router.State

	service.HelpMsg()
	service.Title()

	ioTrigger.ReadInput(func(msg *trigger.Msg, exit *bool) {
		pNo = func() {}

		pYes = func() {
			service.SelectYes()
		}

		pSelectBatch = func() {
			service.Start(service.NewSelector(store.GetBatch(msg.Ctx)))
		}

		pFinish = router.NewState(func() {
			service.FinishMsg()
		}, func(input interface{}) {
			switch input.(int) {
			case views.SelectYes:
				service.Restart()
				curState = pItemQus
			case views.SelectNo:
				curState = pStart
			case views.SelectPost:
				pSelectBatch()
				curState = pItemQus
			}
		})

		pStart = router.NewState(func() {
			service.HelpMsg()
		}, func(input interface{}) {
			switch input.(int) {
			case views.SelectPost:
				pSelectBatch()
				curState = pItemQus
			default:
				curState = pStart
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
			case views.SelectPost:
				pSelectBatch()
				curState = pItemQus
			default:
				curState = pItemAns
			}
		})

		pItemQus = router.NewState(func() {
			service.ItemQus()
		}, func(input interface{}) {
			switch input.(int) {
			case views.SelectPost:
				pSelectBatch()
				curState = pItemQus
			default:
				curState = pItemAns
			}
		})

		if curState == nil {
			curState = pStart
		}
		curState.ChangeState(msg.Flag)
		curState.Todo()
	})
}