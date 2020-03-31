package main

import "fmt"

/**
能不能找零钱
 */

func lemonadeChange(bills []int) bool {
	fives, tens := 0, 0
	for _, row := range bills {
		switch row {
		case 5:
			fives++
		case 10:
			tens++
			fives--
		case 20:
			if tens > 0 {
				tens--
				fives--
			} else {
				fives -= 3
			}
		}
		if fives < 0 || tens < 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(lemonadeChange([]int{5, 5, 10, 10, 20}))
	fmt.Println(lemonadeChange([]int{5, 5, 10}))
	fmt.Println(lemonadeChange([]int{5, 5, 5, 10, 20}))
}
