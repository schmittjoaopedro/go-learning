package merge_sorted_array

import "testing"

func TestExample1(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	Merge(nums1, m, nums2, n)
	expected := []int{1, 2, 2, 3, 5, 6}
	if !equal(nums1, expected) {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, nums1)
	}
}

func TestExample2(t *testing.T) {
	nums1 := []int{1}
	m := 1
	var nums2 []int
	n := 0
	Merge(nums1, m, nums2, n)
	expected := []int{1}
	if !equal(nums1, expected) {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, nums1)
	}
}

func TestExample3(t *testing.T) {
	nums1 := []int{0}
	m := 0
	nums2 := []int{1}
	n := 1
	Merge(nums1, m, nums2, n)
	expected := []int{1}
	if !equal(nums1, expected) {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, nums1)
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
