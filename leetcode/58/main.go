package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	split := strings.Split(s, " ")
	for i:=len(split)-1;i>=0;i--  {
		if len(split[i]) != 0 {
			return len(split[i])
		}
	}
	return 0
}

func main() {
	//res := lengthOfLastWord("Hello World")
	//res := lengthOfLastWord("a ")
	res := lengthOfLastWord(" ")
	fmt.Println(res)
}
