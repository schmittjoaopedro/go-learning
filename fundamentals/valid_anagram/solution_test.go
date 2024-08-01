package valid_anagram

import "testing"

func TestExample1(t *testing.T) {
	sInput := "anagram"
	tInput := "nagaram"
	output := IsAnagram(sInput, tInput)
	if !output {
		t.Errorf("Test failed, expected: '%v', got: '%v'", true, output)
	}
}

func TestExample2(t *testing.T) {
	sInput := "rat"
	tInput := "car"
	output := IsAnagram(sInput, tInput)
	if output {
		t.Errorf("Test failed, expected: '%v', got: '%v'", false, output)
	}
}
