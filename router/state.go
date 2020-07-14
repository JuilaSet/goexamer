package router

type TodoFunc func()
type ChangeStateFunc func(input interface{})

type State struct {
	Todo func()	// 规定做什么
	ChangeState func(input interface{}) // 规定如何切换状态
}

func NewState(Todo func(), ChangeState func(input interface{})) *State {
	return &State{
		Todo, ChangeState,
	}
}