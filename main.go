package main

import (
	"fmt"
	"solutions/datastructures"
)

func main() {
	// h := datastructures.MaxHeap{
	// 	Array: []int{},
	// }

	// // build := []int{50, 16, 48, 14, 8, 34, 20, 9, 1, 5, 7, 63}
	// build := []int{10, 20, 30, 10}
	// for _, v := range build {
	// 	h.Insert(v)
	// 	fmt.Println(h)
	// }

	// fmt.Println(h)

	// l := datastructures.LinkedList{
	// 	Head: nil,
	// }

	// build := []int{5, 20, 30, 10}
	// // build := []int{5}
	// for _, v := range build {
	// 	l.Insert(v)
	// }
	// l.Print()

	// fmt.Println("Done")

	// remove := []int{5, 10, 20, 80}
	// for _, v := range remove {
	// 	l.Remove(v)
	// }
	// l.Print()

	// s := datastructures.NewStack()
	// build := []int{5, 20, 30, 10}
	// for _, v := range build {
	// 	s.Push(v)
	// }
	// fmt.Println(s)
	
	// for range 2 {
	// 	s.Pop()
	// }

	// fmt.Println(s)

	// q := datastructures.NewQueue()
	// build := []int{5, 20, 30, 10}
	// for _, v := range build {
	// 	q.Enqueue(v)
	// }
	// fmt.Println(q)

	// for range 2 {
	// 	fmt.Println(q.Dequeue())
	// }
	
	t := datastructures.NewBinaryTree()
	build := []int{5, 10, 9}
	for _, v := range build {
		t.Insert(v)
	}

	fmt.Println("\ntree:")
	fmt.Println(t.Root)
	fmt.Println(t.Root.L, t.Root.R)

}