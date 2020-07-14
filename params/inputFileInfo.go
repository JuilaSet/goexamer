package params

import "flag"

type InputFileInfo struct {
	Name string
}

func (f *InputFileInfo) Set(s string) error {
	f.Name = s
	return nil
}

func (f *InputFileInfo) String() string {
	return f.Name
}

var fileInfo InputFileInfo

func init(){
	// 输入文件的选择
	fileInfo = InputFileInfo{"exam.txt"}
	// 注册方法
	flag.CommandLine.Var(&fileInfo, "i", "input file's name")
	// 解析
	flag.Parse()
}

func GetInputFileName() string {
	return fileInfo.Name
}