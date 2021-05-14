package main

import "fmt"

// 这种方法会把相对位置打乱
//func oddFirst(nums []int) {
//	length := len(nums)
//	l, r := 0, length-1
//	for l < r {
//		for l < r && nums[l]%2 == 1 {
//			l++
//		}
//
//		for l < r && nums[r]%2 == 0 {
//			r--
//		}
//
//		nums[l], nums[r] = nums[r], nums[l]
//	}
//}

// 保持相对位置
func oddFirst(nums []int) {
	for i, j := 0, 0; i < len(nums); i++ {
		// 从左到右找到第一个偶数
		fmt.Println(i, j, nums)
		for i < len(nums) {
			if nums[i]%2 == 0 {
				break
			}
			i++
		}

		// j从i+1开始找，找到第一个奇数
		for j = i + 1; j < len(nums); j++ {
			if nums[j]%2 == 1 {
				break
			}
		}

		if j == len(nums) {
			break
		}
		tmp := nums[j]
		// 偶数位之后往后移
		for k := j - 1; k >= i; k-- {
			nums[k+1] = nums[k]
		}
		nums[i] = tmp
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	oddFirst(nums)
	fmt.Println(nums)
}
