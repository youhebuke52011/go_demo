package main

import "fmt"

func maxProfit(prices []int) int {
	maxPrice, res := 0, 0
	for i := len(prices)-1; i >= 0; i-- {
		if prices[i] > maxPrice {
			maxPrice = prices[i]
		}
		if maxPrice - prices[i] > res {
			res = maxPrice - prices[i]
		}
	}
	return res
}

func main() {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
	fmt.Println(maxProfit([]int{7,6,4,3,1}))
}
