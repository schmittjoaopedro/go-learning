package remove_duplicates_from_sorted_array_2

// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii
// Started: 5:32pm

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}

	maxSawCnt := 2
	lastSawVal := nums[0]
	lastSawCnt := 1
	instPos := 1
	for i := 1; i < len(nums); i++ {
		curVal := nums[i]

		if curVal == lastSawVal && lastSawCnt < maxSawCnt {
			lastSawCnt++
			if instPos != i {
				temp := nums[i]
				nums[i] = nums[instPos]
				nums[instPos] = temp
			}
			instPos++
		} else if curVal > lastSawVal {
			lastSawVal = curVal
			lastSawCnt = 1
			if instPos != i {
				temp := nums[i]
				nums[i] = nums[instPos]
				nums[instPos] = temp
			}
			instPos++
		}
	}

	return instPos
}
