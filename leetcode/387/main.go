package main

import "fmt"

func firstUniqChar(s string) int {
	sb := []byte(s)
	hashMap := map[byte]int{}

	for _, row := range sb {
		hashMap[row] += 1
	}

	for i, row := range sb {
		if v, ok := hashMap[row]; ok {
			if v == 1 {
				return i
			} else if v > 1 {
				continue
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))
}
