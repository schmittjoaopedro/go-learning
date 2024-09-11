package invert_binary_tree

// https://leetcode.com/problems/invert-binary-tree
// Start: 16:00
// End: 16:07
// Start: 16:18
// End: 16:21

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InvertTree(root *TreeNode) *TreeNode {
	if root != nil {
		temp := root.Left
		root.Left = root.Right
		root.Right = temp
		InvertTree(root.Left)
		InvertTree(root.Right)
	}
	return root
}
