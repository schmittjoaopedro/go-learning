package remove_duplicates_from_sorted_array_2

import (
	"slices"
	"testing"
)

func TestExample1(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := RemoveDuplicates(nums)
	expected := []int{1, 1, 2, 2, 3}
	output := make([]int, k)
	copy(output, nums[:k])
	// Sort output in ascending order
	slices.Sort(output)
	if !equal(output, expected) {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}

func TestExample2(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	k := RemoveDuplicates(nums)
	expected := []int{0, 0, 1, 1, 2, 3, 3}
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
