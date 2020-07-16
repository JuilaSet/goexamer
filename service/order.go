package service

import (
	"goexamer/utils"
	"strconv"
	"strings"
	"time"
)

// 命令函数
var beforeActionFuncMap map[string]ActionFunc
var midActionFuncMap map[string]ActionFunc
var afterActionFuncMap map[string]ActionFunc

func init(){
	beforeActionFuncMap = make(map[string]ActionFunc)
	midActionFuncMap = make(map[string]ActionFunc)
	afterActionFuncMap = make(map[string]ActionFunc)
}

func init(){
	InitBeforeActionFunc()
	InitMidActionFunc()
	InitAfterActionFunc()
}

func InitBeforeActionFunc(){
	beforeActionFuncMap["set"] = func(selector *Selector, params []string){
		switch {
		case len(params) < 2:
			item, _ := selector.NextItem()
			if n, err := strconv.Atoi(params[0]); err != nil {
				panic(err)
			} else {
				selector.SetScore(item.Qus, n)
			}
		default:
			name, value := params[0], params[1]
			if n, err := strconv.Atoi(value); err != nil {
				panic(err)
			} else {
				item, _ := selector.NextItem()
				name = strings.ReplaceAll(name, utils.CurQusPrefix, item.Qus)
				selector.SetScore(name, n)
			}
		}
	}
	beforeActionFuncMap["deduct"] = func(selector *Selector, params []string){
		value := params[0]
		item, _ := selector.NextItem()
		value = strings.ReplaceAll(value, utils.CurQusPrefix, item.Qus)
		selector.DeductItem(value, 1)
	}
	beforeActionFuncMap["mark"] = func(selector *Selector, params []string){
		value := params[0]
		item, _ := selector.NextItem()
		value = strings.ReplaceAll(value, utils.CurQusPrefix, item.Qus)
		selector.DeductItem(value, -1)
	}
}

func InitMidActionFunc() {
	midActionFuncMap["qus"] = func(selector *Selector, params []string) {
		for _, v := range params {
			output.Println(v)
		}
	}
}

func InitAfterActionFunc() {
	// 设置下一项, 如果下一项可以执行就在下次执行
	afterActionFuncMap["link"] = func(selector *Selector, params []string) {
		value := params[0]
		item := selector.CurItem()
		value = strings.ReplaceAll(value, utils.CurQusPrefix, item.Qus)
		selector.SetNext(item.Qus)
	}
	afterActionFuncMap["jmp"] = func(selector *Selector, params []string){
		value := params[0]
		value = strings.ReplaceAll(value, utils.CurQusPrefix, selector.CurItem().Qus)
		selector.SetJmp(value)
	}
	afterActionFuncMap["img"] = func(selector *Selector, params []string) {
		item := selector.CurItem()
		for _, value := range params {
			go func(imgName string) {
				value = strings.ReplaceAll(value, utils.CurQusPrefix, item.Qus)
				ShowImageFunc("./img", value, value)()
			}(value)
			time.Sleep(600 * time.Millisecond)
		}
	}
}