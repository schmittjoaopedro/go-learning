package invert_binary_tree

import "testing"

func TestExample1(t *testing.T) {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}
	output := InvertTree(root)
	if output.Val != 4 {
		t.Errorf("Test failed")
	}
	if output.Left.Val != 7 {
		t.Errorf("Test failed")
	}
	if output.Right.Val != 2 {
		t.Errorf("Test failed")
	}
	if output.Left.Left.Val != 9 {
		t.Errorf("Test failed")
	}
	if output.Left.Right.Val != 6 {
		t.Errorf("Test failed")
	}
	if output.Right.Left.Val != 3 {
		t.Errorf("Test failed")
	}
	if output.Right.Right.Val != 1 {
		t.Errorf("Test failed")
	}
}

func TestExample2(t *testing.T) {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	output := InvertTree(root)
	if output.Val != 2 {
		t.Errorf("Test failed")
	}
	if output.Left.Val != 3 {
		t.Errorf("Test failed")
	}
	if output.Right.Val != 1 {
		t.Errorf("Test failed")
	}
}

func TestExample3(t *testing.T) {
	output := InvertTree(nil)
	if output != nil {
		t.Errorf("Test failed")
	}
}
