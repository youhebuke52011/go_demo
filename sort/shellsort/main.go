package main

import (
	"fmt"
	"go_demo/common"
)

// shellSort 希尔排序
func shellSort(a []int) {
	var (
		n    = len(a)
		gaps = []int{1}
		k    = 1
	)

	for {
		gap := element(2, k) + 1
		if gap > n-1 {
			break
		}
		gaps = append([]int{gap}, gaps...)
		k++
	}
	fmt.Println(gaps)
}

func element(a, b int) int {
	e := 1
	for b > 0 {
		if b&1 != 0 {
			e *= a
		}
		b >>= 1
		a *= a
	}
	return e
}

func main() {
	a := common.GenerateSlice(10)
	shellSort(a)
}
