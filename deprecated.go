package main
//
//import (
//	"goexamer/config"
//	"goexamer/io"
//	"goexamer/store"
//	"goexamer/trigger"
//	"goexamer/utils"
//	"goexamer/views"
//	"strconv"
//	"strings"
//	"time"
//)
//
//var ioTrigger trigger.Trigger // 触发器
//var output io.OutPutter // 输出器
//
//// 命令函数
//var afterActionFuncMap map[string]func(value string)
//var midActionFuncMap map[string]func(value string)
//var beforeActionFuncMap map[string]func(value string)
//
//// 状态对象
//var (
//	curQus string // 当前执行对象, itemQus更新
//	totalCount, curCount int // 总共需要复习多少个, 当前已完成, itemQus更新
//	batch *store.Batch	// 当前batch
//)
//
//func init() {
//	afterActionFuncMap = make(map[string]func(string))
//	midActionFuncMap = make(map[string]func(string))
//	beforeActionFuncMap = make(map[string]func(string))
//	totalCount, curCount = 0, 0
//
//	ioTrigger = config.IoTrigger()
//	output = config.OutPutter()
//}
//
//func setBatch(name string) {
//	batch = store.GetBatch(name)
//}
//
//func GetBatch() *store.Batch {
//	return batch
//}
//
//// 注册action: 在用户交互之后执行
//func initAfterAction(){
//	scoreSet := batch.GetAllScore()
//	afterActionFuncMap["deduct"] = func(value string){
//		value = strings.ReplaceAll(value, utils.CurQusPrefix, curQus)
//		if n, ok := scoreSet[value]; ok {
//			batch.SetScore(value, n - 1)
//		}
//	}
//	afterActionFuncMap["mark"] = func(value string){
//		value = strings.ReplaceAll(value, utils.CurQusPrefix, curQus)
//		if n, ok := scoreSet[value]; ok {
//			batch.SetScore(value, n + 1)
//		}
//	}
//	// 强制跳转并执行某项
//	afterActionFuncMap["jmp"] = func(value string){
//		value = strings.ReplaceAll(value, utils.CurQusPrefix, curQus)
//		if _, ok := scoreSet[value]; ok {
//			batch.SetFinish(value, false)
//			Process(value, 1)
//		}
//	}
//	// 设置下一项, 如果下一项已经执行过, 就不执行
//	afterActionFuncMap["link"] = func(value string){
//		value = strings.ReplaceAll(value, utils.CurQusPrefix, curQus)
//		if n, ok := scoreSet[value]; ok {
//			Process(value, n)
//		}
//	}
//	// 设置当前item的测试次数为大于0的数
//	afterActionFuncMap["set"] = func(value string){
//		if n, err := strconv.Atoi(value); err != nil {
//			panic(err)
//		} else {
//			if _, ok := batch.GetAllScore()[curQus];ok {
//				if  n < 0 {
//					n = 0
//				}
//				batch.SetScore(curQus, n)
//			}
//		}
//	}
//	// 改变当前action的执行对象, 只针对在交互之后执行的action有效
//	afterActionFuncMap["execute"] = func(value string){
//		if _, ok := scoreSet[value]; ok {
//			curQus = value
//		}
//	}
//	// 显示当前对象
//	afterActionFuncMap["current"] = func(string){
//		output.Println("current item name:", curQus)
//	}
//}
//
//// 注册action: 在用户交互之前执行
//func initMidAction(){
//	midActionFuncMap["img"] = func(value string){
//		go func(value string) {
//			value = strings.ReplaceAll(value, utils.CurQusPrefix, curQus)
//			showImageFunc("./img", value, value)()
//		}(value)
//		time.Sleep(600 * time.Millisecond)
//	}
//}
//
//// 注册action: 在显示问题时执行
//func initBeforeAction() {
//	beforeActionFuncMap["qusImg"] = func(value string){
//		value = strings.ReplaceAll(value, utils.CurQusPrefix, curQus)
//		go func(value string) {
//			showImageFunc("./img", value, value)()
//		}(value)
//		time.Sleep(600 * time.Millisecond)
//	}
//	beforeActionFuncMap["ext"] = func(value string){
//		output.Println(value)
//	}
//	// 帮助
//	beforeActionFuncMap["help"] = func(value string){
//		switch value {
//		case "deduct":
//			output.Println("@deduct:item name", "item的测试次数-1")
//		case "mark":
//			output.Println("@mark:item name", "item的测试次数+1")
//		case "jmp":
//			output.Println("@jmp:item name", "当前结束后强制跳转并执行item")
//		case "link":
//			output.Println("@link:item name", "当前结束后进入指定的item, 如果该item已经执行过, 就不执行")
//		case "set":
//			output.Println("@set:count", "设置当前item的测试次数为大于0的数")
//		case "showImg":
//			output.Println("@showImg:image name", "显示当前img文件夹下的一张图片")
//		case "execute":
//			output.Println("@execute:item name", "改变当前action的执行对象为指定对象")
//		case "img":
//			output.Println("@img:image name", "显示当前img文件夹下的一张图片, 在交互之前进行")
//		case "help":
//			output.Println("@help[:action name]", "显示帮助")
//		default:
//			output.Println("@deduct:item name", "item的测试次数-1")
//			output.Println("@mark:item name", "item的测试次数+1")
//			output.Println("@jmp:item name", "当前结束后强制跳转并执行指定的item")
//			output.Println("@link:item name", "当前结束后进入指定的item, 如果该item已经执行过, 就不执行")
//			output.Println("@set:count", "设置当前item的测试次数为大于0的数")
//			output.Println("@showImg:image name", "显示当前img文件夹下的一张图片, 在交互之后进行")
//			output.Println("@execute:item name", "改变当前action的执行对象为指定对象, 只针对在交互之后执行的action有效")
//			output.Println("@img:image name", "显示当前img文件夹下的一张图片, 在交互之前进行")
//			output.Println("@help[:action name]", "显示帮助")
//		}
//	}
//}
//
//// 某个item的每一行回调lineCallBack, 检测动作
//func lineActionAnalyzeFunc(line string) (mid, after  func()) {
//	mid, after = func(){}, func(){}
//	// 以@开头的是特殊动作
//	if strings.HasPrefix(line, utils.ActionPrefix) {
//		action, value := utils.GetActionStr(line)
//		if amFunc, ok := beforeActionFuncMap[action]; ok {
//			amFunc(value)
//		}
//		if midFunc, ok := midActionFuncMap[action]; ok {
//			mid = func(){
//				midFunc(value)
//			}
//		}
//		if aFunc, ok := afterActionFuncMap[action]; ok {
//			after = func(){
//				aFunc(value)
//			}
//		}
//	}
//	return
//}
//
//// 重新计算当前进度
//func reCalcProcess() {
//	allScore := batch.GetAllScore()
//	curCount = 0
//	totalCount = len(allScore)
//	for _, score := range allScore {
//		if score < 1 {
//			curCount++
//		}
//	}
//}
//
//// 处理每一个item, 设置curQus为当前item, 标记为完成状态
//func itemProcess(qus string, ans []string, lineCallBack func(string) (mid, after func())) {
//	// 设置curQus为当前item
//	curQus = qus
//	reCalcProcess()
//
//	// 显示问题
//	output.Println("(" + strconv.Itoa(curCount) + "/" + strconv.Itoa(totalCount) + ")question^" + strconv.Itoa(batch.GetScore(qus)) + ":", qus)
//
//	// 执行函数
//	var afterActionFuncArray, midActionFuncArray []func()
//	var ansStr = "ans:" + func() (str string) {
//		for _, line := range ans {
//			if !strings.HasPrefix(line, "@") {
//				str += line + "\n"
//			}
//			mid, after := lineCallBack(line)
//			midActionFuncArray = append(midActionFuncArray, mid)
//			afterActionFuncArray = append(afterActionFuncArray, after)
//		}
//		return
//	}()
//	ioTrigger.Wait()
//
//	// 输出答案
//	output.Println(ansStr)
//	for _, aFunc := range midActionFuncArray {
//		aFunc()
//	}
//
//	output.Print("(y/N)-> ")
//	if ioTrigger.Judge() {
//		batch.SetScore(qus, batch.GetScore(qus) - 1)
//		output.Println("√\n")
//	} else {
//		output.Println("×\n")
//	}
//	output.Clear()
//
//	for _, aFunc := range afterActionFuncArray {
//		aFunc()
//	}
//
//	// 标记为完成
//	batch.SetFinish(qus, true)
//}
//
//// 显示图片
//func showImageFunc(rootDirPath, imageName string, title string) func() {
//	return views.ShowImage(views.FromImage(rootDirPath, imageName), title)
//}
//
//// 得到中间和之后执行的命令并执行before命令
//func getMidAndAfter(orders []string, callback func(mid, after func(), line string)) (str string) {
//	for _, line := range orders {
//		mid, after := lineActionAnalyzeFunc(line)
//		callback(mid, after, line)
//	}
//	return
//}
//
//// 执行标题
//func Title() {
//	// 设置curQus为当前item
//	curQus = ""
//
//	var actionFuncArray []func()
//	output.SetTitle(func() (str string) {
//		getMidAndAfter(store.GetTitle(), func(mid, after func(), line string){
//			str += line + "\n"
//			actionFuncArray = append(actionFuncArray, mid)
//			actionFuncArray = append(actionFuncArray, after)
//		})
//		return
//	}())
//
//	if batch.Qus != "" {
//		output.Println("======================")
//		output.Println(" chapter: ", batch.Qus)
//		output.Println("======================", "\n")
//	}
//
//	// title不分前后计算action
//	for _, aFunc := range actionFuncArray {
//		aFunc()
//	}
//}
//
//// 初始化Review数据
//func InitReviewData() {
//	// 初始化数据
//	batch.Init()
//}
//
//// 复习
//func Review() {
//	for qus, n := range batch.GetAllScore() {
//		Process(qus, n)
//	}
//	// 恢复访问标记, 为下一轮做准备数据
//	for qus := range batch.GetAllScore() {
//		batch.SetFinish(qus, false)
//	}
//}
//
//// 单个项目测试过程
//func Process(qus string, n int) {
//	// 剩余 > 0并且未完成才进行测试
//	if !batch.IsFinish(qus) && n > 0 {
//		itemProcess(qus, batch.GetQus(qus), lineActionAnalyzeFunc)
//	}
//}
//
//// 检测是否还有错题
//func HasReviewItem() (b bool) {
//	// 检测是否还有错题
//	b = false
//	for _, n := range GetBatch().GetAllScore() {
//		if n > 0 {
//			b = true
//			return
//		}
//	}
//	return
//}
//
//// 打印错误情况
////func PrintScore() string {
////	scores := batch.GetAllScore()
////	for qus, n := range scores {
////		output.Println("question: ", qus, ", wrong: [", n, "]")
////	}
////	output.Println()
////	return ""
////}
//
//// 初始化
//func SelectBatch(batchName string) {
//	setBatch(batchName)
//	initBeforeAction()
//	initMidAction()
//	initAfterAction()
//}
//
