package main

import (
	"github.com/laleman11/go-practices-lab/internal/datastructures"
)

func main() {
	queue := datastructures.NewLinkedList[int]()
	queue.Add(10)
	queue.Add(30)
	queue.Add(60)
	_, exist := queue.Find(10)
	_, exist2 := queue.Find(1)
	size := queue.Size()
	s := queue.ToSlice()
	queue.InsertAfter(1, 23)
	queue.Print()
	queue.Reverse()
	queue.Print()
	println(size)
	println(exist)
	println(exist2)
	println(s)

}
