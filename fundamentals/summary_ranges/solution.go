package summary_ranges

// Start 17:45
// End 18:02
// https://leetcode.com/problems/summary-ranges

import "fmt"

func SummaryRanges(nums []int) []string {
	ranges := []string{}

	if len(nums) > 0 {
		start := nums[0]
		for i, val := range nums {
			if i+1 < len(nums) && val+1 == nums[i+1] {
				continue
			} else {
				if start == val {
					ranges = append(ranges, fmt.Sprintf("%d", start))
				} else {
					ranges = append(ranges, fmt.Sprintf("%d->%d", start, val))
				}
				if i+1 < len(nums) {
					start = nums[i+1]
				}
			}
		}
	}
	return ranges
}
