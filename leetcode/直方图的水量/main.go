package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func trap(height []int) int {
	if len(height) < 2 {
		return 0
	}
	l, r := 0, len(height)-1
	tmpTrap := min(l, r) * (r - l)
	for l < r {
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
		tmpTrap = max(tmpTrap, min(height[l], height[r])*(r-l))
	}
	return tmpTrap
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
