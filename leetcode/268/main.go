package main

import "fmt"

func missingNumber(nums []int) int {
	xor := 0
	// 0,0,1,1,2,3 --> 2,3
	for i, row := range nums {
		xor ^= i ^ row
	}

	// 2,3 -> 2
	return xor ^ len(nums)
}

func main() {
	//fmt.Println(missingNumber([]int{3, 0, 1}))
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}
