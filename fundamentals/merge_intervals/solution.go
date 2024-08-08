package merge_intervals

// Start 18:07
// End 18:52
// https://leetcode.com/problems/merge-intervals

// Scenario 1
// Input: [[1,4],[0,4]]
// Expected: [[0,4]]

// Scenario 2
// Input: [[2,3],[4,5],[6,7],[8,9],[1,10]]
// Expected: [[1,10]]

import (
	"sort"
)

func Merge(intervals [][]int) [][]int {
	merged := [][]int{}
	size := len(intervals)
	if size > 0 {
		sort.Slice(intervals, func(i, j int) bool {
			if intervals[i][0] == intervals[j][0] {
				return intervals[i][1] < intervals[j][1]
			} else {
				return intervals[i][0] < intervals[j][0]
			}
		})
		start := intervals[0][0]
		end := intervals[0][1]
		for i := 0; i < size; i++ {
			if i+1 < size && end >= intervals[i+1][0] {
				end = max(end, intervals[i+1][1])
				continue
			} else {
				merged = append(merged, []int{start, end})
				if i+1 < size {
					start = intervals[i+1][0]
					end = intervals[i+1][1]
				}
			}
		}
	}
	return merged
}
