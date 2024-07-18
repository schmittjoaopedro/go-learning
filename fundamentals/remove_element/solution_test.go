package remove_element

import (
	"slices"
	"testing"
)

func TestExample1(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	target := 3
	k := RemoveElement(nums, target)
	expected := []int{2, 2}
	if k != len(expected) {
		t.Errorf("Test failed, expected: '%v', got: '%v'", len(expected), k)
	}
	output := make([]int, k)
	copy(output, nums[:k])
	// Sort output in ascending order
	slices.Sort(output)
	if !equal(output, expected) {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}

func TestExample2(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	target := 2
	k := RemoveElement(nums, target)
	expected := []int{0, 0, 1, 3, 4}
	if k != len(expected) {
		t.Errorf("Test failed, expected: '%v', got: '%v'", len(expected), k)
	}
	output := make([]int, k)
	copy(output, nums[:k])
	// Sort output in ascending order
	slices.Sort(output)
	if !equal(output, expected) {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}

func equal(a1 []int, a2 []int) bool {
	if len(a1) != len(a2) {
		return false
	}
	for i := 0; i < len(a1); i++ {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}
