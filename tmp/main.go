package main

import "fmt"

type S struct {
}

func InitType() S {
	var s S
	return s
}

func InitType1() *S {
	var s S
	return &s
}

func removeDuplicates(nums []int) int {
	j := 0
	for i, row := range nums {
		if i == 0 || i == 1 || row != nums[i-2] {
			if i != 0 && i != 1 {
				j++
			}
			nums[j] = row
			j++
		}
	}
	nums = nums[:j]
	return j
}

func main() {
	//var number *int
	//*number = 10
	//fmt.Println(*number)
	//fmt.Println(InitType()==nil)
	//fmt.Println(InitType1()==nil)
	//nums := []int{1, 1, 1, 2, 2, 3}
	//removeDuplicates(nums)

	//s := "(()(()))"
	//fmt.Println(scoreOfParentheses(s))
	//s := "()()"
	//fmt.Println(s[0],s[len(s)-1])
	//for i := 0; i < len(s); i++ {
	//	fmt.Printf("%c", s[i])
	//}

	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
	//fmt.Println(partitionLabels("eaaaabaaec"))
}

func partitionLabels(S string) []int {
	result := []int{}
	tmpMap := make(map[byte]int,30)
	before, l, k := -1, 0, 0
	for i :=len(S)-1;i>=0;i--{
		if _,ok := tmpMap[S[i]];!ok{
			tmpMap[S[i]] = i
		}
	}
	for i :=0 ;i<len(S);i++ {
		if r,ok := tmpMap[S[i]];ok {
			//fmt.Printf("%c",S[i])
			if before != -1 && (i == before || i-1==before) {
				result = append(result,l)
				if len(result) != 0 {
					k = before+1
				}
				l = 0
				before = -1
			}
			if l < r - k + 1 {
				l = r - k + 1
				before = r
			}
		}
	}
	return result
}

//var a string
//var done bool
//
//func setup() {
//	a = "hello, world"
//	done = true
//}
//
//func t() {
//	go setup()
//	for !done {
//	}
//	print(a)
//}
