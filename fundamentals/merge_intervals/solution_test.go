package merge_intervals

import "testing"

func TestExample1(t *testing.T) {
	nums := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	output := Merge(nums)
	expected := [][]int{{1, 6}, {8, 10}, {15, 18}}
	if !equal(output, expected) {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}

func TestExample2(t *testing.T) {
	nums := [][]int{{1, 4}, {4, 5}}
	output := Merge(nums)
	expected := [][]int{{1, 5}}
	if !equal(output, expected) {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}

func TestExample3(t *testing.T) {
	nums := [][]int{{1, 4}, {0, 4}}
	output := Merge(nums)
	expected := [][]int{{0, 4}}
	if !equal(output, expected) {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}

func TestExample4(t *testing.T) {
	nums := [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}
	output := Merge(nums)
	expected := [][]int{{1, 10}}
	if !equal(output, expected) {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}

func equal(a1 [][]int, a2 [][]int) bool {
	if len(a1) != len(a2) {
		return false
	}
	for i := range a1 {
		if a1[i][0] != a2[i][0] || a1[i][1] != a2[i][1] {
			return false
		}
	}
	return true
}
