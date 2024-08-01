package ransom_note

import "testing"

func TestExample1(t *testing.T) {
	ransomNote := "a"
	magazine := "b"
	output := CanConstruct(ransomNote, magazine)
	if output {
		t.Errorf("Test failed, expected: '%v', got: '%v'", false, output)
	}
}

func TestExample2(t *testing.T) {
	ransomNote := "aa"
	magazine := "ab"
	output := CanConstruct(ransomNote, magazine)
	if output {
		t.Errorf("Test failed, expected: '%v', got: '%v'", false, output)
	}
}

func TestExample3(t *testing.T) {
	ransomNote := "aa"
	magazine := "aab"
	output := CanConstruct(ransomNote, magazine)
	if !output {
		t.Errorf("Test failed, expected: '%v', got: '%v'", true, output)
	}
}
