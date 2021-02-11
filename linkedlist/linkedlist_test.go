package linkedlist

import (
	"testing"
)

var tests = []struct {
	input []int
	want  bool
	size  int
}{
	{[]int{10}, true, 1},
	{[]int{10, 4, 6, 8, 10}, true, 5},
	{[]int{6, 8, 6, 3, 9, 20, 54, 40}, true, 8},
}

func TestLinkedListFunctionality(t *testing.T) {
	for _, test := range tests {
		l := NewLinkedList()
		var got bool
		var err error
		for _, data := range test.input {
			got, err = l.Add(data)
			if err != nil {
				t.Errorf("Adding element to LinkedList failed %v", data)
			}
			if got != test.want {
				t.Errorf("Test fail for %v \n", test.input)
			}
		}

		if l.Size() != test.size {
			t.Errorf("Test fail for Incorrect size expeted: %v but got: %v\n", test.size, l.Size())
		}

	}
}
