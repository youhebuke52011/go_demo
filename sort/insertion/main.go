package main

import "go_demo/common"

/**
无序a数组->有序a1数组、无序a2数组，循环一次有序数组把无序a2[0]插入到有序数组中适当的位置，保持有序数组有序->n次循环后,有序数组a
*/
// InsertionSort 插入排序
func InsertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		tmp := a[i]
		j := i - 1
		// i前面数组有比i小(大)的话，那就整体后移
		for j >= 0 && tmp < a[j] {
			a[j+1] = a[j]
			j--
		}
		// 符合排序规则 没进j循环 不需要赋值 j+1是因为j的for循环后j是目标位前一位
		if j+1 != i {
			a[j+1] = tmp
		}
	}
}

func main() {
	a := common.GenerateSlice(100000)
	//common.PrintSlice(a, "before:")
	InsertionSort(a)
	//common.PrintSlice(a, "after:")
}
