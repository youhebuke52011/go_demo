package main

import "fmt"

func main() {
	// "bacbaab"
	// "baebbeab"
	// s := "babad"
	// s := "ccc"
	s := "baebbeab"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	sb := []byte(s)
	dp := [1000][1000]bool{}
	start, length := 0, 1
	for i, row := range sb {
		dp[i][i] = true
		if i+1 <= len(sb)-1 && row == sb[i+1] {
			dp[i][i+1] = true
			start = i
			length = 2
		}
	}
	// fmt.Println(dp)

	for l := 3; l <= len(sb); l++ {
		// fmt.Println("-----------------")
		for i := 0; i+l-1 < len(sb); i++ {
			j := i + l - 1
			fmt.Println(i, j)
			if sb[i] == sb[j] && dp[i+1][j-1] {
				dp[i][j] = true
				start = i
				length = l
			}
		}
	}

	fmt.Println(start, length)
	return string(sb[start : start+length])
}
