package main

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
	nums := []int{1, 1, 1, 2, 2, 3}
	removeDuplicates(nums)
}
