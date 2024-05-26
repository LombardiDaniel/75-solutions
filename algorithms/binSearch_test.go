package algorithms_test

import (
	"solutions/algorithms"
	"testing"
)

func TestBinSearch(t *testing.T) {
	sortedArr := []int{1,2,3,4,5,6,7,8,9,99,1200}

	idx := algorithms.BinarySearch(sortedArr, 98)
	if idx != -1 {
		t.Errorf("found 98 @ %d", idx)
	}

	idx = algorithms.BinarySearch(sortedArr, 99)
	if idx != 9 {
		t.Errorf("found 99 @ idx=%d", idx)
	}

	idx = algorithms.BinarySearch(sortedArr, 1200)
	if idx != 10 {
		t.Errorf("found 99 @ idx=%d", idx)
	}
}