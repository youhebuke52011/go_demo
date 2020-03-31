package main

import "fmt"

/**
逆序对数
*/

var count int

func merge(a, b []int) []int {
	res := make([]int, len(a)+len(b))
	i, j, k := 0, 0, 0
	for i < len(a) && j < len(b) {
		if a[i] > b[j] {
			res[k] = a[i]
			i++
			count+=len(b) - 1 - j + 1
		} else {
			res[k] = b[j]
			j++
		}
		k++
	}
	for i < len(a) {
		res[k] = a[i]
		k++
		i++
	}
	for j < len(b) {
		res[k] = b[j]
		k++
		j++
	}
	return res
}

func InversePairsCore(nums []int) []int {
	length := len(nums)
	if length < 2 {
		return nums
	}
	l, r := 0, length
	mid := l + (r-l)/2
	return merge(InversePairsCore(nums[l:mid]), InversePairsCore(nums[mid:r]))
}

func main() {
	fmt.Println(InversePairsCore([]int{9, 8, 7, 6}))
	fmt.Println(count)
}
