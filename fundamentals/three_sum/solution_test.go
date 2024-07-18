package three_sum

import "testing"

func TestExample1(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	output := ThreeSum(nums)
	expected := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	if !equal(output, expected) {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}

func TestExample2(t *testing.T) {
	nums := []int{0, 1, 1}
	output := ThreeSum(nums)
	expected := [][]int{}
	if !equal(output, expected) {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}

func TestExample3(t *testing.T) {
	nums := []int{0, 0, 0}
	output := ThreeSum(nums)
	expected := [][]int{{0, 0, 0}}
	if !equal(output, expected) {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
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
