package merge_two_sorted_lists

import "testing"

func TestExample1(t *testing.T) {
	list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	output := MergeTwoLists(list1, list2)
	expected := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}}
	if !equal(output, expected) {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}

func TestExample2(t *testing.T) {
	list1 := &ListNode{}
	list2 := &ListNode{}
	output := MergeTwoLists(list1, list2)
	expected := &ListNode{}
	if !equal(output, expected) {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}

func TestExample3(t *testing.T) {
	list1 := &ListNode{}
	list2 := &ListNode{0, nil}
	output := MergeTwoLists(list1, list2)
	expected := &ListNode{0, nil}
	if !equal(output, expected) {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}

func equal(output *ListNode, expected *ListNode) bool {
	for expected != nil {
		if expected.Val != output.Val {
			return false
		}
		expected = expected.Next
		output = output.Next
	}
	return true
}
