package main

import "fmt"

func singleNumber(nums []int) int {
	res := 0
	for _, row := range nums {
		res = res ^ row
	}
	return res
}

func main() {
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println(singleNumber([]int{2,2,1}))
}
