package datastructures

// FIFO
type Queue[T any] struct {
	arr []T
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{
		arr: []T{},
	}
}

func (q *Queue[T]) Enqueue(v T) {
	q.arr = append(q.arr, v)
}

func (q *Queue[T]) Dequeue() *T {
	if len(q.arr) == 0 {
		return nil
	}

	ret := q.arr[0]

	q.arr = q.arr[1:]
	return &ret
}