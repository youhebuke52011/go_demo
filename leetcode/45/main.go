package main

import "fmt"

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func jump(nums []int) int {
	length := len(nums)
	res := 0
	cur, per, i := 0, 0, 0
	for cur < length-1 {
		per = cur
		for ; i <= per; i++ {
			cur = Max(cur, i+nums[i])
		}
		if per == cur {
			return -1
		}
		res += 1
	}
	return res
}

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	fmt.Println(jump([]int{3,2,1,0,4}))
}
