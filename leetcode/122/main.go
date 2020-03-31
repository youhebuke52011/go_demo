package main

import "fmt"

/**
多次买入卖出，最大收益
每一个上坡加起来
 */

func maxProfit(prices []int) int {
	length := len(prices)
	if length < 2 {
		return 0
	}
	res := 0
	curMin, curMax := prices[0], prices[0]
	for i := 0; i < length-1; i++ {
		if prices[i] < prices[i+1] {
			if curMin > prices[i] {
				curMin = prices[i]
			}
			if curMax < prices[i+1] {
				curMax = prices[i+1]
			}
		} else {
			res += curMax - curMin
			curMin, curMax = prices[i+1], prices[i+1]
		}
	}
	res += curMax - curMin
	return res
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfit([]int{7,6,4,3,1}))
}
