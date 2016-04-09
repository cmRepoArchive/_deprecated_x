// very simple stack implement for a clear view
package stack

import (
	"fmt"
)

type Element struct {
	prev  *Element
	stack *Stack
	Value interface{}
}

type Stack struct {
	top *Element
	len int
}

func (s *Stack) Init() *Stack {
	s.len = 0
	return s
}

func New() *Stack {
	return new(Stack).Init()
}

func (s *Stack) Len() int {
	return s.len
}

func (s *Stack) Empty() bool {
	return !(s.len > 0)
}

// func (s *Stack) TopElem() *Element {
// 	if s.len > 0 {
// 		return s.top
// 	}
// 	return nil
// }

func (s *Stack) Top() interface{} {
	if s.len > 0 {
		return s.top.Value
	}
	return nil
}

func (s *Stack) Pop() (i interface{}) {
	if s.len > 0 {
		i = s.top.Value
		s.top = s.top.prev
		s.len--
		return i
	}
	return nil
}

// func (s *Stack) PushElm(e *Element) {
// 	e.prev = s.top
// 	e.stack = s
// 	s.top = e
// 	s.len++
// }

func (s *Stack) Push(v interface{}) {
	e := &Element{Value: v}
	e.prev = s.top
	e.stack = s
	s.top = e
	s.len++
}

func (s *Stack) String() (str string) {
	cur := s.top
	for cur != nil {
		str += fmt.Sprint("->", cur.Value)
		cur = cur.prev
	}
	str += fmt.Sprint("\n")
	return str
}
