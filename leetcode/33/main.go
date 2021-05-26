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
//func search(nums []int, target int) int {
//	n := len(nums)
//	l, r := 0, n-1
//	for l <= r {
//		mid := l + (r-l)/2
//		if nums[mid] == target {
//			return mid
//		} else if nums[mid] > nums[r] {
//			//	nums[l:mid]有序(4,5,6,7,0,1,2)
//			if nums[l] <= target && target < nums[mid] {
//				r = mid - 1
//			} else {
//				l = mid + 1
//			}
//		} else {
//			//	nums[mid:r]有序(4,5,6,7,0,1,2)
//			if nums[mid] < target && target <= nums[r] {
//				l = mid + 1
//			} else {
//				r = mid - 1
//			}
//		}
//	}
//	return -1
//}

func search(nums []int, target int) int {
	rotated := indexOfMin(nums) /* 数组旋转了的距离 */
	fmt.Println(rotated)
	size := len(nums)
	left, right := 0, size-1

	for left <= right {
		mid := (left + right) / 2
		/* nums 是 rotated，所以需要使用 rotatedMid 来获取 mid 的值 */
		rotatedMid := (rotated + mid) % size
		switch {
		case nums[rotatedMid] < target:
			left = mid + 1
		case target < nums[rotatedMid]:
			right = mid - 1
		default:
			return rotatedMid
		}
	}

	return -1
}

/* nums 是被旋转了的递增数组 */
func indexOfMin(nums []int) int {
	size := len(nums)
	left, right := 0, size-1
	for left < right {
		mid := (left + right) / 2
		if nums[right] < nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	// fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 5))
	// fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 2))
	// fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}
