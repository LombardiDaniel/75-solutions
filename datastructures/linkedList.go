package datastructures

import (
	"errors"
	"fmt"
)

type listNode struct {
	Val 		int
	Next		*listNode
}

type LinkedList struct {
	Head		*listNode
}

func (l *LinkedList) Insert(v int) {
	newHead := listNode{
		Val: v,
		Next: l.Head,
	}
	l.Head = &newHead
}

func (l *LinkedList) Remove(v int) error {

	if l.Head == nil {
		return errors.New("not present")
	}

	if l.Head.Val == v {
		l.Head = l.Head.Next
		return nil
	}

	// seek val, or end
	prevToDel := l.Head
	if prevToDel.Next == nil {
		return errors.New("not present")
	}
	for prevToDel.Next.Val != v { // aqui pode dar SEGFAULT se nao tiver next
		prevToDel = prevToDel.Next
		if prevToDel.Next == nil {
			return errors.New("not present")
		}
	}

	// deletion
	prevToDel.Next = prevToDel.Next.Next

	return errors.New("not present")
}

func (l *LinkedList) Print() {
	currNode := l.Head
	for currNode != nil {
		fmt.Println(currNode)
		currNode = currNode.Next
	}
}