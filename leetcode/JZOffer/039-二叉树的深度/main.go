package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	max := 0
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil{
			return
		}
		if max < depth {
			max = depth
		}
		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)
	}
	dfs(root, 1)
	return max
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
	fmt.Println(maxDepth(&root))
}
