package common

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GenerateSlice(size int) []int {
	result := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < size; i++ {
		result[i] = rand.Intn(999) - rand.Intn(999)
	}
	return result
}

func PrintSlice(a []int, message string) {
	fmt.Println(message)
	for _, row := range a {
		fmt.Printf(strconv.Itoa(row) + " ")
	}
	fmt.Println()
}
