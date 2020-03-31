package main

import "fmt"

/**
//在一条环路上有 N 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
//
// 你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
//
// 如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1。
*/

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
