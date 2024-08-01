package word_pattern

import "testing"

func TestExample1(t *testing.T) {
	pattern := "abba"
	s := "dog cat cat dog"
	output := WordPattern(pattern, s)
	if !output {
		t.Errorf("Test failed, expected: '%v', got: '%v'", true, output)
	}
}

func TestExample2(t *testing.T) {
	pattern := "abba"
	s := "dog cat cat fish"
	output := WordPattern(pattern, s)
	if output {
		t.Errorf("Test failed, expected: '%v', got: '%v'", false, output)
	}
}

func TestExample3(t *testing.T) {
	pattern := "aaaa"
	s := "dog cat cat dog"
	output := WordPattern(pattern, s)
	if output {
		t.Errorf("Test failed, expected: '%v', got: '%v'", false, output)
	}
}
