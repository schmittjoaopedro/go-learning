package valid_palindrome

import "testing"

func TestExample1(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	output := IsPalindrome(s)
	if !output {
		t.Errorf("Test failed, expected: '%v', got: '%v'", true, output)
	}
}

func TestExample2(t *testing.T) {
	s := "race a car"
	output := IsPalindrome(s)
	if output {
		t.Errorf("Test failed, expected: '%v', got: '%v'", false, output)
	}
}

func TestExample3(t *testing.T) {
	s := " "
	output := IsPalindrome(s)
	if !output {
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
