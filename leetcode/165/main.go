package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
比较版本号
如果 version1 > version2 返回 1，如果 version1 < version2 返回 -1， 除此之外返回 0。
*/

func compareVersion(version1 string, version2 string) int {
	sb1 := strings.Split(version1, ".")
	sb2 := strings.Split(version2, ".")
	n := len(sb1)
	if n > len(sb2) {
		n = len(sb2)
	}
	var v1, v2 int
	i := 0
	for ; i < n; i++ {
		v1, _ = strconv.Atoi(sb1[i])
		v2, _ = strconv.Atoi(sb2[i])

		if v1 > v2 {
			return 1
		} else if v1 < v2 {
			return -1
		}
	}
	if len(sb1) > len(sb2) {
		for ; i < len(sb1); i++ {
			v1, _ = strconv.Atoi(sb1[i])
			if v1 != 0 {
				break
			}
		}
		if i == len(sb1) {
			return 0
		}
		return 1
	} else if len(sb1) < len(sb2) {
		for ; i < len(sb2); i++ {
			v2, _ = strconv.Atoi(sb2[i])
			if v2 != 0 {
				break
			}
		}
		if i == len(sb2) {
			return 0
		}
		return -1
	} else {
		return 0
	}
}

func main() {
	fmt.Println(compareVersion("1.0.1", "1.0.1"))
	fmt.Println(compareVersion("7.5.2.4", "7.5.3"))
	fmt.Println(compareVersion("01", "1"))
	fmt.Println(compareVersion("1", "1.1"))
	fmt.Println(compareVersion("1.0", "1"))
}
