package maximum_depth_of_binary_tree

import "testing"

func TestExample1(t *testing.T) {
	tree := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	output := MaxDepth(tree)
	expected := 3
	if output != expected {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}

func TestExample2(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
		},
	}
	output := MaxDepth(tree)
	expected := 2
	if output != expected {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}

func TestExample3(t *testing.T) {
	output := MaxDepth(nil)
	expected := 0
	if output != expected {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}
