package queue

import "testing"

var tests = []struct {
	capacity int
	input    []int
	want     bool
}{
	{1, []int{10}, true},
	{1, []int{10, 20}, false},
}

func TestQueueFunctionality(t *testing.T) {
	for _, test := range tests {
		queue := Queue{Capacity: test.capacity}
		var got bool
		for _, data := range test.input {
			got = queue.Add(data)
		}

		if got != test.want {
			t.Errorf("Test fail for %v \n", test.input)
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
