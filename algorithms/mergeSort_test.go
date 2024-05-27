package algorithms_test

import (
	"fmt"
	"solutions/algorithms"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{7,8,7,5,4,10,3}
	fmt.Println(algorithms.MergeSort(arr))
}