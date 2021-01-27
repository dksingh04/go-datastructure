package main

import (
	"fmt"
	"go-datastructure/stack"
)

func main() {
	st := stack.Stack{Capacity: 5}
	st.Push(10)
	st.Push(20)
	st.Push(30)
	st.Push(40)
	st.Push(50)
	st.Push(60)
	st.Print()

	val := st.Pop()
	fmt.Printf("Pop Value: %v\n", val)
	st.Pop()
	st.Pop()
	st.Pop()
	st.Pop()
	st.Pop()

	st.Print()
}
