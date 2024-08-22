package reverse_linked_list_ii

// start 19:09
// end 19:55

// Scenario 1
// head = [3,5]
// left = 1
// right = 2
// expect = [5,3]

// Scenario 2
// head = [1,2,3]
// left = 1
// right = 3
// expected = [3,2,1]

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	var nodes []*ListNode
	curr := head
	for curr != nil {
		nodes = append(nodes, curr)
		curr = curr.Next
	}

	for left < right {
		temp := nodes[left-1]
		nodes[left-1] = nodes[right-1]
		nodes[right-1] = temp
		left++
		right--
	}

	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	nodes[len(nodes)-1].Next = nil
	head = nodes[0]

	return head
}
