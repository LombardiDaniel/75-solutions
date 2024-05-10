package main

import "fmt"

func maxProfit(prices []int) int {
    l := 0
	r := 0
	maxProfit := 0

	for r < len(prices) {
		if prices[l]< prices[r] {
			profit := prices[r] - prices[l]
			if profit > maxProfit {
				maxProfit = profit
			}
		} else {
			l = r
		}

		r += 1
	}

	return maxProfit
}

func main() {
	
	arr := []int{7,1,5,3,6,4}

	r := maxProfit(arr)

	fmt.Println(r)
}