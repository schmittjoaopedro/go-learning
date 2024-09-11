package same_tree

import "testing"

func TestExample1(t *testing.T) {
	pTree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	qTree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	output := IsSameTree(pTree, qTree)
	if !output {
		t.Errorf("Test failed")
	}
}

func TestExample2(t *testing.T) {
	pTree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}
	qTree := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
		},
	}
	output := IsSameTree(pTree, qTree)
	if output {
		t.Errorf("Test failed")
	}
}

func TestExample3(t *testing.T) {
	pTree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 1,
		},
	}
	qTree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
	output := IsSameTree(pTree, qTree)
	if output {
		t.Errorf("Test failed")
	}
}
