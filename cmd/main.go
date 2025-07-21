package main

import (
	"github.com/laleman11/go-practices-lab/internal/datastructures"
)

func main() {
	queue := datastructures.NewLinkedList[int]()
	queue.Add(10)
	queue.Add(20)
	queue.Add(50)
	queue.Prepend(1)
	queue.Prepend(-1)
	queue.Delete(10)
	_, exist := queue.Find(10)
	_, exist2 := queue.Find(1)
	size := queue.Size()
	s := queue.ToSlice()

	println(queue)
	println(size)
	println(exist)
	println(exist2)
	println(s)
	queue.Print()

}
