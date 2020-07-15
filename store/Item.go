package store

import (
	"goexamer/utils"
	"strings"
)

type Action struct {
	Name, Value string
}

type Item struct {
	Qus string
	Ans []string
	Action []Action
}

var (
	EofItem *Item
	NilItem *Item
)

func init() {
	NilItem = NewItem("<this batch is empty>", []string{"<nothing todo>"})
	EofItem = NewItem("<End>", []string{"<nothing todo>"})
}

func NewItem(qus string, rawAns []string) *Item {
	action, ans := make([]Action, 0), make([]string, 0)
	for _, line := range rawAns {
		if strings.HasPrefix(line, utils.ActionPrefix) {
			// 特殊动作
			name, value := utils.GetActionStr(line)
			action = append(action, Action{name, value})
		} else {
			ans = append(ans, line)
		}
	}
	return &Item{
		qus,
		ans,
		action,
	}
}
