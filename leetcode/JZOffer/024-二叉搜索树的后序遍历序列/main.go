package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// var ch chan int

var ch = make(chan int)

//func print() []int {
//	res := []int{}
//	res = append(res, <-ch)
//	return res
//}

func postorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	postorderTraversal(root.Left)
	ch <- root.Val
	// fmt.Println(root.Val)
	postorderTraversal(root.Right)
}

func main() {
	root := TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	res := []int{}
	go func() {
		for {
			select {
			case t := <-ch:
				res = append(res, t)
			}
		}
	}()
	postorderTraversal(&root)
	close(ch)
	fmt.Println(res)
}
