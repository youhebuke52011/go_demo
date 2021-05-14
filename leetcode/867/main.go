package main

import "fmt"

func transpose(matrix [][]int) [][]int {
	res := make([][]int, len(matrix[0]))
	for i, _ := range res {
		res[i] = make([]int, len(matrix))
	}
	fmt.Println(res)
	for i, row := range matrix {
		for j, val := range row {
			res[j][i] = val
		}
	}
	return res
}

// transposeS90 顺时针90度
func transposeS90(matrix [][]int) [][]int {
	// res := make([][]int, len(matrix[0]))
	// for i, _ := range res {
	// 	res[i] = make([]int, len(matrix))
	// }
	// // fmt.Println(res)
	// for i, row := range matrix {
	// 	for j, val := range row {
	// 		res[j][len(row)-1-i] = val
	// 	}
	// }
	// return res

	for i, _ := range matrix {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i, row := range matrix {
		for j := 0; j < len(row)/2; j++ {
			matrix[i][j], matrix[i][len(row)-1-j] = matrix[i][len(row)-1-j], matrix[i][j]
		}
	}
	return matrix
}

// transposeN90 逆时针90度
func transposeN90(matrix [][]int) [][]int {
	// res := make([][]int, len(matrix[0]))
	// for i, _ := range res {
	// 	res[i] = make([]int, len(matrix))
	// }
	// // fmt.Println(res)
	// for i, row := range matrix {
	// 	for j, val := range row {
	// 		res[len(row)-1-j][i] = val
	// 	}
	// }
	// return res

	for j, _ := range matrix {
		for i := 0; i < j; i++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for j, row := range matrix {
		for i := 0; i < len(row)/2; i++ {
			matrix[i][j], matrix[i][len(row)-1-j] = matrix[i][len(row)-1-j], matrix[i][j]
		}
	}
	return matrix
}

func main() {
	a := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	// res := transpose(a)
	// res := transposeS90(a)
	res := transposeN90(a)
	for _, row := range res {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}
