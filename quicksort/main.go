package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func generateSlice(size int) []int {
	result := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < size; i++ {
		result[i] = rand.Intn(999) - rand.Intn(999)
	}
	return result
}

func printSlice(a []int, message string) {
	fmt.Println(message)
	for _, row := range a {
		fmt.Printf(strconv.Itoa(row) + " ")
	}
	fmt.Println()
}

func quickSort(a []int) []int {
	length := len(a)
	if length < 2 {
		return a
	}
	left, right := 0, length-1

	ranVal := rand.Intn(length+1)
	// 随机选个
	index := ranVal % length

	a[index], a[right] = a[right], a[index]

	// 循环后a[index]左边都比它小
	for i, row := range a {
		if row < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	//
	a[left], a[right] = a[right], a[left]
	quickSort(a[:left])
	quickSort(a[left+1:])
	return a
}

func main() {
	a := generateSlice(10)
	printSlice(a, "before:")
	quickSort(a)
	printSlice(a, "after:")
}
