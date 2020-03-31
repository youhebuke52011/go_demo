package main

import (
	"fmt"
	"strings"
)

func WordPattern(pattern, str string) bool {
	tmpMap := make(map[byte]string, len(pattern))
	tMap := make(map[string]byte, len(pattern))
	splitStr := strings.Split(str, " ")
	for i, row := range pattern {
		if tmp, ok := tmpMap[byte(row)]; ok {
			if tmp != splitStr[i] {
				return false
			}
		}
		tmpMap[byte(row)] = splitStr[i]

		if tmp, ok := tMap[splitStr[i]]; ok {
			if tmp != byte(row) {
				return false
			}
		}
		tMap[splitStr[i]] = byte(row)

	}

	return true
}

func main() {
	fmt.Println(WordPattern("abba", "dog cat cat dog"))
	fmt.Println(WordPattern("abba", "dog cat cat fish"))
	fmt.Println(WordPattern("aaaa", "dog cat cat dog"))
	fmt.Println(WordPattern("abba", "dog dog dog dog"))
}
