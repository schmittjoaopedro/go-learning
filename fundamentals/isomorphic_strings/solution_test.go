package isomorphic_strings

import "testing"

func TestExample1(t *testing.T) {
	sInput := "egg"
	tInput := "add"
	output := IsIsomorphic(sInput, tInput)
	if !output {
		t.Errorf("Test failed, expected: '%v', got: '%v'", true, output)
	}
}

func TestExample2(t *testing.T) {
	sInput := "foo"
	tInput := "bar"
	output := IsIsomorphic(sInput, tInput)
	if output {
		t.Errorf("Test failed, expected: '%v', got: '%v'", false, output)
	}
}

func TestExample3(t *testing.T) {
	sInput := "paper"
	tInput := "title"
	output := IsIsomorphic(sInput, tInput)
	if !output {
		t.Errorf("Test failed, expected: '%v', got: '%v'", true, output)
	}
}
