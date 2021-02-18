package linkedlist

import (
	"fmt"
	"testing"
)

var tests = []struct {
	input []int
	want  bool
	size  int
	rVal  int
	fSize int
}{
	{[]int{10}, true, 1, 10, 0},
	{[]int{10, 4, 6, 8, 10}, true, 5, 10, 4},
	{[]int{6, 8, 6, 3, 9, 20, 54, 40}, true, 8, 40, 7},
}

func TestLinkedListFunctionality(t *testing.T) {
	for _, test := range tests {
		l := NewLinkedList()
		var got bool
		var err error
		// Test Add functionality
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
		fmt.Println("**** Before removing elements *****")
		l.Print()

		// Test Remove functionality
		rVal, err := l.Remove()
		if err != nil {
			t.Errorf("Remove of an element from the LinkedList failed %v", rVal)
		}
		if rVal != test.rVal {
			t.Errorf("Test fail for expected: %v but got: %v\n", test.rVal, rVal)
		}

		if l.Size() != test.fSize {
			t.Errorf("Test fail for Incorrect size expeted: %v but got: %v\n", test.size, l.Size())
		}
		fmt.Println("**** After removing elements *****")
		l.Print()

	}

}
