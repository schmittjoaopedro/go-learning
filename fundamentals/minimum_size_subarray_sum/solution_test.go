package minimum_size_subarray_sum

import "testing"

func TestExample1(t *testing.T) {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	output := MinSubArrayLen(target, nums)
	expected := 2
	if output != expected {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}

func TestExample2(t *testing.T) {
	target := 4
	nums := []int{1, 4, 4}
	output := MinSubArrayLen(target, nums)
	expected := 1
	if output != expected {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}

func TestExample3(t *testing.T) {
	target := 11
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1}
	output := MinSubArrayLen(target, nums)
	expected := 0
	if output != expected {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}
