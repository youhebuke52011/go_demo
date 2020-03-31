package main

import "fmt"

/**
两个数组的交集
 */

func intersection(nums1 []int, nums2 []int) []int {
	res := []int{}
	if len(nums2) == 0 || len(nums1) == 0 {
		return res
	}
	tmp := map[int]int{}
	resTmp := map[int]bool{}
	for i := 0; i < len(nums1); i++ {
		tmp[nums1[i]]++
	}
	for j := 0; j < len(nums2); j++ {
		if tmp[nums2[j]] > 0 {
			//tmp[nums2[j]]++
			resTmp[nums2[j]] = true
		}
	}
	for k, _ := range resTmp {
		res = append(res, k)
	}
	return res
}

func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
	fmt.Println(intersection([]int{}, []int{1, 1}))
	fmt.Println(intersection([]int{2}, []int{1, 1}))
}
