package main

import "fmt"

func maxProfit1(prices []int) int {
    l := 0
	r := 0

	maxProfit := 0

	for i, v := range prices {
		if v <= prices[l] {
			l = i
		}

		if v >= prices[r] {
			r = i
		}

		if l > r {
			r = l
		}

		profit := prices[r] - prices[l]
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	fmt.Println(l, r)
	fmt.Println(prices[l], prices[r])

	return maxProfit
}

// func main() {
	
// 	arr := []int{7,1,5,3,6,4}

// 	r := maxProfit(arr)

// 	fmt.Println(r)
// }