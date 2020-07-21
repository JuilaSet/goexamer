package service

import (
	"goexamer/config"
	"goexamer/io"
	"goexamer/store"
	"goexamer/utils"
	"goexamer/views"
	"regexp"
	"strconv"
	"strings"
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

func HelpFileMsg() {
	output.Clear()
	output.Println(" Please select a file")
}

func HelpBatchMsg() {
	output.Clear()
	output.Println(" Please select a batch")
}

func Start(s *Selector) {
	selector = s
	RestartBatch()
}

func RestartBatch() {
	selector.Init()
	output.Clear()
	Batch()
}

func BatchName() {
	output.Println(batchMsg)
}

func Batch() {
	batch := selector.Batch()
	item := NewItem(batch.Name, batch.Lines())
	selector.ExecuteBeforeFuncFromItem(item)
	batchMsg = "======================\n" +
	" chapter: " + selector.ReplaceStringAccordingToTempVar(item.Qus) + "\n"
	for _, line := range item.ActionAfter {
		if line.Name == "line" {
			batchMsg += " " + line.Param[0] + "\n"
		}
	}
	batchMsg += "======================" + "\n"
	selector.ExecuteMidFuncFromItem(item)
	selector.ExecuteAfterFuncFromItem(item)
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

func RefreshItemQus() {
	selector.RefreshCurItem()
}

func ItemQus() {
	output.Clear()
	BatchName()
	// 选择调度系数最大的先执行
	selector.SetNext(selector.MinHotFactorItemName())
	selector.ExecuteBeforeFunc()
	item, totalCount := selector.PopItem(), len(selector.Batch().GetAllQus())
	output.Println("(" + strconv.Itoa(selector.FinishCount()) + "/" + strconv.Itoa(totalCount) + ")question^" +
		selector.DispatchCoefficientString(item.Qus) + ":" + selector.ReplaceStringAccordingToTempVar(item.Qus))
	selector.ExecuteMidFunc()
}

func ItemAnsRefresh() {
	output.Clear()
	BatchName()
	item, totalCount := selector.CurItem(), len(selector.Batch().GetAllQus())
	selector.ExecuteBeforeFuncFromItem(item)
	output.Println("(" + strconv.Itoa(selector.FinishCount()) + "/" + strconv.Itoa(totalCount) + ")question^" +
		selector.DispatchCoefficientString(item.Qus) + ":" + item.Qus)
	selector.ExecuteMidFunc()
	// 显示问题
	output.Println("ans:")
	selector.ExecuteAfterFunc()
	output.Print("(y/N)-> ")
}


func ItemAns() {
	// 显示问题
	output.Println("ans:")
	selector.ExecuteAfterFunc()
	output.Print("(y/N)-> ")
}

func SelectYes() {
	if i := selector.CalcDispatchCoefficient(1, true); selector.IsFinish(i) {
		selector.Deduct(1)
	}
}

func SelectNo() {
	selector.CalcDispatchCoefficient(0, false)
}

func EditItem() {
	views.FromItem(selector.Batch().Name).Run()
}

func SaveItem(itemString string) {
	check := utils.CheckLinesHandler()
	lines := strings.Split(itemString, "\n")
	var formatMarks []string
	// 语法检查
	for n, line := range lines {
		_, _, m := check(line, n)
		formatMarks = append(formatMarks, string(m))
	}
	if rule, _ := regexp.Compile(utils.ItemFormatRule); rule.MatchString(strings.Join(formatMarks, "")) {
		var qus string
		var ans []string
		for _, line := range lines {
			// 取出item header
			utils.GetPair(line, func(s string) bool {
				if s == "" {
					return false
				} else if strings.HasPrefix(s, utils.ItemPrefix) {
					return true
				} else {
					ans = append(ans, line)
					return false
				}
			})(func(pair [2]string) {
				qus = pair[0]
				ans = append(ans, pair[1])
			})
		}
		selector.AddNewItem(qus, ans)
	}
}