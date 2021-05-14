package main

import "fmt"

/**
数据流中的中位数
如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。
如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。
我们使用Insert()方法读取数据流，使用GetMedian()方法获取当前读取数据的中位数。
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
