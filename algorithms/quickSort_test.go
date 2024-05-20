package algorithms_test

import (
	"solutions/algorithms"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{7,8,7,5,4,10,3}
	println(algorithms.QuickSort(arr))
}