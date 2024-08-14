package min_stack

import (
	"testing"
)

func TestExample1(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	if minStack.GetMin() != -3 {
		t.Fail()
	}
	minStack.Pop()
	if minStack.Top() != 0 {
		t.Fail()
	}
	if minStack.GetMin() != -2 {
		t.Fail()
	}
}
