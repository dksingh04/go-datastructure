package stack

import (
	"fmt"
)

//Stack defenition
type Stack struct {
	items    []interface{}
	pos      int
	Capacity int
}

//Push element
func (s *Stack) Push(data interface{}) bool {
	if !s.isStackFull() {
		s.items = append(s.items, data)
		s.pos++
		return true
	}
	fmt.Printf("Stack is full now, Capacity of the Stack is: %v and size: %v of the Satck is Same\n", s.Capacity, s.pos)
	return false

}

//Pop element
func (s *Stack) Pop() interface{} {
	if !s.isStackEmpty() {
		item := s.items[s.pos-1]
		s.items = s.items[:s.pos-1]
		s.pos--
		return item
	}
	fmt.Printf("Stack is Empty now, Capacity of the Stack is: %v and size: %v of the Satck is Zero\n", s.Capacity, s.pos)
	return -1

}

// Size of the Stack
func (s *Stack) Size() int {
	return s.pos
}

// IsStackEmpty Stack
func (s *Stack) isStackEmpty() bool {
	return s.pos == 0
}

//IsStackFull - check to find if Stack is full
func (s *Stack) isStackFull() bool {
	return s.pos == s.Capacity
}

//Print stack data
func (s *Stack) Print() {
	fmt.Println(s.items)
}
