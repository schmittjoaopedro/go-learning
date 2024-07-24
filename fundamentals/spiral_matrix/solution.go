package spiral_matrix

// https://leetcode.com/problems/spiral-matrix
// Start: 14:00
// End: 14:22

/*
[ 1, 2, 3, 4, 5]
[ 6, 7, 8, 9,10]
[11,12,13,14,15]
[16,17,18,19,20]
[21,22,23,24,25]

[1,2,3,4,5,10,15,20,25,24,23,22,21,16,11,6,7,8,9,14,19,18,17,12,7]
*/
func SpiralOrder(matrix [][]int) []int {
	totalSteps := len(matrix) * len(matrix[0])
	output := make([]int, totalSteps)
	leftMax := 0
	rightMax := len(matrix[0]) - 1
	topMax := 1 // because we start first row
	bottomMax := len(matrix) - 1
	dir := 'r' // r = right, d = down, l = left, u = up
	row := 0
	col := 0

	for i := 0; i < totalSteps; i++ {
		output[i] = matrix[row][col]
		if dir == 'r' {
			if col == rightMax {
				rightMax--
				dir = 'd'
				row++
			} else {
				col++
			}
		} else if dir == 'd' {
			if row == bottomMax {
				bottomMax--
				dir = 'l'
				col--
			} else {
				row++
			}
		} else if dir == 'l' {
			if col == leftMax {
				leftMax++
				dir = 'u'
				row--
			} else {
				col--
			}
		} else {
			if row == topMax {
				topMax++
				dir = 'r'
				col++
			} else {
				row--
			}
		}
	}
	return output

}
