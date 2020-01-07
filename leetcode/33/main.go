package main

import "fmt"

// O(n)
//func search(nums []int, target int) int {
//	for i := 0; i < len(nums); i++ {
//		if target == nums[i] {
//			return i + 1
//		}
//	}
//	return -1
//}

// O(log n)
func search(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > nums[r] {
			//	nums[l:mid]有序(4,5,6,7,0,1,2)
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			//	nums[mid:r]有序(4,5,6,7,0,1,2)
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 5))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 2))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}
