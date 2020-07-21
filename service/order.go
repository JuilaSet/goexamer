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
	beforeActionFuncMap["def"] = func(selector *Selector, params []string){
		name, value := params[0], params[1]
		selector.SetTempVar(name, value, true, false)
	}
	beforeActionFuncMap["defReg"] = func(selector *Selector, params []string){
		name, value := params[0], params[1]
		selector.SetTempVar(name, value, true, true)
	}
	beforeActionFuncMap["appDef"] = func(selector *Selector, params []string){
		name, value := params[0], params[1]
		selector.SetTempVar(name, value, false, false)
	}
	beforeActionFuncMap["undef"] = func(selector *Selector, params []string){
		name := params[0]
		selector.RemoveTempVar(name)
	}
}

func InitMidActionFunc() {
	midActionFuncMap["qus"] = func(selector *Selector, params []string) {
		var lines []string
		for _, v := range params {
			lines = append(lines, selector.ReplaceStringAccordingToTempVar(v))
		}
		output.Println(strings.Join(lines, ":"))
	}
	midActionFuncMap["ext"] = func(selector *Selector, params []string) {
		var lines []string
		for _, v := range params {
			lines = append(lines, selector.ReplaceStringAccordingToTempVar(v))
		}
		output.Println(strings.Join(lines, ":"))
	}
}

func InitAfterActionFunc() {
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
	afterActionFuncMap["ext"] = func(selector *Selector, params []string) {
		var lines []string
		for _, v := range params {
			lines = append(lines, selector.ReplaceStringAccordingToTempVar(v))
		}
		output.Println(strings.Join(lines, ":"))
	}
	afterActionFuncMap["line"] = func(selector *Selector, params []string) {
		var lines []string
		for _, v := range params {
			lines = append(lines, selector.ReplaceStringAccordingToTempVar(v))
		}
		output.Println(strings.Join(lines, ""))
	}
}