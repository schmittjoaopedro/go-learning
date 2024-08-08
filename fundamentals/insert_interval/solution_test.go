package insert_interval

import "testing"

func TestExample1(t *testing.T) {
	intervals := [][]int{{1, 3}, {6, 9}}
	newInterval := []int{2, 5}
	output := Insert(intervals, newInterval)
	expected := [][]int{{1, 5}, {6, 9}}
	if !equal(output, expected) {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}

func TestExample2(t *testing.T) {
	intervals := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	newInterval := []int{4, 8}
	output := Insert(intervals, newInterval)
	expected := [][]int{{1, 2}, {3, 10}, {12, 16}}
	if !equal(output, expected) {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}

func TestExample3(t *testing.T) {
	intervals := [][]int{}
	newInterval := []int{1, 2}
	output := Insert(intervals, newInterval)
	expected := [][]int{{1, 2}}
	if !equal(output, expected) {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}

func TestExample4(t *testing.T) {
	intervals := [][]int{{1, 5}}
	newInterval := []int{6, 8}
	output := Insert(intervals, newInterval)
	expected := [][]int{{1, 5}, {6, 8}}
	if !equal(output, expected) {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}

func TestExample5(t *testing.T) {
	intervals := [][]int{{1, 5}}
	newInterval := []int{0, 3}
	output := Insert(intervals, newInterval)
	expected := [][]int{{0, 5}}
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
