package queue

import (
	"testing"
)

var tests = []struct {
	capacity         int
	input            []int
	want             bool
	befRemovalSize   int
	afterRemovalSize int
}{
	{1, []int{10}, true, 1, 0},
	{2, []int{10, 20}, true, 2, 1},
	{3, []int{10, 20, 30}, true, 3, 2},
	{3, []int{10, 20, 30, 40}, false, 3, 2},
}

func TestQueueFunctionality(t *testing.T) {
	for _, test := range tests {
		queue := Queue{Capacity: test.capacity}
		var got bool
		// Test Add functionality
		for _, data := range test.input {
			got = queue.Add(data)
		}
		if queue.Size() != test.befRemovalSize {
			t.Errorf("Test fail for, Elements didin't got added successfully %v\n", test.input)
		}
		if got != test.want {
			t.Errorf("Test fail for %v \n", test.input)
		}

		//Test Peek Functionality
		pVal := queue.Peek()
		if pVal != test.input[0] {
			t.Errorf("Test fail while Peek element function not worked %v \n", test.input)
		}

		//Also Peek should not remove element
		if queue.Size() != test.befRemovalSize {
			t.Errorf("Test fail After Peek, size of the Queue is not same as earlier Queue Size after Peek: %v, Queue Size Before Peek: %v \n", queue.Size(), test.befRemovalSize)
		}

		// Test Removal Functionality
		val := queue.Remove()
		if val != test.input[0] {
			t.Errorf("Test fail while removing element %v \n", test.input)
		}

		if test.befRemovalSize != test.afterRemovalSize+1 {
			t.Errorf("Elment didn't got removed properly Before Queue Size: %v and After Removal Queue Size: %v", test.befRemovalSize, test.afterRemovalSize)
		}

	}
}

func TestSuccessfulAddOperation(t *testing.T) {
	queue := Queue{Capacity: 1}
	if !queue.Add(10) {
		t.Errorf("Error in Adding item to the Queue!")
	}
}

func TestQueuefulAddOperation(t *testing.T) {
	queue := Queue{Capacity: 1}
	queue.Add(10)
	if queue.Add(20) {
		t.Errorf("Queue is allowing to add element even Queue is Full!")
	}

}
