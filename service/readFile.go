package service

import (
	"errors"
	"goexamer/io"
	"goexamer/store"
	"goexamer/utils"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	rule *regexp.Regexp
	lastIndex string // 上一次读取的index
	lastBatch *store.Batch // 上一次的batch
)

type LineInfo struct {
	Line string
	Mark rune
	N    int
}

func init() {
	var err error
	rule, err = regexp.Compile(utils.FormatRule)
	if err != nil {
		panic(err)
	}
}

func SaveTitle(line string) {
	if !strings.HasPrefix(line,"title:") {
		panic(errors.New("unknown error " + line))
	}
	lastIndex = "title"
	store.SetTitle(utils.GetTitle(line))
}

func ReadLineOfTitle(line string) {
	if len(lastIndex) <= 0 || lastIndex != "title" {
		panic(errors.New("unknown error " + line))
	}
	store.SetTitle(line)
}

func ReadLineOfItem(line string) {
	if len(lastIndex) <= 0 || lastIndex == "title"  {
		panic(errors.New("unknown error " + line))
	}
	if lastBatch == nil {
		store.SaveQus(lastIndex, line)
	} else {
		lastBatch.SaveQus(lastIndex, line)
	}
}

func ReadBatchLine(line string) {
	if lastBatch == nil  {
		panic(errors.New("unknown error " + line))
	}
	lastBatch.AppendLine(line)
}

func NewBatch(line string) {
	if rule, _ := regexp.Compile(utils.Batch); !rule.MatchString(line) {
		panic(errors.New("unknown error " + line))
	}
	lastBatch = store.CreateBatch(line[1:len(line)-1])
	store.SaveBatch(lastBatch)
}

func ReadItem(line string) {
	// 每一行可以看做为一个pair
	if line != "" {
		utils.GetPair(line, func(str string) bool {
			if str[0] == utils.ItemPrefix[0] {
				return true
			}
			return false
		})(func(pair [2]string) {
			if lastBatch == nil {
				store.SaveQus(pair[0], pair[1])
			} else {
				lastBatch.SaveQus(pair[0], pair[1])
			}
			lastIndex = pair[0]
		})
	}
}

func ReadFile(fileName string, controllerCallBack func(info *LineInfo)) string {
	store.SetFileName(fileName)
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var errMsg string	// 去除空格后的mark, 错误信息
	var markString string // 解析完成最后的fileLinesMark
	var titleSet = false
	var info = &LineInfo{}
	io.EachLine(file)(func(line string, n int, fileLinesMark rune){
		if fileLinesMark == utils.TitleMark {
			titleSet = true
		}
		if titleSet {
			if fileLinesMark != utils.LineMark {
				line = strings.TrimSpace(line)
			}
			if fileLinesMark != utils.EmptyMark {
				markString += string(fileLinesMark)
			}
			markCheck := rule.MatchString(markString)
			if markCheck {
				errMsg = ""
			} else {
				errMsg += "Error: line[" + strconv.Itoa(n) + "] " + line + ", mark: " + markString + "\n"
			}
			// 获取当前信息
			if markString != "" {
				// 过滤空行
				if fileLinesMark != utils.EmptyMark {
					info.Mark = fileLinesMark
					info.Line = line
					info.N = n
					// 状态机切换
					controllerCallBack(info)
				}
			}
		}else{
			store.AppendBeforeTitle(line + "\n")
		}
	})

	// 错误信息
	if errMsg != "" {
		panic(errors.New(errMsg))
	}

	// 回收全局数据
	lastIndex = ""
	lastBatch = nil
	return ""
}
