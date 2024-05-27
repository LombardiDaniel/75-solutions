package algorithms

import "fmt"

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	midIdx := len(arr) / 2
	left := make([]int, midIdx)
	right := make([]int, len(arr) - midIdx)

	copy(left, arr[:midIdx])
	copy(right, arr[midIdx:])
	fmt.Println(left, right)

	l := MergeSort(left)
	r := MergeSort(right)

	return merge(l, r)
}

func merge(left []int, right []int) []int {
	arr := []int{}
	lIdx := 0
	rIdx := 0

	for lIdx < len(left) && rIdx < len(right) {
		if left[lIdx] <= right[rIdx] {
			arr = append(arr, left[lIdx])
			lIdx++
		} else {
			arr = append(arr, right[rIdx])
			rIdx++
		}
	}

	for lIdx < len(left) {
		arr = append(arr, left[lIdx])
		lIdx++
	}

	for rIdx < len(right) {
		arr = append(arr, right[rIdx])
		rIdx++
	}

	return arr
}
