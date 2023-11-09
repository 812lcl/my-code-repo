package structure

import "fmt"

type IStack interface {
	Push(interface{})
	Pop() (interface{}, bool)
	Top() (interface{}, bool)
	IsEmpty() bool
}
type Stack []interface{}

func (s *Stack) Push(val interface{}) {
	*s = append(*s, val)
}

func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	index := len(*s) - 1
	val := (*s)[index]
	*s = (*s)[:index]
	return val, true
}

func (s *Stack) Top() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	index := len(*s) - 1
	return (*s)[index], true
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

type LinkedStack struct {
	top *node
}

type node struct {
	val  interface{}
	next *node
}

func (s *LinkedStack) Push(val interface{}) {
	s.top = &node{val: val, next: s.top}
}

func (s *LinkedStack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	val := s.top.val
	s.top = s.top.next
	return val, true
}

func (s *LinkedStack) Top() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	return s.top.val, true
}

func (s *LinkedStack) IsEmpty() bool {
	return s.top == nil
}

func StackFunc() {
	s := Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Printf("Value of s is %+v\n", s)
	if v, ok := s.Top(); ok {
		fmt.Printf("Value of s top is %+v\n", v)
	}
	fmt.Println(s.Pop()) // 3 true
	fmt.Println(s.Pop()) // 2 true
	fmt.Println(s.Pop()) // 1 true
	fmt.Println(s.Pop()) // nil false

	ls := LinkedStack{}
	ls.Push(1)
	ls.Push(2)
	ls.Push(3)
	fmt.Printf("Value of ls is %+v\n", ls)
	if v, ok := ls.Top(); ok {
		fmt.Printf("Value of ls top is %+v\n", v)
	}
	fmt.Println(ls.Pop()) // 3 true
	fmt.Println(ls.Pop()) // 2 true
	fmt.Println(ls.Pop()) // 1 true
	fmt.Println(ls.Pop()) // nil false
}
