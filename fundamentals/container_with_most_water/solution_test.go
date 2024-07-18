package container_with_most_water

import "testing"

func TestExample1(t *testing.T) {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	output := MaxArea(nums)
	expected := 49
	if output != expected {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}

func TestExample2(t *testing.T) {
	nums := []int{1, 1}
	output := MaxArea(nums)
	expected := 1
	if output != expected {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}
