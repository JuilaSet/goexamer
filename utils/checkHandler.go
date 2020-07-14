package utils

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	Batch = `^(\[.*?\])`
	Item = `^(#.+:)`
	Line = `^(\\.+)`
	Title = `^title:.+`
)

const (
	BatchMark = 'B'
	ItemMark = 'I'
	LineMark = 'n'
	TitleMark = 'T'
	ErrorMark = ','
	EmptyMark = ','
)

// 语法检测
func CheckFileHandler() func(line string, n int) (errMsg string, noFailed bool, fileLinesMark rune) {
	noFailed := true
	errMsg := ""
	var fileLinesMark rune
	ruleStrArr := []string{Batch, Item, Line, Title}
	return func(line string, n int) (string, bool, rune) {
		if line == "" {
			fileLinesMark = EmptyMark
		}else {
			matchFailed := true
			for _, rule := range ruleStrArr {
				if regexp.MustCompile(rule).MatchString(line) {
					switch rule {
					case Batch:
						fileLinesMark = BatchMark
					case Item:
						fileLinesMark = ItemMark
					case Line:
						fileLinesMark = LineMark
					case Title:
						fileLinesMark = TitleMark
					}
					matchFailed = false
				}
			}
			if matchFailed {
				fileLinesMark = ErrorMark
				sb := &strings.Builder{}
				fmt.Fprintf(sb, "syntax error line [%v]: %v\n", n, line)
				errMsg += sb.String()
				noFailed = false
			}
		}
		return errMsg, noFailed, fileLinesMark
	}
}