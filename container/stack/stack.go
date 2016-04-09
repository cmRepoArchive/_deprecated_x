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

func (s *Stack) Top() *Element {
	return s.top
}

func (s *Stack) TopValue() interface{} {
	return s.top.Value
}

func (s *Stack) Pop() {
	s.top = s.top.prev
	s.len--
}

func (s *Stack) Push(e *Element) {
	e.prev = s.top
	s.top = e
	s.len++
}

func (s *Stack) PushValue(v interface{}) {
	e := &Element{Value: v}
	e.prev = s.top
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
