package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	// a是当前剩余多少汽油,b是欠多少汽油,开始的位置
	a, b, start := 0, 0, 0
	for i, g := range gas {
		a += g - cost[i]
		// 不够钱到下一站
		if a < 0 {
			b += a
			start = i + 1
			a = 0
		}
	}

	if a+b < 0 {
		return -1
	}
	return start
}

func main() {
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
}
