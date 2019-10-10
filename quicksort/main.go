package main

import (
	"go_demo/common"
	"math/rand"
)

/**
随机选个下标，通过比较，得到改下标的最终位置（即该下标左边都小于它、右边都大于它）
然后该下标左边、右边的数组继续执行上面的操作（递归下就好），数组长度==1就return
O(nlogn)
*/
// QuickSort 快速排序
func QuickSort(a []int) []int {
	length := len(a)
	if length < 2 {
		return a
	}
	left, right := 0, length-1

	//ranVal := rand.Intn(length+1)
	ranVal := rand.Int()
	// 随机选个
	index := ranVal % length

	a[index], a[right] = a[right], a[index]

	// 循环后a[index]左边都比它小
	for i, row := range a {
		if row < a[right] {
			//if i != left {
			a[i], a[left] = a[left], a[i]
			//}
			left++
		}
	}

	//
	a[left], a[right] = a[right], a[left]
	QuickSort(a[:left])
	QuickSort(a[left+1:])
	return a
}

func QuickSort2(a []int) []int {
	length := len(a)
	if length < 2 {
		return a
	}
	// a[l...lt) < val , a[lt...i) == val , a[rt...r] > val

	return a
}

func main() {
	a := common.GenerateSlice(10000000)
	//a := []int{7,4,8,1,2,9,3,5,6}
	//printSlice(a, "before:")
	QuickSort(a)
	//printSlice(a, "after:")
}
