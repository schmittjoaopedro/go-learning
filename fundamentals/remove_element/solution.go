package remove_element

// https://leetcode.com/problems/remove-element

func RemoveElement(nums []int, val int) int {
	fwd := 0
	bwd := len(nums) - 1
	k := 0
	for fwd <= bwd {
		if nums[fwd] == val {
			for fwd < bwd && nums[fwd] == nums[bwd] {
				bwd--
			}
			if fwd < bwd {
				temp := nums[bwd]
				nums[bwd] = nums[fwd]
				nums[fwd] = temp
				k++
			}
		} else {
			k++
		}
		fwd++
	}
	return k
}
