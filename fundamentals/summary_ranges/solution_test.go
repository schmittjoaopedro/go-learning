package summary_ranges

import "testing"

func TestExample1(t *testing.T) {
	nums := []int{0, 1, 2, 4, 5, 7}
	output := SummaryRanges(nums)
	expected := []string{"0->2", "4->5", "7"}
	if !equal(output, expected) {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}

func TestExample2(t *testing.T) {
	nums := []int{0, 2, 3, 4, 6, 8, 9}
	output := SummaryRanges(nums)
	expected := []string{"0", "2->4", "6", "8->9"}
	if !equal(output, expected) {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}

func equal(a1 []string, a2 []string) bool {
	if len(a1) != len(a2) {
		return false
	}
	for i := range a1 {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}
