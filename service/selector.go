package service

import (
	"goexamer/store"
	"math"
	"strconv"
)

const (
	C = 0.3 // 冷却系数
	C0 = 0.5
)

type Selector struct {
	score      map[string]int     // 剩余需要进行的测试次数
	arraySet   []string           // 调度顺序, 根据原始数据每次随机生成
	i          int                // 当前位置只负责调度
	curItem    *Item              // 当前item, 可能会出现当前的item与s.i指向不一样的情况
	lastWrongN map[string]int     // item的出现次数
	lastN      map[string]int     // item的上次出现次数
	hotFactor  map[string]float64 // 熟悉因子
	batch      *store.Batch
	n          int						// 弹出item次数
}

func NewSelector(batch *store.Batch) *Selector {
	s := &Selector{
		make(map[string]int),
		make([]string, 0),
		0,
		nil,
		make(map[string]int),
		make(map[string]int),
		make(map[string]float64),
		batch,
		0,
	}
	s.Init()
	return s
}

// 置为初始状态
func (s *Selector) Init() {
	s.arraySet = []string{}
	for name := range s.batch.GetAllQus() {
		s.score[name] = 1
		s.lastWrongN[name] = -1
		s.hotFactor[name] = 1.5
		s.lastN[name] = 0
		s.arraySet = append(s.arraySet, name)
	}
	s.i = 0
	s.n = 0
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
	return s.curItem.Qus
}

// 取出当前item
func (s *Selector) CurItem() *Item {
	if s.curItem == nil {
		return NilItem
	}
	return s.curItem
}

// 直接设置当且执行对象的指针
func (s *Selector) SetCurItemDangerous(item *Item) {
	s.curItem = item
}

func (s *Selector) IsEmpty() bool {
	return len(s.batch.GetAllQus()) == 0
}

// 指向下一项
func (s *Selector) next(i int) int {
	if i++; i >= len(s.batch.GetAllQus()) {
		i = 0
	}
	return i
}

// 设置下一项
func (s *Selector) SetNext(name string) {
	if _, ok := s.score[name]; ok {
		ai := 0
		for k, qus := range s.arraySet {
			if qus == name {
				ai = k
				break
			}
		}
		bi := s.i + 1
		if bi >= len(s.arraySet) {
			bi = 0
		}
		s.arraySet[bi], s.arraySet[ai] = s.arraySet[ai], s.arraySet[bi]
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

// 跳转到某项
func (s *Selector) SetJmp(itemName string) {
	if n, ok := s.score[itemName]; ok && n <= 1 {
		s.score[itemName] = 2
	} else if !ok {
		return
	}
	s.SetNext(itemName)
}

// 取出下一项Item
func (s *Selector) NextItem() (*Item, int) {
	if !s.HasNext() {
		return EofItem, -1
	}
	if len(s.batch.GetAllQus()) == 0 {
		return NilItem, -2
	} else {
		i := s.next(s.i)
		var qus string
		for {
		qus = s.arraySet[i]
		if s.score[qus] > 0 {
			break
		}
		i = s.next(i)
		}
		return NewItem(qus, s.batch.GetAllQus()[qus]), i
	}
}

// 取出一个Item, 指针指向当前item, 只弹出剩余测试大于0的
func (s *Selector) PopItem() *Item {
	s.curItem, s.i = s.NextItem()
	return s.curItem
}

// 获取调度系数
func (s *Selector) DispatchCoefficient(name string) float64 {
	return s.hotFactor[name]
}

// 获取调度系数字符串
func (s *Selector) DispatchCoefficientString(name string) string {
	return strconv.FormatFloat(selector.DispatchCoefficient(name),'f',2,64)
}

// 找到调度系数最小的项名称, 每次选择hotFactor最小的弹出来
func (s *Selector) MinHotFactorItemName() string {
	focus, key := 100.0, ""
	for k, v := range s.hotFactor {
		if focus > v {
			focus, key = v, k
		}
	}
	return key
}

// 计算调度系数(q ∈(0, 4), 当q=4减去的是0)
func (s *Selector) CalcDispatchCoefficient(q float64, correct bool) float64 {
	name := s.curItem.Qus
	s.n++
	if correct {
		if s.lastWrongN[name] == -1 {
			s.hotFactor[name] += q *  1.0
		} else {
			s.hotFactor[name] += q *  float64(s.n - s.lastN[name])
		}
	} else {
		s.lastWrongN[name] = s.n
		s.hotFactor[name] *= math.Exp(-C0 * float64(s.n - s.lastN[name]))
	}
	s.lastN[s.curItem.Qus] = s.n
	// 其他项目, 没有出现错误的是-1, 不需要进行计算
	for qus := range s.hotFactor {
		if qus != name && s.lastWrongN[qus] != -1 && !s.IsFinish(s.hotFactor[qus]){
			s.hotFactor[qus] *= math.Exp(-C * float64(s.n-s.lastN[qus]))
		}
	}
	return s.hotFactor[name]
}

// 是否完成
func (s *Selector) IsFinish(ef float64) bool {
	if ef >= 2.5 {
		return true
	}
	return false
}

// 获得batch对象
func (s *Selector) Batch() *store.Batch {
	return s.batch
}

// 减少item的出现次数
func (s *Selector) DeductItem(itemName string, v int) {
	if _, ok := s.score[itemName]; ok {
		s.score[itemName] -= v
		if s.score[itemName] < 0 {
			s.score[itemName] = 0
		}
	}
}

// 减少当前item的重复次数
func (s *Selector) Deduct(v int) {
	s.DeductItem(s.curItem.Qus, v)
}

// 设置当前item的执行次数
func (s *Selector) SetScore(itemName string, v int) {
	if _, ok := s.score[itemName]; ok {
		s.score[itemName] = v
	}
}

// 执行命令
func (s *Selector) ExecuteBeforeFunc() {
	item, _ := s.NextItem()
	for _, action := range item.ActionBefore {
		action.Func(s, action.Param)
	}
}
func (s *Selector) ExecuteMidFunc() {
	for _, action := range s.curItem.ActionMid{
		action.Func(s, action.Param)
	}
}
func (s *Selector) ExecuteAfterFunc() {
	for _, action := range s.curItem.ActionAfter{
		action.Func(s, action.Param)
	}
}
