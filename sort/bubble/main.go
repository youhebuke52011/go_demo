package main

import (
	"go_demo/common"
)

/**
主要通过两两交换，循环一次找到最大（小）值，然后循环n次，那长度为n的数组就排好序了
*/
// bubbleSort 冒泡排序
func bubbleSort(a []int) {
	sorted := false
	length := len(a)
	for !sorted {
		swap := false
		// 每一次循环通过两两交换找到最大（小）值
		for i := 0; i < length-1; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				swap = true
			}
		}
		// 没有发生交换,说明数组已有序
		if !swap {
			sorted = true
		}
		// 已找到最大（小）值 , 剔除最后一个
		length--
	}
}

func main() {
	a := common.GenerateSlice(100000)
	//common.PrintSlice(a, "before:")
	bubbleSort(a)
	//common.PrintSlice(a, "after:")
}
