// very simple queue implement for specific scene
package queue

import (
	"fmt"
)

type Element struct {
	prev, next *Element
	queue      *Queue
	Value      interface{}
}

type Queue struct {
	head, tail *Element
	len        int
}

func (q *Queue) Init() *Queue {
	q.len = 0
	return q
}

func New() *Queue {
	return new(Queue).Init()
}

func (q *Queue) Len() int {
	return q.len
}

func (q *Queue) Empty() bool {
	return !(q.len > 0)
}

func (q *Queue) Front() interface{} {
	if q.len > 0 {
		return q.head.Value
	}
	return nil
}

func (q *Queue) Back() interface{} {
	if q.len > 0 {
		return q.tail.Value
	}
	return nil
}
func (q *Queue) Push(v interface{}) {
	e := &Element{Value: v}
	if q.len > 0 {
		e.prev = q.tail
		q.tail.next = e
	}
	q.tail = e
	q.len++
	if q.len == 1 {
		q.head = q.tail
	}
}

func (q *Queue) Pop() (i interface{}) {
	if q.len > 0 {
		i = q.head.Value
		q.head = q.head.next
		q.len--
	}
	return nil
}

func (q *Queue) String() (str string) {
	cur := q.head
	for cur != nil {
		str += fmt.Sprint("->", cur.Value)
		cur = cur.next
	}
	str += fmt.Sprint("\n")
	return str
}
