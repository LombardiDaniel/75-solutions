package algorithms

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	left := []int{}
	right := []int{}

	for _, v := range arr[1:] {
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	ret := []int{}
	ret = append(ret, QuickSort(left)...)
	ret = append(ret, pivot)
	ret = append(ret, QuickSort(right)...)
	return ret
}
