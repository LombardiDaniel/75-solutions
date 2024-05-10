package datastructures

// FIFO
type Queue struct {
	arr []int
}

func NewQueue() Queue {
	return Queue{
		arr: []int{},
	}
}

func (q *Queue) Enqueue(v int) {
	q.arr = append(q.arr, v)
}

func (q *Queue) Dequeue() int {
	ret := q.arr[0]

	q.arr = q.arr[1:]
	return ret
}