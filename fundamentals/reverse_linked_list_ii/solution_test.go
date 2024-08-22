package reverse_linked_list_ii

import "testing"

func TestExample1(t *testing.T) {
	head1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	head2 := ReverseBetween(head1, 2, 4)
	expected := []int{1, 4, 3, 2, 5}
	if !equals(head2, expected) {
		t.Fail()
	}

}

func TestExample2(t *testing.T) {
	head1 := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 5,
		},
	}

	head2 := ReverseBetween(head1, 1, 2)
	expected := []int{5, 3}
	if !equals(head2, expected) {
		t.Fail()
	}

}

func TestExample3(t *testing.T) {
	head1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	head2 := ReverseBetween(head1, 1, 3)
	expected := []int{3, 2, 1}
	if !equals(head2, expected) {
		t.Fail()
	}

}

func equals(head *ListNode, vals []int) bool {
	curr := head
	for _, val := range vals {
		if curr.Val != val {
			return false
		}
		curr = curr.Next
	}
	return true
}
