package evaluate_reverse_polish_notation

import (
	"testing"
)

func TestExample1(t *testing.T) {
	tokens := []string{"2", "1", "+", "3", "*"}
	output := EvalRPN(tokens)
	if output != 9 {
		t.Fail()
	}
}

func TestExample2(t *testing.T) {
	tokens := []string{"4", "13", "5", "/", "+"}
	output := EvalRPN(tokens)
	if output != 6 {
		t.Fail()
	}
}

func TestExample3(t *testing.T) {
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	output := EvalRPN(tokens)
	if output != 22 {
		t.Fail()
	}
}
