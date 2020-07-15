package service

import (
	"goexamer/config"
	"goexamer/io"
	"goexamer/store"
	"strconv"
)

var output io.OutPutter // 输出器
var selector *store.Selector // 调度器

func init() {
	output = config.OutPutter()
}

func FinishMsg() {
	output.Clear()
	output.Println("Finish!")
	output.Println("Continue? (y/N)->")
}

func HelpMsg() {
	output.Clear()
	output.Println(" Please select a batch")
}

func Start(s *store.Selector) {
	selector = s
	selector.Init()
	output.Clear()
}

func BatchName() {
	batch := selector.Batch()
	if batch.Name != "" {
		output.Println("======================")
		output.Println(" chapter: ", batch.Name)
		output.Println("======================", "\n")
	}
}

func Title() {
	output.SetTitle(func() (str string) {
		for _, line := range store.GetTitle() {
			str += line + "\n"
		}
		return
	}())
}

func IsEnd() bool {
	return selector != nil && !selector.HasNext()
}

func ItemQus() {
	output.Clear()
	BatchName()
	item := selector.PopItem()
	totalCount := len(selector.Batch().GetAllQus())
	output.Println("(" + strconv.Itoa(selector.FinishCount()) + "/" + strconv.Itoa(totalCount) + ")question^" +
		strconv.Itoa(selector.ItemScore(item.Qus)) + ":", item.Qus)
}

func ItemAns() {
	// 显示问题
	item := selector.CurItem()
	var ansStr = "ans:" + func() (str string) {
		for _, line := range item.Ans {
			str += line + "\n"
		}
		return
	}()
	output.Println(ansStr)
	output.Print("(y/N)-> ")
}

func SelectYes() {
	selector.Deduct(1)
}