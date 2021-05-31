package main

import "fmt"

func uniquePaths(m int, n int) int {
	if m == 0 && n == 0 {
		return 0
	}
	if m == 0 || n == 0 {
		return 1
	}
	dp := [101][101]int{}
	for i := 0; i <= m-1; i++ {
		dp[i][0] = 1
	}
	for i := 0; i <= n-1; i++ {
		dp[0][i] = 1
	}
	// dp[m][n] = dp[m-1][n] + dp[m][n-1]
	for i := 1; i <= m-1; i++ {
		for j := 1; j <= n-1; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
			fmt.Println(i, j, dp[i][j])
		}
	}
	// for i := 1; i <= m-1; i++ {
	// 	for j := 1; j <= n-1; j++ {
	// 		fmt.Printf("%d ", dp[i][j])
	// 	fmt.Println()
	// 	}
	// }
	return dp[m-1][n-1]
}

func main() {
	// fmt.Println(uniquePaths(3, 3))
	// fmt.Println(uniquePaths(3, 7))
	fmt.Println(uniquePaths(1, 1))
	// fmt.Println(uniquePaths(2, 2))
}
