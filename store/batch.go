package store

import (
	"github.com/pkg/errors"
	"sort"
)

var fileName string			// 当前文件名称
var beforeTitle string		// 设置在正式开始前的内容
var title []string			// 多行标题

type Batch struct {
	Name  string
	lines []string				// batch的每一行
	store map[string][]string	// 测试项目的每一行
}

var (
	batchGroup map[string]*Batch
	batchArray []*Batch
)

func Init() {
	batchGroup = make(map[string]*Batch)
	batchArray = nil
	title = make([]string, 0)
	beforeTitle = ""
	batchGroup[""] = CreateBatch("")
}

// 设置当前文件名称
func SetFileName(name string){
	fileName = name
}

// 设置当前文件名称
func CurFileName() string{
	return fileName
}

// 设置正式开始前的内容
func AppendBeforeTitle(b string) {
	beforeTitle += b
}

func GetBeforeTitle() string {
	return beforeTitle
}

// 变为数组
func BatchArray() []*Batch {
	if batchArray == nil {
		var nameArr []string
		for name := range GetAllBatch() {
			nameArr = append(nameArr, name)
		}
		sort.Strings(nameArr)
		for _, name := range nameArr {
			batchArray = append(batchArray, batchGroup[name])
		}
	}
	return batchArray
}

// 创建已经存在name的将会获得原来就有的
func CreateBatch(name string) *Batch {
	if _, ok := batchGroup[name]; ok {
		return batchGroup[name]
	}
	return &Batch{
		name,
		make([]string, 0),
		make(map[string][]string),
	}
}

func SaveBatch(batch *Batch) {
	if _, ok := batchGroup[batch.Name]; !ok {
		batchGroup[batch.Name] = batch
	}
}

func GetAllBatch() map[string]*Batch {
	return batchGroup
}

func GetBatch(name string) *Batch {
	if v, ok := batchGroup[name]; ok {
		return v
	} else {
		panic(errors.New("batch named " + name + " not found"))
	}
}

func SetTitle(str string) {
	title = append(title, str)
}

func GetTitle() []string {
	return title
}

func (b *Batch) WriteQus(qus string, ans []string) {
	b.store[qus] = ans
}

func (b *Batch) SaveQus(qus string, ans string) {
	b.store[qus] = append(b.store[qus], ans)
}

func SaveQus(qus string, ans string) {
	batchGroup[""].store[qus] = append(batchGroup[""].store[qus], ans)
}

func (b *Batch) GetAllQus() map[string][]string {
	return b.store
}

func GetAll() map[string][]string {
	return batchGroup[""].store
}

func (b *Batch) GetQus(qus string) []string {
	return b.store[qus]
}

func (b *Batch) AppendLine(line string) {
	b.lines = append(b.lines, line)
}

func (b *Batch) Lines() []string {
	return b.lines
}

func (b *Batch) IsEmpty() bool {
	return len(b.store) == 0
}

func (b *Batch) ToString() (str string) {
	if len(b.store) <= 0 {
		return ""
	}
	if b.Name != "" {
		str = "[" + b.Name + "]\n"
		for _, batchLine := range b.lines {
			str += batchLine + "\n"
		}
	} else {
		str = "[<default batch>]\n"
	}
	str += "\n"
	for qus, ans := range b.store {
		s := ""
		for _, line := range ans {
			s += line + "\n"
		}
		str += "#" + qus + ":" + s + "\n"
	}
	return
}


