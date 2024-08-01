package group_anagrams

import "testing"

func TestExample1(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	output := GroupAnagrams(strs)
	expected := [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}
	if !equal(output, expected) {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}

func TestExample2(t *testing.T) {
	strs := []string{""}
	output := GroupAnagrams(strs)
	expected := [][]string{{""}}
	if !equal(output, expected) {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}

func TestExample3(t *testing.T) {
	strs := []string{"a"}
	output := GroupAnagrams(strs)
	expected := [][]string{{"a"}}
	if !equal(output, expected) {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}

func equal(output [][]string, expected [][]string) bool {
	if len(output) != len(expected) {
		return false
	}
	for _, groupExpected := range expected {
		found := false
		for _, groupOutput := range output {
			if len(groupExpected) != len(groupOutput) {
				continue
			}
			count := 0
			for _, strExpected := range groupExpected {
				for _, strOutput := range groupOutput {
					if strExpected == strOutput {
						count++
						break
					}
				}
			}
			if count == len(groupExpected) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
