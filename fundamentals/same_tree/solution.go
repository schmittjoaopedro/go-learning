package same_tree

// https://leetcode.com/problems/same-tree
// Start: 16:00
// End: 16:07

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p != nil && q != nil {
		if p.Val == q.Val {
			if IsSameTree(p.Left, q.Left) {
				return IsSameTree(p.Right, q.Right)
			} else {
				return false
			}
		} else {
			return false
		}
	} else if p == nil && q == nil {
		return true
	} else {
		return false
	}
}
