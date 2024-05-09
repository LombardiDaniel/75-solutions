package main

func productExceptSelf(nums []int) []int {
	n := len(nums)
    leftProds := make([]int, n)
	rightProds := make([]int, n)

	leftProds[0] = 1
	rightProds[n-1] = 1

	for i := 1; i < n; i++ {
		leftProds[i] = nums[i-1] * leftProds[i-1]
	}

	for i := n-2; i >= 0; i-- {
		rightProds[i] = nums[i+1] * rightProds[i+1]
	}

	ret := make([]int, n)
	for i := range n {
		ret[i] = leftProds[i] * rightProds[i]
	}

	return ret
}