package main

import "fmt"

/**
//给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
//
// 此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
三路快排思想
 */

func sortColors(nums []int) {
	i, j, k := 0, 0, len(nums)-1

	for j <= k {
		switch nums[j] {
		case 0:
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		case 1:
			j++
		case 2:
			nums[k], nums[j] = nums[j], nums[k]
			k--
		}
	}
}

func main() {
	a := []int{2, 0, 2, 1, 1, 0}
	sortColors(a)
	fmt.Println(a)
}
