package main

import "fmt"

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func addBinary(a string, b string) string {
	na := len(a)
	nb := len(b)
	n := Max(na, nb)
	sb := make([]byte, n+1)
	i := na - 1
	j := nb - 1
	t := byte(0)
	for i >= 0 && j >= 0 {
		tmp := a[i] + b[j] + t
		temp := byte(48)
		switch int(tmp) {
		case 96:
			t = 0
		case 97:
			t = 0
			temp = byte(49)
		case 98:
			t = 1
		}
		sb[n] = temp
		n--
		i--
		j--
	}
	for i >= 0 {
		tmp := a[i] + t
		temp := byte(48)
		switch int(tmp) {
		case 96:
			t = 0
		case 97:
			t = 0
			temp = byte(49)
		case 98:
			t = 1
		}
		sb[n] = temp
		n--
		i--
	}
	for j >= 0 {
		tmp := b[j] + t
		temp := byte(48)
		switch int(tmp) {
		case 96:
			t = 0
		case 97:
			t = 0
			temp = byte(49)
		case 98:
			t = 1
		}
		sb[n] = temp
		n--
		j--
	}
	if t == 1 {
		sb[n] = byte(49)
	}
	return string(sb)
}

func main() {
	fmt.Println(addBinary("1010", "1011"))
	fmt.Println(addBinary("111", "111"))
	fmt.Println(addBinary("111", "1"))
}
