package container_with_most_water

// https://leetcode.com/problems/container-with-most-water
// Start: 19:17
// End: 19:28

func MaxArea(heights []int) int {
	max_area := 0
	head := 0
	tail := len(heights) - 1

	for head < tail {
		width := tail - head
		height := min(heights[head], heights[tail])
		new_area := width * height
		max_area = max(max_area, new_area)
		if heights[head] > heights[tail] {
			tail--
		} else {
			head++
		}
	}

	return max_area
}
