package store

import (
	"github.com/pkg/errors"
)

var title []string			// 多行标题
type Batch struct {
	store map[string][]string // 测试项目, 保存每一行
	score map[string]int    // 剩余需要进行的测试次数
	Name  string
}

var (
	batchGroup map[string]*Batch
)

func init() {
	batchGroup = make(map[string]*Batch)
	batchGroup[""] = &Batch{
		make(map[string][]string),
		make(map[string]int),
		"",
	}
}

// 创建已经存在name的将会获得原来就有的
func CreateBatch(name string) *Batch {
	if _, ok := batchGroup[name]; ok {
		return batchGroup[name]
	}
	return &Batch{
		make(map[string][]string),
		make(map[string]int),
		name,
	}
}

func SaveBatch(batch *Batch) {
	if _, ok := batchGroup[batch.Name]; !ok {
		batchGroup[batch.Name] = batch
	}
}

func GetAnonBatch() *Batch {
	return batchGroup[""]
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

func (b *Batch) SetScore(qus string, n int) {
	if n < 0 {
		n = 0
	}
	b.score[qus] = n
}

func SetScore(qus string, n int) {
	if n < 0 {
		n = 0
	}
	batchGroup[""].score[qus] = n
}

func (b *Batch) GetScore(qus string) int {
	return b.score[qus]
}

func GetScore(qus string) int {
	return batchGroup[""].score[qus]
}

func (b *Batch) GetAllScore() map[string]int {
	return b.score
}

func GetAllScore() map[string]int {
	return batchGroup[""].score
}

func (b *Batch) SaveQus(qus string, ans string) {
	b.store[qus] = append(b.store[qus], ans)
}

func SaveQus(qus string, ans string) {
	batchGroup[""].store[qus] = append(batchGroup[""].store[qus], ans)
}

func (b *Batch) GetAll() map[string][]string {
	return b.store
}

func GetAll() map[string][]string {
	return batchGroup[""].store
}

func (b *Batch) GetQus(qus string) []string {
	return b.store[qus]
}

func GetQus(qus string) []string {
	return batchGroup[""].store[qus]
}

func (b *Batch) ToString() (str string) {
	str = "[" + b.Name + "]\n"
	for k, v := range b.store {
		s := ""
		for _, line := range v {
			s += line + "\n"
		}
		str += k + ":" + s
	}
	return
}

func ToString() (str string) {
	for k, v := range batchGroup[""].store {
		s := ""
		for _, line := range v {
			s += line + "\n"
		}
		str += k + ":" + s
	}
	return
}

