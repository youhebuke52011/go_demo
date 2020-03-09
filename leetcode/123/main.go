package main

import "fmt"

func maxProfit(prices []int) int {
	size := len(prices)
	if size <= 1 {
		return 0
	}

	profits := []int{}
	temp := 0
	for i := 1; i < size; i++ {
		if prices[i]-prices[i-1] != 0 {
			profits = append(profits, prices[i]-prices[i-1])
		}
	}

	//for i := 1; i < size; i++ {
	//	diff := prices[i] - prices[i-1]
	//
	//	if temp*diff >= 0 {
	//		temp += diff
	//		continue
	//	}
	//
	//	profits = append(profits, temp)
	//	temp = diff
	//}
	//profits = append(profits, temp)

	res := 0
	for i := 0; i < len(profits); i++ {
		temp = max(profits[:i]) + max(profits[i:])
		if res < temp {
			res = temp
		}
	}

	return res
}

func max(ps []int) int {
	max, tmp := 0, 0

	for _, p := range ps {
		if tmp < 0 {
			tmp = 0
		}

		tmp += p

		if max < tmp {
			max = tmp
		}
	}

	return max
}

func main() {
	fmt.Println(maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
