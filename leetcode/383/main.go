package main

import "fmt"

/**
判断第一个字符串ransom能不能由第二个字符串magazines里面的字符构成
 */

func canConstruct(ransomNote string, magazine string) bool {
	n1 := len(ransomNote)
	n2 := len(magazine)
	if n2 < n1 {
		return false
	}
	hashMap := map[byte]int{}
	sb1 := []byte(ransomNote)
	sb2 := []byte(magazine)
	for _, row := range sb2 {
		hashMap[row] += 1
	}
	for _, row := range sb1 {
		if v, ok := hashMap[row]; ok {
			if v == 0 {
				return false
			} else if v > 0 {
				hashMap[row] -= 1
			}
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aa", "aab"))
}
