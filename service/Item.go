package service

import (
	"goexamer/utils"
	"strings"
)
type ActionFunc func(selector *Selector, params []string)

type Action struct {
	Name string
	Param []string
	Func ActionFunc
}

type Item struct {
	Qus string
	ActionBefore, ActionMid, ActionAfter []Action
}

var (
	EofItem *Item
	NilItem *Item
)

func init() {
	NilItem = NewItem("<this batch is empty>", []string{"<nothing todo>"})
	EofItem = NewItem("<End>", []string{"<nothing todo>"})
}

func setFunc(item *Item, line string) {
	if strings.HasPrefix(line, utils.ActionPrefix) {
		action, param := utils.GetActionStr(line)
		if bFunc, ok := beforeActionFuncMap[action]; ok {
			item.ActionBefore = append(item.ActionBefore, Action{action, param, bFunc})
		}
		if midFunc, ok := midActionFuncMap[action]; ok {
			item.ActionMid = append(item.ActionMid, Action{action, param, midFunc})
		}
		if aFunc, ok := afterActionFuncMap[action]; ok {
			item.ActionAfter = append(item.ActionAfter, Action{action, param, aFunc})
		}
	} else {
		item.ActionAfter = append(item.ActionAfter, Action{"line", []string{line}, afterActionFuncMap["line"]})
	}
}

func NewItem(qus string, rawAns []string) (item *Item) {
	item = &Item{
		qus,
		make([]Action, 0),
		make([]Action, 0),
		make([]Action, 0),
	}
	for _, line := range rawAns {
		setFunc(item, line)
	}
	return
}
