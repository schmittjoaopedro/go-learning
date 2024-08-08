package insert_interval

// Start 19:00
// End 19:34
// https://leetcode.com/problems/insert-interval

// Scenario 1
// Input: [] <- [1,2]
// Expected [[1,2]]

// Scenario 2
// Input: [[1,6]] <- [6,8]
// Expected: [[1,8]]

// Scenario 3
// Input: [[1,5]] <- [0,3]
// Expected: [[0,5]]

func Insert(intervals [][]int, newInterval []int) [][]int {
	var newIntervals [][]int
	size := len(intervals)
	insPos := size
	if size > 0 {
		for i := 0; i < size; i++ {
			if newInterval[0] <= intervals[i][0] || newInterval[0] <= intervals[i][1] {
				insPos = i
				break
			}
		}
		for i, curInterval := range intervals {
			if i < insPos {
				newIntervals = append(newIntervals, curInterval)
			} else if i == insPos {
				if newInterval[1] < curInterval[0] {
					newIntervals = append(newIntervals, newInterval)
					newIntervals = append(newIntervals, curInterval)
				} else {
					start := min(curInterval[0], newInterval[0])
					end := max(curInterval[1], newInterval[1])
					newIntervals = append(newIntervals, []int{start, end})
				}
			} else {
				lastInterval := newIntervals[len(newIntervals)-1]
				if lastInterval[1] >= curInterval[0] {
					lastInterval[1] = max(lastInterval[1], curInterval[1])
				} else {
					newIntervals = append(newIntervals, curInterval)
				}
			}
		}
	}
	if insPos == size {
		newIntervals = append(newIntervals, newInterval)
	}
	return newIntervals
}
