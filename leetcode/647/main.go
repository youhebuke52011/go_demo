package main

import "fmt"

// func countSubstrings(s string) int {
// 	cnt := 0
// 	sb := []byte(s)
// 	for i := 0; i < len(sb); i++ {
// 		for j := i + 1; j <= len(sb); j++ {
// 			if isOk(sb[i:j]) {
// 				// fmt.Println(i, j, sb[i:j])
// 				cnt++
// 			}
// 		}
// 	}
// 	return cnt
// }

func countSubstrings(s string) int {
	cnt := 0
	dp := [1000][1000]bool{}
	sb := []byte(s)
	n := len(sb)
	for i := 0; i < n; i++ {
		dp[i][i] = true
		cnt++
	}
	for i := 0; i < n-1; i++ {
		if sb[i] == sb[i+1] {
			dp[i][i+1] = true
			cnt++
		}
	}

	for l := 3; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			j := i + l - 1
			if sb[i] == sb[j] && dp[i+1][j-1] {
				dp[i][j] = true
				cnt++
			}
		}
	}

	return cnt
}

func isOk(sb []byte) bool {
	l, r := 0, len(sb)-1
	// fmt.Println(string(sb))
	for l <= r {
		if sb[l] != sb[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func main() {
	// sb := []byte("10")
	// fmt.Println(sb[0], sb[1])
	fmt.Println(countSubstrings("aaa"))
	fmt.Println(countSubstrings("abc"))
}
