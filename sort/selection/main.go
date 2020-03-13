package main

import "go_demo/common"

/**
把未排好序的数组循环一次通过比较找到最大（小）值，然后和数组最后（最前）下标交换，
一次循环找到为排好序数组的最大（小）值，并移到已排好序数组（就是数组下标--或++）
*/
// SelectionSort 选择排序
func SelectionSort(a []int) {
	length, n := len(a), len(a)
	for j := 0; j < length; j++ {
		tmp := 0

		// 循环一次找出数组最大（小）值的下标
		for i := 1; i < n; i++ {
			if a[tmp] < a[i] {
				tmp = i
			}
		}

		// 最大（小）值的下标 与 当前数组最后一位交换
		if tmp != n-1 {
			a[tmp], a[n-1] = a[n-1], a[tmp]
		}

		// n后面已经排序好
		n--

	}
}

func main() {
	a := common.GenerateSlice(100000)
	//common.PrintSlice(a, "before:")
	SelectionSort(a)
	//common.PrintSlice(a, "after:")
}
