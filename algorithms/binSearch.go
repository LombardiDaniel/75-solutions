package algorithms

import "fmt"

func BinarySearch(sortedArr []int, val int) int {
	fmt.Println(sortedArr)
	// return binarySearchInternalRecursive(sortedArr, val, 0, len(sortedArr) - 1)
	return binarySearchInternalIterative(sortedArr, val)
}

func binarySearchInternalRecursive(sortedArr []int, val int, startIdx int, finishIdx int) int {
	searchIdx := startIdx + (finishIdx - startIdx) / 2
	searchedVal := sortedArr[searchIdx]
	deltaIdx := finishIdx - startIdx

	fmt.Println(searchIdx, searchedVal)

	if val == searchedVal {
		return searchIdx
	}

	if deltaIdx == 1 {
		return -1
	}

	if val > searchedVal {
		return binarySearchInternalRecursive(sortedArr, val, searchIdx + 1, finishIdx)
	}

	if val < searchedVal {
		return binarySearchInternalRecursive(sortedArr, val, startIdx, searchIdx - 1)
	}

	return -1
}

func binarySearchInternalIterative(sortedArr []int, val int) int {
	low := 0
	high := len(sortedArr) - 1

	for low <= high {
		mid := (high + low) / 2
		guess := sortedArr[mid]

		if guess == val {
			return mid
		}

		if val > guess {
			low = mid + 1
		} else if val < guess {
			high = mid - 1
		}
	}

	return -1
}