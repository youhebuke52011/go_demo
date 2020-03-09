package main

import "fmt"

func solve(nums, tmp []int, n, index int, res *[][]int, visit []bool) {
	if index == n {
		t := make([]int, n)
		copy(t, tmp)
		*res = append(*res, t)
		//*res = append(*res, tmp)
		return
	}
	used := make(map[int]bool, n-index)
	for i := 0; i < n; i++ {
		if !visit[i] && !used[nums[i]] {
			visit[i] = true
			used[nums[i]] = true
			tmp[index] = nums[i]
			solve(nums, tmp, n, index+1, res, visit)
			visit[i] = false
		}
	}

}

func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	solve(nums, make([]int, len(nums)), len(nums), 0, &res, make([]bool, len(nums)))
	return res
}

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}
