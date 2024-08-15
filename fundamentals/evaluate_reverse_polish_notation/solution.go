package evaluate_reverse_polish_notation

// Start 18:26
// End 18:42

import (
	"strconv"
)

func EvalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		if token == "+" {
			result := stack[len(stack)-2] + stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] = result
		} else if token == "-" {
			result := stack[len(stack)-2] - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] = result
		} else if token == "*" {
			result := stack[len(stack)-2] * stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] = result
		} else if token == "/" {
			result := stack[len(stack)-2] / stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] = result
		} else {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}
	return stack[0]
}
