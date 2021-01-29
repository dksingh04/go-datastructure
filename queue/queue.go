package queue

import "fmt"

// Queue Data-Structure
type Queue struct {
	items    []interface{}
	pos      int
	Capacity int
}

// Add item in the Queue
func (q *Queue) Add(data interface{}) bool {
	if !q.isQueueFull() {
		q.items = append(q.items, data)
		q.pos++
		return true
	}
	fmt.Printf("Queue is full now, Capacity of the Queue : %v and size : %v is Same\n", q.Capacity, q.pos)
	return false

}

// Remove data from the Queue
func (q *Queue) Remove() interface{} {
	if !q.isQueueEmpty() {
		data := q.items[0]
		q.items = q.items[:-q.pos-1]
		q.pos--
		return data
	}
	fmt.Printf("Queue is Empty now, Capacity of the Queue is: %v and size of Queue is: %v \n", q.Capacity, q.pos)
	return -1
}

// Peek to look for the value, it does not remove the element from the Queue
func (q *Queue) Peek() interface{} {
	if !q.isQueueEmpty() {
		data := q.items[0]
		return data
	}
	fmt.Printf("Queue is Empty now, Capacity of the Queue is: %v and size of Queue is: %v \n", q.Capacity, q.pos)
	return -1
}

//Size of the Queue
func (q *Queue) Size() int {
	return q.pos
}

//IsQueueFull function to check, If Queue is full
func (q *Queue) isQueueFull() bool {
	return q.pos == q.Capacity
}

//IsQueueEmpty function to check if Queue is Empty
func (q *Queue) isQueueEmpty() bool {
	return q.pos == 0
}
