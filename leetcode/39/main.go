package main

import (
	"fmt"
)

// func combinationSum(candidates []int, target int) [][]int {
// 	sort.Ints(candidates)

// 	res := [][]int{}
// 	solution := []int{}
// 	cs(candidates, solution, target, &res)

// 	return res
// }

// func cs(candidates, solution []int, target int, result *[][]int) {
// 	if target == 0 {
// 		*result = append(*result, solution)
// 	}

// 	if len(candidates) == 0 || target < candidates[0] {
// 		// target < candidates[0] 因为candidates是排序好的
// 		return
// 	}

// 	// 这样处理一下的用意是，让切片的容量等于长度，以后append的时候，会分配新的底层数组
// 	// 避免多处同时对底层数组进行修改，产生错误的答案。
// 	// 可以注释掉以下语句，运行单元测试，查看错误发生。
// 	solution = solution[:len(solution):len(solution)]

// 	cs(candidates, append(solution, candidates[0]), target-candidates[0], result)

// 	cs(candidates[1:], solution, target, result)
// }

// func slove(nums, tmp []int, target, index int, res *[][]int) {

// 	if index >= len(nums) || target < 0 {
// 		return
// 	}
// 	if target == 0 {
// 		*res = append(*res, tmp)
// 	}

// 	for i := index; i < len(nums); i++ {
// 		//tmp = append(tmp, nums[i])
// 		slove(nums, append(tmp, nums[i]), target-nums[i], index+1, res)
// 		//remove(tmp, i)
// 	}

// 	////tmp = tmp[:len(tmp):len(tmp)]
// 	//slove(nums, append(tmp, nums[0]), target-nums[0], index, res)
// 	//slove(nums[1:], tmp, target, index, res)

// }

// func remove(s []int, i int) []int {
// 	s[len(s)-1], s[i] = s[i], s[len(s)-1]
// 	return s[:len(s)-1]
// }

// func combinationSum(candidates []int, target int) [][]int {
// 	sort.Ints(candidates)
// 	res := [][]int{}
// 	slove(candidates, []int{}, target, 0, &res)
// 	return res
// }

func slove(andidates []int, target, index int, tmp []int, res *[][]int) {
	// index == len(andidates)-1 &&
	if target == 0 {
		t := make([]int, len(tmp))
		copy(t, tmp)
		*res = append(*res, t)
		// 下面这样会有问题
		// *res = append(*res, tmp)
		return
	}
	// || index > len(andidates)-1
	if target < 0 {
		return
	}
	for i := index; i < len(andidates); i++ {
		tmp = append(tmp, andidates[i])
		slove(andidates, target-andidates[i], i, tmp, res)
		tmp = tmp[:len(tmp)-1]
	}
}

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	slove(candidates, target, 0, []int{}, &res)
	return res
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
	fmt.Println(combinationSum([]int{7, 3, 2}, 18))
}
