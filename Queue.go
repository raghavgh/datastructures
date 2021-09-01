package DataStructures

import (
	"github.com/datastructures/node"
	"reflect"
)

type Queue struct {
	Hetero     bool
	tail, head *node.Node
	len        int
	typ        reflect.Type
}

func NewQueue(data interface{}, hetero bool) *Queue {
	return &Queue{
		Hetero: hetero,
		tail:   nil,
		head:   nil,
		len:    0,
		typ:    reflect.TypeOf(data),
	}
}

func (q *Queue) Size() int {
	return q.len
}
func (q *Queue) Enqueue(data interface{}) {
	typeCheck(data, q.typ, q.Hetero)
	temp := &node.Node{
		Data: data,
		Next: nil,
		Prev: nil,
	}
	if q.head == nil {
		q.head = temp
		q.tail = temp
	} else {
		temp.Prev = q.tail
		q.tail.Next = temp
		q.tail = temp
	}
	q.len++
}

func (q *Queue) Dequeue() interface{} {
	var res interface{}
	if q.head == nil {
		panic("Queue is Empty")
	}
	temp := q.head.Next
	res = q.head.Data
	q.head = temp
	q.len--
	return res
}
func (q *Queue) Front() interface{} {
	if q.head == nil {
		panic("Queue is Empty")
	}
	return q.head.Data
}

func (q *Queue) IsEmpty() bool {
	return q.head == nil
}
