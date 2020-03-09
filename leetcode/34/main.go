package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) == 1 && nums[0] == target {
		return []int{0, 0}
	}
	l, r := 0, len(nums)-1
	res := []int{-1, -1}
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			// 向左找最开始位置(改进，也是用二分法找)
			for i := mid; i >= 0 && nums[i] == target ; i-- {
				res[0] = i
			}

			// 向右找最末尾位置(改进，也是用二分法找)
			for j := mid; j < len(nums) && nums[j] == target ; j++ {
				res[1] = j
			}
			break
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{1}, 1))
	fmt.Println(searchRange([]int{2, 2}, 2))
	fmt.Println(searchRange([]int{1, 4}, 4))
}
