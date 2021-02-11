package linkedlist

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

// Size function return size of the LinkedList
func (l *LinkedList) Size() int {
	return l.size
}
