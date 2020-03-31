package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
路径总和
是否一条路径=sum
 */

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	sum -= root.Val
	if root.Left == nil && root.Right == nil && sum == 0 {
		return true
	}

	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)

}

func main() {
}
