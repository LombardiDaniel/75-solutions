package main

func lengthOfContinuousLIS(nums []int) int {
    currLen := 1
    largest := 1
    last := nums[0]

    for _, v := range nums {
        if v > last {
            currLen++
            largest = currLen
        } else {
            currLen = 1
        }
        last = v
    }

    return largest
}