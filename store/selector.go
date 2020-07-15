package store

import (
	"github.com/pkg/errors"
)

type Selector struct {
	score map[string]int		// 剩余需要进行的测试次数
	arraySet []string			// 调度顺序, 根据原始数据每次随机生成
	i int						// 当前位置
	curItem *Item				// 当前item
	batch *Batch
}

func NewSelector(batch *Batch) *Selector {
	score, arraySet := make(map[string]int), make([]string, 0)
	for name := range batch.store {
		score[name] = 1
		arraySet = append(arraySet, name)
	}
	return &Selector{
		score,
		arraySet,
		0,
		nil,
		batch,
	}
}

// 置为初始状态
func (s *Selector) Init() {
	s.arraySet = []string{}
	for name := range s.batch.store {
		s.score[name] = 1
		s.arraySet = append(s.arraySet, name)
	}
	s.i = 0
}

// 当前完成进度
func (s *Selector) FinishCount() (c int) {
	for _, score := range s.score {
		if score < 1 {
			c++
		}
	}
	return
}

// item分数
func (s *Selector) ItemScore(itemName string) int {
	return s.score[itemName]
}

// 取出当前item字符串
func (s *Selector) CurQus() string {
	return s.arraySet[s.i]
}

// 取出当前item
func (s *Selector) CurItem() *Item {
	if s.curItem == nil {
		panic(errors.New("CurItem is nil"))
	}
	return s.curItem
}

func (s *Selector) IsEmpty() bool {
	return len(s.batch.store) == 0
}

// 指向下一项
func (s *Selector) next() {
	if s.i++; s.i >= len(s.batch.store) {
		s.i = 0
	}
}

// 有下一项
func (s *Selector) HasNext() bool {
	for _, n := range s.score {
		if n > 0 {
			return true
		}
	}
	return false
}

// 取出一个Item, 指针指向当前item
func (s *Selector) PopItem() *Item {
	if !s.HasNext() {
		s.curItem = EofItem
		return EofItem
	}
	if len(s.batch.store) == 0 {
		s.curItem = NilItem
		return s.curItem
	} else {
		s.next()
		// 只弹出剩余测试大于0的
		qus := s.arraySet[s.i]
		for s.score[qus] <= 0 {
			s.next()
			qus = s.arraySet[s.i]
		}
		s.curItem = NewItem(qus, s.batch.store[qus])
		return s.curItem
	}
}

// 获得batch对象
func (s *Selector) Batch() *Batch {
	return s.batch
}

// 减少当前item的重复次数
func (s *Selector) Deduct(v int) {
	s.score[s.curItem.Qus] -= v
	if s.score[s.curItem.Qus] < 0 {
		s.score[s.curItem.Qus] = 0
	}
}

