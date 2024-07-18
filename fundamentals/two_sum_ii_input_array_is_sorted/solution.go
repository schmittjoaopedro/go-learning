package two_sum_ii_input_array_is_sorted

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted
// Start: 19:10
// End: 19:16

func TwoSum(numbers []int, target int) []int {
	i1 := 0
	i2 := len(numbers) - 1
	for i1 < i2 {
		sum := numbers[i1] + numbers[i2]
		if sum == target {
			break
		} else if sum > target {
			i2--
		} else {
			i1++
		}
	}
	return []int{i1 + 1, i2 + 1}
}
