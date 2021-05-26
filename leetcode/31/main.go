package main

import "fmt"

/**
1.从后向前找最长递增序列【i:】，翻转，二分查找翻转序列第一个大过[i-1]的 ，swap(i-1, 翻转序列第一个大过[i-1]的）
*/

func reverse(nums []int, l int) {
	r := len(nums) - 1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l += 1
		r -= 1
	}
}

func nextPermutation(nums []int) {
	length := len(nums)
	if length < 2 {
		return
	}
	i := length - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i -= 1
	}

	// 翻转从后到前的 最长递增序列
	reverse(nums, i+1)

	if i < 0 {
		// 已经是最大的
		fmt.Println(nums)
		return
	}

	// 二分查找 在翻转序列中找刚好大过i-1的值，并swap
	l, r := i, length-1
	fmt.Println(nums[:i+1])
	fmt.Println(nums[i+1:])
	target := nums[i]
	for l+1 < r {
		//mid := l + (r-l)/2
		mid := (l + r) / 2
		if target < nums[mid] {
			r = mid
		} else {
			l = mid
		}
	}
	nums[i], nums[r] = nums[r], nums[i]
	fmt.Println(nums)
}

func main() {
	// nextPermutation([]int{1, 2, 3})
	// nextPermutation([]int{3, 2, 1})
	// nextPermutation([]int{1, 1, 5})
	// nextPermutation([]int{1, 5, 1})
	// nextPermutation([]int{1, 3, 2})
	nextPermutation([]int{2, 2, 7, 5, 4, 3, 2, 2, 1})
}
