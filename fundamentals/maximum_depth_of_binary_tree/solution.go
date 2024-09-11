package maximum_depth_of_binary_tree

// https://leetcode.com/problems/maximum-depth-of-binary-tree
// Start: 16:00
// End: 16:07

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return deepSearch(root, 1)
}

func deepSearch(node *TreeNode, cnt int) int {
	ls, rs := 0, 0
	if node.Left != nil {
		ls = deepSearch(node.Left, cnt)
	}
	if node.Right != nil {
		rs = deepSearch(node.Right, cnt)
	}
	if ls == 0 && rs == 0 {
		return cnt
	} else {
		return cnt + max(ls, rs)
	}
}
