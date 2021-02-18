package linkedlist

import (
	"fmt"

	"github.com/pkg/errors"
)

//TODO other functionality implementation is in progress
type node struct {
	data interface{}
	prev *node
	next *node
}

//LinkedList Data-Structure
type LinkedList struct {
	head *node
	size int
}

//NewLinkedList function to create new LinkedList data-structure
func NewLinkedList() *LinkedList {
	return &LinkedList{head: nil}
}

// Add item to the LinkedList
func (l *LinkedList) Add(data interface{}) (bool, error) {
	cNode := &node{data: data, prev: nil, next: nil}
	if l.head == nil {
		l.head = cNode
		l.size++
		return true, nil
	}
	cNode.next = l.head
	l.head.prev = cNode
	l.head = cNode
	l.size++
	return true, nil
}

// Remove will remove the element from the head
func (l *LinkedList) Remove() (interface{}, error) {
	if l.head == nil {
		return nil, errors.New("linkedlist is empty")
	}
	data := l.head.data
	l.head.data = nil
	l.head = l.head.next
	l.size--
	return data, nil
}

// Print will print formatted data in the LinkedList
func (l *LinkedList) Print() {
	var builder []interface{}
	current := l.head
	i := 0
	for current != nil {
		if i == 0 {
			builder = append(builder, "nil<-[prev |")
			builder = append(builder, current.data)
			builder = append(builder, "| next]")
		} else {
			builder = append(builder, "-> <-[prev |")
			builder = append(builder, current.data)
			builder = append(builder, "| next]")
		}
		i++
		current = current.next
	}
	for _, val := range builder {
		fmt.Printf("%v", val)
	}
	fmt.Println()
}

// Size function return size of the LinkedList
func (l *LinkedList) Size() int {
	return l.size
}
