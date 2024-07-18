package minimum_size_subarray_sum

// https://leetcode.com/problems/minimum-size-subarray-sum
/*
O(n) solution
Start: 18:06
End: 18:39
*/

func MinSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if nums[0] >= target {
			return 1
		} else {
			return 0
		}
	}
	if nums[0] >= target {
		return 1
	}
	length := len(nums) + 1
	p1, p2 := 0, 1
	sum := nums[p1] + nums[p2]
	for p1 < len(nums) {
		if sum >= target {
			length = min(length, p2-p1+1)
			sum -= nums[p1]
			p1++
		} else if p2 < len(nums)-1 {
			p2++
			sum += nums[p2]
		} else {
			sum -= nums[p1]
			p1++
		}
		if length == 1 {
			break
		}
	}
	if length <= len(nums) {
		return length
	} else {
		return 0
	}
}
