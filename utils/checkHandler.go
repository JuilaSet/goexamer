package utils

import (
	"regexp"
)

const (
	//regExp = `(^\s$)|(^Tn*((Bn*)?(In*)+)*$)`
	FormatRule = `(^\s$)|(^` +
		string(TitleMark) +
		string(LineMark) + `*((` +
		string(BatchMark) +
		string(LineMark) + `*)?(` +
		string(ItemMark) +
		string(LineMark) + `*)+)*$)`
)

const (
	Batch    = `^(\[.*?\])`
	Item     = `^(#.+:)`
	Title    = `^title:.*`
)

const (
	BatchMark    = 'B'
	ItemMark     = 'I'
	LineMark     = 'n'
	TitleMark    = 'T'
	EmptyMark    = ','
)

// 语法检测
func CheckFileHandler() func(line string, n int) (errMsg string, noFailed bool, fileLinesMark rune) {
	noFailed := true
	errMsg := ""
	var fileLinesMark rune
	ruleStrArr := []string{Batch, Item, Title}	// 匹配规则
	return func(line string, n int) (string, bool, rune) {
		if line == "" {
			fileLinesMark = EmptyMark
		} else {
			matchFailed := true
			for _, rule := range ruleStrArr {
				if regexp.MustCompile(rule).MatchString(line) {
					switch rule {
					case Batch:
						fileLinesMark = BatchMark
					case Item:
						fileLinesMark = ItemMark
					case Title:
						fileLinesMark = TitleMark
					}
					matchFailed = false
					break
				}
			}
			if matchFailed {
				fileLinesMark = LineMark
			}
		}
		return errMsg, noFailed, fileLinesMark
	}
}
