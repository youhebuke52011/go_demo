package main

import "fmt"

func singleNumber(nums []int) int {
	ones, twos := 0, 0
	for i := 0; i < len(nums); i++ {
		ones = (ones ^ nums[i]) & ^twos
		twos = (twos ^ nums[i]) & ^ones
	}
	return ones
}

func main() {
	fmt.Println(singleNumber([]int{2,2,3,2}))
}
