package main

import "fmt"

/**
给定一个数组和滑动窗口的大小，找出所有滑动窗口里数值的最大值。例如，如果输入数组{2,3,4,2,6,2,5,1}及滑动窗口的大小3，
那么一共存在6个滑动窗口，他们的最大值分别为{4,4,6,6,6,5}； 针对数组{2,3,4,2,6,2,5,1}的滑动窗口有以下6个： 
{[2,3,4],2,6,2,5,1}， {2,[3,4,2],6,2,5,1}， {2,3,[4,2,6],2,5,1}， {2,3,4,[2,6,2],5,1}， {2,3,4,2,[6,2,5],1}， {2,3,4,2,6,[2,5,1]}。
窗口大于数组长度的时候，返回空
*/

func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	var window []int
	for index := 0; index < len(nums); index++ {
		// 最大值已经超出窗口范围
		if index >= k && window[0] <= index-k {
			window = window[1:]
		}
		// 当前值比队列末尾大则替换末尾
		for len(window) > 0 && nums[window[len(window)-1]] < nums[index] {
			window = window[:len(window)-1]
		}
		// 添加元素进队列
		window = append(window, index)
		// 记录当前最大值
		if index >= k-1 {
			res = append(res, nums[window[0]])
		}
	}
	return res
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
