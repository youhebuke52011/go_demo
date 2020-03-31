package main

import "fmt"

/**
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。
^int  按位取反
// 那么，如果是三次的话，香农定理，需要用 2 bits 进行记录

// 当 5 第 1 次出现的时候，ones = 5, twos = 0,  ones 记录这个数字
// 当 5 第 2 次出现的时候，ones = 0, twos = 5， twos 记录了这个数字
// 当 5 第 3 次出现的时候，ones = 0, twos = 0， 都清空了，可以去处理其他数字了
// 所以，如果有某个数字出现了 1 次，就存在 ones 中，出现了 2 次，就存在 twos 中
*/

func singleNumber(nums []int) int {
	ones, twos := 0, 0
	for i := 0; i < len(nums); i++ {
		fmt.Printf("%b  %b\n", ones^nums[i], uint(^twos))
		ones = (ones ^ nums[i]) & ^twos
		fmt.Printf("%b  %b\n", twos^nums[i], uint(^ones))
		twos = (twos ^ nums[i]) & ^ones
	}
	return ones
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 3, 2}))
}
