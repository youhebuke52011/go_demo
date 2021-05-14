package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i, row := range nums {
		if k, ok := hashMap[target-row]; ok {
			return []int{k, i}
		}
		hashMap[row] = i
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{2}, 2))
	fmt.Println(twoSum([]int{}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))

	a := 3
	b := 4
	fmt.Println(a, b)
	a, b = b, a
	// a = a + b
	// b = a - b
	// a = a - b
	fmt.Println(a, b)

}
