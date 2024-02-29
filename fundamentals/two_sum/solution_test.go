package two_sum

import "testing"

func TestExample1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	output := TwoSum(nums, target)
	expected := []int{0, 1}
	if !equal(output, expected) {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}

func TestExample2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	output := TwoSum(nums, target)
	expected := []int{1, 2}
	if !equal(output, expected) {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}

func TestExample3(t *testing.T) {
	nums := []int{3, 3}
	target := 6
	output := TwoSum(nums, target)
	expected := []int{0, 1}
	if !equal(output, expected) {
		t.Errorf("Test failed, expected '%v', got: '%v'", expected, output)
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
