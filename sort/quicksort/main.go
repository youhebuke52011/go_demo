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

// QuickSort3Way 三路快排
func QuickSort3Way(a []int) []int {
	n := len(a)
	if n < 2 {
		return a
	}

	// a[0...zero] < val , a[zero+1...i) == val , a[two...n-1] > val
	zero := -1
	i := 0
	two := n

	// 随机选个
	val := a[rand.Int() % n]

	for i < two {
		if a[i] < val {
			a[i], a[zero+1] = a[zero+1], a[i]
			zero++
			i++
		} else if a[i] > val {
			a[i], a[two-1] = a[two-1], a[i]
			two--
		} else {
			i++
		}
	}

	QuickSort3Way(a[:zero+1])
	QuickSort3Way(a[two:])

	return a
}

func main() {
	//a := common.GenerateSlice(10000000)
	a := []int{8,8,3,3,2,1,2,3,6,1,1}
	common.PrintSlice(a, "before:")
	//QuickSort(a)
	QuickSort3Way(a)
	common.PrintSlice(a, "after:")
}
