package main

import "fmt"

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := [200][200]int{}

	dp[0][0] = grid[0][0]
	for i, row := range grid {
		if i == 0 {
			continue
		}
		dp[i][0] = dp[i-1][0] + row[0]
	}
	for j, val := range grid[0] {
		if j == 0 {
			continue
		}
		dp[0][j] = dp[0][j-1] + val
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = Min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	// tmp = min(dp[i-1][j] ,dp[i][j-1]), dp[i][j] = tmp+grid[i][j]
	return dp[m-1][n-1]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// {1, 2, 3},
	// {4, 5, 6},
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(minPathSum(grid))
}
