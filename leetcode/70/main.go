package main

import "fmt"

func main() {
	fmt.Println(climbStairs(20))
}

func climbStairs(n int) int {
	dp := [100]int{}
	dp[1] = 1
	dp[2] = 2
	if n < 3 {
		return dp[n]
	}
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
