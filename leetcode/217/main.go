package main

func containsDuplicate(nums []int) bool {
	hsMap := make(map[int]bool, len(nums))
	for _, row := range nums {
		if _, ok := hsMap[row]; ok {
			return true
		}
		hsMap[row] = true
	}
	return false
}

func main() {

}
