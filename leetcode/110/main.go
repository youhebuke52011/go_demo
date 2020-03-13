package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func slove(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	leftHigh, leftIsBalanced := slove(root.Left)
	rightHigh, rightIsBalanced := slove(root.Right)

	if leftIsBalanced && rightIsBalanced && leftHigh-rightHigh >= -1 && leftHigh-rightHigh <= 1 {
		//return max(leftHigh, rightHigh) + 1, true
		return max(leftHigh, rightHigh) + 1, true
	}
	return 0, false

}

func isBalanced(root *TreeNode) bool {
	_, res := slove(root)
	return res
}

func main() {

}
