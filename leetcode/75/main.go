package main

import "fmt"

/**
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
