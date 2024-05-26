package algorithms

import "fmt"

type node struct {
	val			int
	next		*node
	prev		*node
}

type OrderQueue struct {
	Size		int
	firstNode  	*node
	lastNode 	*node
}

func NewOrderQueue() OrderQueue {
	return OrderQueue{
		firstNode: nil,
		lastNode: nil,
		Size: 0,
	}
}

func (q *OrderQueue) Print() {

	currNode := q.firstNode
	// fmt.Println(currNode)
	for currNode != nil {
		fmt.Printf("%d -> ", currNode.val)
		currNode = currNode.next
	}
	fmt.Println("")
}

func (q *OrderQueue) PrintBackward() {

	currNode := q.lastNode
	// fmt.Println(currNode)
	for currNode != nil {
		fmt.Printf("%d <- ", currNode.val)
		currNode = currNode.prev
	}
	fmt.Println("")
}

func (q *OrderQueue) Enqueue(v int) {
	newNode := &node{
		val: v,
		prev: nil,
		next: nil,
	}

	if q.firstNode == nil {
		q.firstNode = newNode
		q.lastNode = newNode
	} else {
		newNode.next = q.firstNode
		q.firstNode.prev = newNode
		q.firstNode = newNode
	}

	q.Size++
}

func (q *OrderQueue) Dequeue() int {

	if q.lastNode == nil {
		return -1
	}

	ret := q.lastNode.val

	if q.lastNode.prev == nil {
		q.firstNode = nil
		q.lastNode = nil
	} else {
		q.lastNode = q.lastNode.prev
		q.lastNode.next = nil
	}

	q.Size--
	return ret
}