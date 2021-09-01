package DataStructures

import (
	"fmt"
	"github.com/datastructures/node"
	"reflect"
	"sync"
)

var (
	head *node.Node
	tail *node.Node
	mux  *sync.RWMutex
)

type CustomList struct {
	len    int
	Hetero bool
	//to keep indexes it will cost us log(n) time to get data
	hash map[int]*node.Node
	typ  reflect.Type
}

func add(a, b int) (res int) {
	res = a + b
	return
}

type List interface {
	PushBack(data interface{})           //	O(logn)
	PopBack() interface{}                //	O(logn)
	PopFront() interface{}               //	O(logn)
	PushFront() interface{}              //	O(logn)
	Remove(index int) interface{}        //	O(logn)
	Get(index int) interface{}           //	O(logn)
	size() int                           //	O(1)
	Add(index int, data interface{})     //	O(n*logn)
	Replace(index int, data interface{}) //	O(logn)
	Make(data interface{}, Hetero bool) *CustomList
	GetAll() []interface{}                       //	O(n)
	GetInRange(start int, end int) []interface{} //	O(n)
}

func (c *CustomList) PushFront(data interface{}) {
	typeCheck(data, c.typ, c.Hetero)
	mux.Lock()
	temp := &node.Node{
		Data: data,
		Next: head,
		Prev: nil,
	}
	head.Prev = temp
	head = temp
	ind := 0
	for temp != nil {
		c.hash[ind] = temp
		ind++
		temp = temp.Next
	}
	c.len++
	mux.Unlock()
}

func (c *CustomList) PopBack() interface{} {
	indexCheck(c.len-1, c.len)
	var res interface{}
	mux.Lock()
	temp := tail
	res = temp.Data
	delete(c.hash, c.len-1)
	tail = temp.Prev
	temp.Prev = nil
	tail.Next = nil
	c.len--
	mux.Unlock()
	return res
}

func (c *CustomList) PopFront() interface{} {
	indexCheck(0, c.len)
	var res interface{}
	mux.Lock()
	temp := head
	res = temp.Data
	delete(c.hash, 0)
	head = temp.Next
	temp.Next = nil
	head.Prev = nil
	temp = head
	index := 0
	for temp != nil {
		c.hash[index] = temp
		temp = temp.Next
		index++
	}
	c.len--
	mux.Unlock()
	return res
}

func (c *CustomList) Remove(index int) interface{} {
	indexCheck(index, c.len)
	var res interface{}
	mux.Lock()
	if index == 0 {
		c.PopFront()
	} else if index == c.len-1 {
		c.PopBack()
	} else {
		temp := c.hash[index]
		prev := temp.Prev
		next := temp.Next
		prev.Next = next
		next.Prev = prev
		for next != nil {
			c.hash[index] = next
			next = next.Next
			index++
		}
		c.len--
	}
	mux.Unlock()
	return res
}

func indexCheck(index int, len int) {
	if index >= len && index < 0 {
		panic("wrong index")
	}
}

func typeCheck(data interface{}, typ reflect.Type, hetero bool) {
	if reflect.TypeOf(data) != typ && hetero == false {
		panic("Type Mismatched : " +
			"\n" +
			"Actual Type: " + fmt.Sprint(typ) + "\n" +
			"Given Type: " + fmt.Sprint(reflect.TypeOf(data)))
	}
}

func (c *CustomList) GetAll() []interface{} {
	temp := head
	index := 0
	res := make([]interface{}, c.len+1)
	for temp != nil {
		mux.RLock()
		res[index] = temp.Data
		index++
		temp = temp.Next
		mux.RUnlock()
	}
	return res
}

func (c *CustomList) GetInRange(start int, end int) []interface{} {
	indexCheck(start, c.len)
	indexCheck(end, c.len)
	siz := end - start + 1
	count := 0
	res := make([]interface{}, siz)
	temp := c.hash[start]
	for count < siz {
		mux.RLock()
		res[count] = temp.Data
		temp = temp.Next
		mux.RUnlock()
		count++
	}
	return res
}

func NewList(data interface{}, Hetero bool) *CustomList {
	head = nil
	tail = nil
	mux = &sync.RWMutex{}
	return &CustomList{
		len:    0,
		Hetero: Hetero,
		hash:   make(map[int]*node.Node),
		typ:    reflect.TypeOf(data),
	}
}

func (c *CustomList) PushBack(data interface{}) {
	typeCheck(data, c.typ, c.Hetero)
	temp := &node.Node{
		Data: data,
		Next: nil,
		Prev: nil,
	}
	mux.Lock()
	if head == nil && tail == nil {
		head = temp
		tail = temp
		c.hash[c.len] = temp
	} else {
		tail.Next = temp
		temp.Prev = tail
		tail = temp
		c.hash[c.len] = temp
	}
	mux.Unlock()
	c.len++
}

func (c *CustomList) Get(index int) interface{} {
	indexCheck(index, c.len)
	mux.RLock()
	temp := c.hash[index].Data
	mux.RUnlock()
	return temp
}

func (c *CustomList) size() int {
	return c.len
}

func (c *CustomList) Add(index int, data interface{}) {
	indexCheck(index, c.len+1)
	typeCheck(data, c.typ, c.Hetero)
	temp := &node.Node{
		Data: data,
		Next: head,
		Prev: nil,
	}
	mux.Lock()
	if index == 0 {
		temp.Next = head
		head.Prev = temp
		head = temp
	} else if index == c.len-1 {
		c.PushBack(data)
	} else {
		prev := c.hash[index-1]
		next := c.hash[index]
		prev.Next = temp
		temp.Next = next
		temp.Prev = prev
		next = temp.Next
		for temp != nil {
			c.hash[index] = temp
			temp = temp.Next
			index++
		}
	}
	c.len++
	mux.Unlock()
}

func (c *CustomList) Replace(index int, data interface{}) {
	typeCheck(data, c.typ, c.Hetero)
	indexCheck(index, c.len)
	mux.Lock()
	c.hash[index].Data = data
	mux.Unlock()
}
