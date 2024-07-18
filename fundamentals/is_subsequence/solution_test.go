package is_subsequence

import "testing"

func TestExample1(t *testing.T) {
	s1 := "abc"
	t1 := "ahbgdc"
	output := IsSubsequence(s1, t1)
	if !output {
		t.Errorf("Test failed, expected: '%v', got: '%v'", true, output)
	}
}

func TestExample2(t *testing.T) {
	s1 := "axc"
	t1 := "ahbgdc"
	output := IsSubsequence(s1, t1)
	if output {
		t.Errorf("Test failed, expected: '%v', got: '%v'", true, output)
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
