package main

import (
	"fmt"

	"github.com/laleman11/go-practices-lab/internal/datastructures"
)

func main() {
	stack := datastructures.NewStack[int]()
	stack.Push(12)
	stack.Push(32)

	for i := 0; i < 3; i++ {
		resp, err := stack.Pop()

		if err == nil {
			fmt.Printf("the last element deleted is %d \n", resp)
		} else {
			fmt.Printf("⚠️ Error: %s\n", err)
		}
	}
	resp, err := stack.Peek()
	if err != nil {
		fmt.Printf("⚠️ Error: %s\n", err)
	} else {
		fmt.Printf("🔍 the last element in the list is : %d\n", resp)
	}

	fmt.Printf("the size in the stack is: %d \n", stack.Size())
	fmt.Println("the stack is empty: ", stack.IsEmpty())
}
