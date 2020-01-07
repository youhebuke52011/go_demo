package main

import "fmt"

func compress(chars []byte) int {
	n := len(chars)
	if n == 0 || n == 1 {
		return n
	}
	//res := make([]byte, n)
	index := 0
	var before byte
	for i := 0; i < n; {
		tmp := 1
		if i >= 0 {
			before = chars[i]
		}
		flag := false
		chars[index] = chars[i]
		index += 1
		//chars[index] = '1'
		i += 1
		for i-1 >= 0 && i < n && chars[i] == before {
			before = chars[i]
			tmp += 1
			i += 1
			flag = true
		}
		if flag {
			if tmp/10 == 0 {
				chars[index] = byte(tmp%10 + '0')
			} else {
				chars[index] = byte(tmp/10 + '0')
				index += 1
				chars[index] = byte(tmp%10 + '0')
			}
			index += 1
		}
	}
	if index < n {
		chars = chars[:index]
	}
	return index
}

func main() {
	sb := []byte("aabbccc")
	//sb := []byte("a")
	//sb := []byte("abbbbbbbbbbbb")
	fmt.Println(compress(sb))
}
