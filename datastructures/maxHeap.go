package datastructures

import "fmt"

// type maxHeap struct {
// 	root 	*node
// }

// type node struct {
// 	val 	int
// 	l		*node
// 	r 		*node
// }

// complete tree

// maxHeap -> root = maxVal
// minHeap -> root = minVal

// can be represented as arr very easily -> sortedArr
// Insert and extract is O(h) = O(log n)

type MaxHeap struct {
	Array []int
}

func leftChildIdx(parentIdx int) int {
	return 2 * parentIdx + 1
}

func rightChildIdx(parentIdx int) int {
	return 2 * parentIdx + 2
}

func parentIdx(idx int) int {
	return (idx - 1) / 2
}

func swap[T any](a *T, b *T) {
	tmp := *a

	*a = *b
	*b = tmp
}

func (h *MaxHeap) greaterChildIdx(parentIdx int) int {
	lIdx := leftChildIdx(parentIdx)
	rIdx := rightChildIdx(parentIdx)

	l := h.Array[lIdx]
	r := h.Array[rIdx]

	if l > r {
		return lIdx
	}

	return rIdx
}

// always added to the bottom right, then we check:
// if its greater than the parent, we swap
// and keep doing that, caled heapify
func (h *MaxHeap) Insert(val int) {
	h.Array = append(h.Array, val)
	h.heapifyUp(len(h.Array) - 1)
}

// remove root (greatest) el. and then we heapify down
func (h *MaxHeap) Remove() int {
	if len(h.Array) == 0 {
		fmt.Println("cannot extract")
		return -1
	}

	ret := h.Array[len(h.Array) - 1]

	h.Array[0] = h.Array[len(h.Array) - 1]
	h.Array = h.Array[:len(h.Array) - 1]

	h.heapifyDown(0)

	return ret
}

func (h *MaxHeap) heapifyDown(idx int) {

	lastIdx := len(h.Array) - 1
	lIdx := leftChildIdx(idx)
	rIdx := rightChildIdx(idx)

	greaterChildIdx := 0

	// loop while index has at least one child
	for lIdx <= lastIdx {
		if lIdx == lastIdx { // when left child in the only child
			greaterChildIdx = lIdx
		} else if h.Array[lIdx] > h.Array[rIdx] { // when left child is larger
			greaterChildIdx = lIdx
		} else { // when right child is larger
			greaterChildIdx = rIdx
		}

		if h.Array[idx] < h.Array[greaterChildIdx] {
			swap(&h.Array[lIdx], &h.Array[idx])
			lIdx = idx
		} else {
			return
		}
	}
}

func (h *MaxHeap) heapifyUp(idx int) {
	for h.Array[parentIdx(idx)] < h.Array[idx] {
		swap(&h.Array[parentIdx(idx)], &h.Array[idx])
		idx = parentIdx(idx)
	}
}