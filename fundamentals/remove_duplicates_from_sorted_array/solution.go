package remove_duplicates_from_sorted_array

// https://leetcode.com/problems/remove-duplicates-from-sorted-array
// Started: 5:15pm
// Finished: 5:29pm

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}

	curVal := nums[0]
	insIdx := 1
	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != curVal {
			temp := nums[insIdx]
			nums[insIdx] = nums[i]
			nums[i] = temp
			curVal = nums[insIdx]
			insIdx++
			k++
		}
	}
	return k
}
