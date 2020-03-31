package main

import (
	"fmt"
)

/**
//给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
用一个值来标识, 质数*每个字符
*/

func groupAnagrams(strs []string) [][]string {
	resMap := map[int][]string{}
	su := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103}
	res := [][]string{}
	for _, row := range strs {
		sb := []byte(row)
		weight := 0
		for j := 0; j < len(sb); j++ {
			//weight += int(sb[j])
			weight += su[int(sb[j])%int('a')] * int(sb[j])
		}
		key := weight
		resMap[key] = append(resMap[key], row)
	}
	for _, v := range resMap {
		res = append(res, v)
	}
	return res
}

func main() {
	//s := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	//s := []string{"cab","tin","pew","duh","may","ill","buy","bar","max","doc"}
	//s := []string{"ray", "cod", "abe", "ned", "arc", "jar", "owl", "pop", "paw", "sky", "yup", "fed", "jul", "woo", "ado", "why", "ben", "mys", "den", "dem", "fat", "you", "eon", "sui", "oct", "asp", "ago", "lea", "sow", "hus", "fee", "yup", "eve", "red", "flo", "ids", "tic", "pup", "hag", "ito", "zoo"}
	s := []string{"pal","ugh","boy","sir","nay","sap","zoe","tan","pym","tho","lea","hod","vim","mes","qua","lag","jib","meg","guy","tat","hue","fez","zit","owe","maj","die","dos","bus","dew","yum","gos","deb","pis","mod","poi","hug","ado","dot","fob","tee","oil","bah","wok","sip","eve","dee","bob","cid","han","mit","ram","wis","ark","pry","mys","wag","eng"}
	res := groupAnagrams(s)
	fmt.Println(res)
}
