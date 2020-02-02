package main

import "fmt"

/**
题目: 在一个二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
思路: 从矩阵右上角(左下角也可以的)出发,如果current>target 列向左移动,如果相等直接return,如果current<target 行向下移动
*/

func twoDimensionalMatrix(a [][]int, target int) bool {
	var ln int
	rn := len(a)
	if rn == 0 {
		return false
	}
	ln = len(a[0])
	for i := 0; i < rn; i++ {
		for j := ln - 1; j >= 0; j-- {
			if a[i][j] == target {
				return true
			} else if a[i][j] > target {
				ln -= 1
				continue
			} else {
				break
			}
		}
	}
	return false
}

func main() {
	a := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}
	fmt.Println(twoDimensionalMatrix(a, 6))
}
