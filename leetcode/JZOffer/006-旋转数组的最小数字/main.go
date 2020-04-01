package main

import "fmt"

/**
题目:题目描述：
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。 输入一个非递减排序的数组的一个旋转，输出旋转数组的最小元素。
例如数组{3,4,5,1,2}为{1,2,3,4,5}的一个旋转，该数组的最小值为1。 NOTE：给出的所有元素都大于0，若数组大小为0，请返回0。
变种题：leetcode:33

思路：旋转了nums[0]>nums[length-1],否则就是没旋转；如果nums[mid]>nums[length-1];说明nums[:mid+1]是有序递增的；那最小元素肯定在右边,left=mid+1，
否则right=mid
*/

func indexOfMin(nums []int) int {
	length := len(nums)
	l, r := 0, length-1
	if nums[l] < nums[r] {
		return nums[0]
	}
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]
}

func main() {
	fmt.Println(indexOfMin([]int{3, 4, 5, 1, 2}))
}
