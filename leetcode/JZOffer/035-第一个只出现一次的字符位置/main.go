package main

import "fmt"

func count(str string) int {
	hsMap := make(map[byte]int,26)
	for _,row := range str {
		hsMap[byte(row)]++
	}
	for i,row := range str {
		if count,ok := hsMap[byte(row)];ok {
			if count == 1 {
				return i
			}
		}
	}
	return -1
}

func main() {
	fmt.Println("alskdphkjsbcjkvhaklsuhdtq: ",count("alskdphkjsbcjkvhaklsuhdtq"))
	fmt.Println("aaa2bbbfff3gggdddsssr3rrwwsss4: ",count("aaa2bbbfff3gggdddsssrrrwwsss4"))
}
