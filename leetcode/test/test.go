package main

import (
	"fmt"
	"sort"
)

/**
题目：对输入的数组进行排序，返回的是，从小到大数组对应于原数组的下标（序号） 。比如 1 7 3 2 9  返回：0 3 2 1 4
 */

type T struct {
	value, index int
}

type Y []T

func (t Y) Len() int {
	return len(t)
}

func (t Y) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t Y) Less(i, j int) bool {
	if t[i].value < t[j].value {
		return true
	}
	return false
}

func main() {
	// 6,2,5,0
	tmp := Y{
		{6, 0},
		{2, 1},
		{5, 2},
		{0, 3},
	}
	sort.Sort(tmp)
	for _, row := range tmp {
		fmt.Printf("%d ", row.value)
	}
	fmt.Println()
	for _, row := range tmp {
		fmt.Printf("%d ", row.index)
	}
}
