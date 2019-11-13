package main

import (
	"fmt"
	"sort"
)



type IntSlice []int

func (s IntSlice) Swap(i,j int) {
	s[i],s[j] = s[j],s[i]
}

func (s IntSlice) Len() int {
	return len(s)
}

func (s IntSlice) Less(i,j int) bool {
	return s[i] < s[j]
}

func isNStraightHand(hand []int, W int) bool {
	if len(hand) % W != 0 {
		return false
	}
	sort.Sort(IntSlice(hand))
	tmpMap := map[int]int{}
	minN, maxN := 1<<32-1,0
	for _,row := range hand{
		if minN > row {
			minN = row
		}
		if maxN < row {
			maxN = row
		}
		tmpMap[row] ++
	}
	tmpSlice := []int{}
	before := 0
	flag := false
	count := 0
	tmp := 0
	for i:=minN;i<=maxN;{
		if _,ok:=tmpMap[i];!ok || tmpMap[i] == 0{
			tmp ++
			if tmp >= 10 {
				for k,row := range hand{
					if k<len(hand)-1 && row == tmpSlice[len(tmpSlice)-1] {
						i += hand[k+1]
						break
					}
				}
			}
			continue
		}
		tmp = 0
		tmpMap[i] --
		if !flag && tmpMap[i] != 0 {
			before = i
			flag = true
		}
		tmpSlice = append(tmpSlice,i)
		if len(tmpSlice) == W {
			for j,row := range tmpSlice{
				if j!=0 && row-1 != tmpSlice[j-1] {
					return false
				}
			}
			tmpSlice = []int{}
			if before != 0{
				i = before - 1
			}
			flag = false
			count++
		}
		i++
	}
	if count == int(len(hand) / W){
		return true
	}
	return false
}




func main() {

	fmt.Println(isNStraightHand([]int{1,2,3,6,2,3,4,7,8},3))
	//fmt.Println(isNStraightHand([]int{986335089,694522119,762403981,769866323,522519161,245109154,678349141,344764950,616558613,515692211},5))
	//fmt.Println(isNStraightHand([]int{238859626,731251114,722361905,684867006,12086001,585836946,570894279,887859796,747750025,937898166},5))
	//fmt.Println(isNStraightHand([]int{1,2,3,4,5,6},2))
	//fmt.Println(isNStraightHand([]int{1,2,3,6,2,3,4,7,8},3))
	//fmt.Println(isNStraightHand([]int{5,1},2))
}