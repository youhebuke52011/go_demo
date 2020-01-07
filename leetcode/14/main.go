package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 0{
		return ""
	}
	minL := 9999
	for _, row := range strs {
		if minL > len(row) {
			minL = len(row)
		}
	}
	result := []byte{}
	outer:
		for j := 0; j < minL; j++ {
			tmp := strs[0][j]
			for _, row := range strs {
				if tmp != row[j] {
					break outer
				}
			}
			result = append(result, tmp)
		}
	t := string(result)
	return t
}

func main() {
	//s := []string{"dog", "racecar", "car"}
	//s := []string{"flower","flow","flight"}
	//s := []string{""}
	s := []string{}
	fmt.Println(longestCommonPrefix(s))
}
