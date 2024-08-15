package linked_list_cycle

// Start 18:47
// End 19:05

type ListNode struct {
	Val  int
	Next *ListNode
}

func HasCycle(head *ListNode) bool {
	visited := map[*ListNode]bool{}
	curr := head
	for curr != nil {
		_, ok := visited[curr]
		if ok {
			return true
		} else {
			visited[curr] = true
		}
		curr = curr.Next
	}
	return false
}
