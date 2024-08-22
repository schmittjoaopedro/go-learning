package merge_two_sorted_lists

// Start 18:26
// End 18:32
// https://leetcode.com/problems/merge-two-sorted-lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode = nil
	var tail *ListNode = nil
	for {
		var next *ListNode = nil
		if list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				next = list1
				list1 = list1.Next
			} else {
				next = list2
				list2 = list2.Next
			}
		} else if list1 != nil {
			next = list1
			list1 = list1.Next
		} else if list2 != nil {
			next = list2
			list2 = list2.Next
		}
		if next != nil {
			if head == nil {
				head = next
				tail = next
			} else {
				tail.Next = next
				tail = next
			}
		} else {
			break
		}
	}
	return head
}
