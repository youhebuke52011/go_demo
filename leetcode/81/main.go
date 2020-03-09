package main

import "fmt"

func search(nums []int, target int) bool {
	length := len(nums)
	l, r := 0, length-1
	k := 1
	for k < length && nums[k-1] <= nums[k] {
		k++
	}

	for l <= r {
		mid := l + (r-l)/2
		med := (mid + k) % length
		if nums[med] == target {
			return true
		} else if nums[med] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}

func main() {
	//fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
	//fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 5))
	//fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 2))
	//fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 11))
	//fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 1))
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 6))
}
