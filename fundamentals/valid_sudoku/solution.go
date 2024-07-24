package valid_sudoku

// https://leetcode.com/problems/valid-sudoku
// Start: 13:11
// End: 13:44

func IsValidSudoku(board [][]byte) bool {
	nRows := map[byte]map[int]bool{}
	nCols := map[byte]map[int]bool{}
	nGrid := map[byte]map[int]bool{}
	skipMark := byte('.')

	for i, row := range board {
		for j, cell := range row {
			if cell != byte(skipMark) {
				g := 10*int(i/3) + int(j/3)
				dupRow := isDuplicated(nRows, cell, i)
				dupCol := isDuplicated(nCols, cell, j)
				dupGrd := isDuplicated(nGrid, cell, g)
				if dupRow || dupCol || dupGrd {
					return false
				}
			}
		}
	}

	return true
}

func isDuplicated(dict map[byte]map[int]bool, key byte, sub_key int) bool {
	_, ok := dict[key]
	if !ok {
		dict[key] = map[int]bool{}
		dict[key][sub_key] = true
		return false
	} else {
		_, subOk := dict[key][sub_key]
		if !subOk {
			dict[key][sub_key] = true
			return false
		} else {
			return true
		}
	}
}
