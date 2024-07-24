package rotate_image

import "testing"

func TestExample1(t *testing.T) {
	nums := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	Rotate(nums)
	expected := [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}
	if !equal(nums, expected) {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, nums)
	}
}

func TestExample2(t *testing.T) {
	nums := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	Rotate(nums)
	expected := [][]int{
		{15, 13, 2, 5},
		{14, 3, 4, 1},
		{12, 6, 8, 9},
		{16, 7, 10, 11},
	}
	if !equal(nums, expected) {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, nums)
	}
}

func equal(a1 [][]int, a2 [][]int) bool {
	if len(a1) != len(a2) {
		return false
	}
	for i := 0; i < len(a1); i++ {
		if len(a1[i]) != len(a2[i]) {
			return false
		}
		for j := 0; j < len(a1[i]); j++ {
			if a1[i][j] != a2[i][j] {
				return false
			}
		}
	}
	return true
}
