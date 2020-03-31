package main

import "fmt"

/**
题目： 数组中只出现一次的数字,有两个数字只出现一次
思路： 通过异或去重的原理，可以找出只出现过一次的数字，两个不同就是两个数据的亦或结果，就把两个不同分别分到不同数组里面去
		怎样分就看， 把所有数字的亦或结果 取二进制最右 为1的数字，即111 -> 001,  010 -> 010,然后再用这个结果和每个数字
        & 运算，相同的分在同一个数组，不同的分开
*/

func FindNumsAppearOnce(nums []int) []int {
	r1, r2 := 0, 0
	tmp := 0
	for _, row := range nums {
		tmp ^= row
	}

	tmp = tmp & (-tmp)

	for _, row := range nums {
		if tmp&row == 0 {
			r1 ^= row
		} else {
			r2 ^= row
		}
	}
	return []int{r1, r2}
}

func main() {
	fmt.Println(FindNumsAppearOnce([]int{2, 4, 3, 6, 3, 2, 5, 5}))
	fmt.Println(FindNumsAppearOnce([]int{1, 1, 3, 4}))
	fmt.Println(FindNumsAppearOnce([]int{2, 2, 3, 2}))
}
