package main

import "fmt"

/**
二叉搜索树：son.left<father<son.right
中序遍历可以找到第k个节点
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func KthNode(root *TreeNode, k int) *TreeNode {
	if root == nil || k == 0 {
		return nil
	}
	var res *TreeNode
	found := false
	var inOrder func(*TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil || found {
			return
		}
		inOrder(node.Left)
		if k == 1 {
			res = node
			found = true
		}
		k--
		inOrder(node.Right)
	}
	inOrder(root)
	return res
}

func print(root *TreeNode) {
	if root == nil {
		return
	}
	print(root.Left)
	fmt.Printf("%d -> ", root.Val)
	print(root.Right)
}

func main() {
	root := &TreeNode{5, &TreeNode{3, &TreeNode{1, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{7, &TreeNode{6, nil, nil}, &TreeNode{8, nil, nil}}}
	print(root)
	fmt.Println()

	res := KthNode(root, 3)
	fmt.Println(res.Val)

}
