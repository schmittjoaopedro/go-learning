package linked_list_cycle

import (
	"testing"
)

func TestExample1(t *testing.T) {
	n1 := ListNode{Val: 3}
	n2 := ListNode{Val: 2}
	n3 := ListNode{Val: 0}
	n4 := ListNode{Val: -4}
	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4
	n4.Next = &n2
	output := HasCycle(&n1)
	if !output {
		t.Fail()
	}
}

func TestExample2(t *testing.T) {
	n1 := ListNode{Val: 1}
	n2 := ListNode{Val: 2}
	n1.Next = &n2
	n2.Next = &n1
	output := HasCycle(&n1)
	if !output {
		t.Fail()
	}
}

func TestExample3(t *testing.T) {
	n1 := ListNode{Val: 1}
	output := HasCycle(&n1)
	if output {
		t.Fail()
	}
}
