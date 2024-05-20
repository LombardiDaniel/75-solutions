package algorithms

func QuickSort(arr []int) []int {
	return quickSortPivot(arr, 0, len(arr) - 1)
}

func quickSortPivot(arr []int, leftIdx int, rightIdx int) []int {
	if leftIdx < rightIdx {
		return arr[leftIdx:rightIdx]
	}
	return nil
}
