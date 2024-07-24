package rotate_image

// https://leetcode.com/problems/rotate-image
// S: 14:38
// E: 15:31

func Rotate(matrix [][]int) {
	nLayers := int(len(matrix) / 2)
	for i := 0; i < nLayers; i++ {
		trow := i                     // top row
		lcol := i                     // left column
		brow := len(matrix) - 1 - i   // bottom row
		rcol := len(matrix) - 1 - i   // right column
		swps := len(matrix) - 1 - 2*i // swaps
		ins := 0                      // insert
		rem := 0                      // remove
		for j := 0; j < swps; j++ {
			ins = matrix[trow][lcol+j]
			rem = matrix[trow+j][rcol]
			matrix[trow+j][rcol] = ins
			ins = rem
			rem = matrix[brow][rcol-j]
			matrix[brow][rcol-j] = ins
			ins = rem
			rem = matrix[brow-j][lcol]
			matrix[brow-j][lcol] = ins
			ins = rem
			matrix[trow][lcol+j] = ins
		}
	}
}
