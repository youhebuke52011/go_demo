package main

import "fmt"

/**
快指针
a -> b -> c  ->  d
          |      |
          f  <-  e
slow指针走的路程是N,fast走的路程是2N,然后他们在e点相遇（fast可能已经转了几圈了）
然后fast设置为a点重新开始跑，并和slow的速度一样，所以fast走N回到e,slow在环走N部也会到e,所以他们肯定会在c点相遇，因为c->e的路程一样。
*/

func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	fast = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

func main() {
	fmt.Println(findDuplicate([]int{3, 1, 3, 4, 2}))
}
