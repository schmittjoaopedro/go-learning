package longest_substring_without_repeating_characters

import "testing"

func TestExample1(t *testing.T) {
	s := "abcabcbb"
	output := LengthOfLongestSubstring(s)
	expected := 3
	if output != expected {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}

func TestExample2(t *testing.T) {
	s := "bbbbb"
	output := LengthOfLongestSubstring(s)
	expected := 1
	if output != expected {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}

func TestExample3(t *testing.T) {
	s := "pwwkew"
	output := LengthOfLongestSubstring(s)
	expected := 3
	if output != expected {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}
