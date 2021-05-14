package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := len(matrix[i]) - 1; j >= 0; j-- {
			if matrix[i][j] == target {
				return true
			} else if matrix[i][j] < target {
				break
			}
		}
	}
	return false
}

// [1,3,5,7],[10,11,16,20],[23,30,34,60]

func main() {
	a := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	fmt.Println(searchMatrix(a, 24))
}
