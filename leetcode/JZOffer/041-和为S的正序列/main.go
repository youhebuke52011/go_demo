package main

import "fmt"

/**
输出所有和为S的连续正数序列。序列内按照从小至大的顺序，序列间按照开始数字从小到大的顺序。
思路：
*/

func arraySum(target int) [][]int {
	start, end := 1, 2
	res := [][]int{}
	var curSum int
	for start < end {
		curSum = (start + end) * (end - start + 1) / 2
		if curSum == target {
			tmp := []int{}
			for i:=start;i<=end;i++{
				tmp = append(tmp, i)
			}
			res = append(res, tmp)
			start++
		} else if curSum < target {
			end++
		} else {
			start++
		}
	}
	return res
}

func main() {
	fmt.Println(arraySum(100))
}
