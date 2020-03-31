package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
二叉树最小深度
 */

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDepth(root *TreeNode) int {
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 1
		}
		return min(dfs(node.Left)+1, dfs(node.Right)+1)
	}
	return dfs(root)
}

func main() {

}
