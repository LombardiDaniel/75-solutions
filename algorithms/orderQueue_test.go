package algorithms_test

import (
	"solutions/algorithms"
	"testing"
)


func TestOrderQueue(t *testing.T) {
	queue := algorithms.NewOrderQueue()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(5)

	queue.Print()
	// queue.PrintBackward()

	v := queue.Dequeue()
	if v != 1 {
		t.Errorf("got %d, expected 1", v)
	}

	queue.Print()

	v = queue.Dequeue()
	if v != 2 {
		t.Errorf("got %d, expected 2", v)
	}

	queue.Print()
	
	queue.Enqueue(7)

	queue.Print()

	v = queue.Dequeue()
	if v != 5 {
		t.Errorf("got %d, expected 5", v)
	}
}