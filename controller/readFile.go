package controller

import (
	"github.com/pkg/errors"
	"goexamer/router"
	"goexamer/service"
	"goexamer/store"
	"goexamer/utils"
	"strconv"
)

func ReadFile(fileName string){
	var pStart, pNewBatch, pReadLineOfTitle, pReadLineOfBatch, pReadLineOfItem, pSetTitle, pReadItem *router.State
	var curState *router.State

	store.Init()
	service.ReadFile(fileName, func(info *service.LineInfo) {
		pStart = router.NewState(func() {}, func(input interface{}) {
			switch rune(input.(rune)) {
			case utils.TitleMark:
				curState = pSetTitle
			default:
				// 直到读到title为止
				curState = pStart
				//panic(errors.New("line[" + strconv.Itoa(info.N) + "] Need batch, title or item here"))
			}
		})
		pReadItem = router.NewState(func() {
			service.ReadItem(info.Line)
		}, func(input interface{}) {
			switch info.Mark {
			case utils.LineMark:
				curState = pReadLineOfItem
			case utils.BatchMark:
				curState = pNewBatch
			case utils.ItemMark:
				curState = pReadItem
			default:
				panic(errors.New("line[" + strconv.Itoa(info.N) + "] Need batch or 'next line' here"))
			}
		})
		pReadLineOfItem = router.NewState(func() {
			service.ReadLineOfItem(info.Line)
		}, func(input interface{}) {
			switch info.Mark {
			case utils.LineMark:
				curState = pReadLineOfItem
			case utils.BatchMark:
				curState = pNewBatch
			case utils.ItemMark:
				curState = pReadItem
			default:
				panic(errors.New("line[" + strconv.Itoa(info.N) + "] Need batch, item or 'next line' here"))
			}
		})
		pSetTitle = router.NewState(func() {
			service.SaveTitle(info.Line)
		}, func(input interface{}) {
			switch info.Mark {
			case utils.LineMark:
				curState = pReadLineOfTitle
			case utils.BatchMark:
				curState = pNewBatch
			case utils.ItemMark:
				curState = pReadItem
			default:
				panic(errors.New("line[" + strconv.Itoa(info.N) + "] Need batch, item or 'next line' here"))
			}
		})
		pReadLineOfTitle = router.NewState(func() {
			service.ReadLineOfTitle(info.Line)
		}, func(input interface{}) {
			switch info.Mark {
			case utils.LineMark:
				curState = pReadLineOfTitle
			case utils.BatchMark:
				curState = pNewBatch
			case utils.ItemMark:
				curState = pReadItem
			default:
				panic(errors.New("line[" + strconv.Itoa(info.N) + "] Need batch, item or 'next line' here"))
			}
		})
		pReadLineOfBatch = router.NewState(func() {
			service.ReadBatchLine(info.Line)
		}, func(input interface{}) {
			switch info.Mark {
			case utils.ItemMark:
				curState = pReadItem
			case utils.LineMark:
				curState = pReadLineOfBatch
			default:
				panic(errors.New("line[" + strconv.Itoa(info.N) + "] Need item or 'next line' here"))
			}
		})
		pNewBatch = router.NewState(func() {
			service.NewBatch(info.Line)
		}, func(input interface{}) {
			switch info.Mark {
			case utils.ItemMark:
				curState = pReadItem
			case utils.LineMark:
				curState = pReadLineOfBatch
			default:
				panic(errors.New("line[" + strconv.Itoa(info.N) + "] Need item or 'next line' here"))
			}
		})
		if curState == nil {
			curState = pStart
		}
		curState.ChangeState(info.Mark)
		curState.Todo()
	})
}
