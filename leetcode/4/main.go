package main

import "fmt"

/**
数据流中的中位数
 */

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)
	res := make([]int, n1+n2)
	i, j, k := 0, 0, 0
	for ; k < n1+n2; k++ {
		if j == n2 || (i < n1 && j < n2 && nums1[i] <= nums2[j]) {
			res[k] = nums1[i]
			i += 1
		} else if i == n1 || (i < n1 && j < n2 && nums1[i] > nums2[j]) {
			res[k] = nums2[j]
			j += 1
		}
	}

	if (n1+n2) % 2 == 0 {
		return float64(res[k/2]+res[k/2-1]) / 2.0
	}
	return float64(res[k/2])
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
}
