package service

import (
	"goexamer/config"
	"goexamer/io"
	"goexamer/store"
	"goexamer/views"
	"strconv"
)

var output io.OutPutter // 输出器
var selector *Selector // 调度器
var batchMsg string

func init() {
	output = config.OutPutter()
}

// 显示图片
func ShowImageFunc(rootDirPath, imageName string, title string) func() {
	return views.ShowImage(views.FromImage(rootDirPath, imageName), title)
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

func Start(s *Selector) {
	selector = s
	selector.Init()
	output.Clear()
	Batch()
}

func Restart() {
	selector.Init()
	output.Clear()
}

func BatchName() {
	output.Println(batchMsg)
}

func Batch() {
	batch := selector.Batch()
	item := NewItem(batch.Name, batch.Lines())
	selector.SetCurItemDangerous(item)
	batchMsg = "======================\n" +
	" chapter: " + item.Qus + "\n"
	for _, line := range item.Ans {
		batchMsg += " " + line + "\n"
	}
	batchMsg += "======================" + "\n"
	selector.ExecuteBeforeFunc()
	selector.ExecuteMidFunc()
	selector.ExecuteAfterFunc()
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
	selector.ExecuteBeforeFunc()
	item, totalCount := selector.PopItem(), len(selector.Batch().GetAllQus())
	output.Println("(" + strconv.Itoa(selector.FinishCount()) + "/" + strconv.Itoa(totalCount) + ")question^" +
		strconv.Itoa(selector.ItemScore(item.Qus)) + ":", item.Qus)
	selector.ExecuteMidFunc()
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
	selector.ExecuteAfterFunc()
	output.Println(ansStr)
	output.Print("(y/N)-> ")
}

func SelectYes() {
	selector.Deduct(1)
}