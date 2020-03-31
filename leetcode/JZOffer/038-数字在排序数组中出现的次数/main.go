package main

import "fmt"

func searchRange(nums []int, target int) int {
	l, r := 0, len(nums)
	res := 0
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			res++
			//	向左找
			for i := mid; nums[i] == nums[mid]; i-- {
				res++
			}
			//  向右找
			for i:=mid;nums[i]==mid;i++{
				res++
			}
			return res
		} else if nums[mid] > target{
			r = mid-1
		} else{
			l = mid+1
		}
	}
	return res
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{1,2,3,4,4,4,4,5,6,7,8}, 4))
}
