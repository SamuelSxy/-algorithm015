/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {

	row, col, blk := make([]map[byte]bool, 9), make([]map[byte]bool, 9), make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		row[i] = map[byte]bool{}
		col[i] = map[byte]bool{}
		blk[i] = map[byte]bool{}
	}

	for i := 0; i < 9; i++ {

		for j := 0; j < 9; j++ {

			val := board[i][j]

			if val == '.' {
				continue
			}

			if _, ok := row[i][val]; ok {
				return false
			}

			if _, ok := col[j][val]; ok {
				return false
			}

			if _, ok := blk[i/3*3+j/3][val]; ok {
				return false
			}

			row[i][val] = true
			col[j][val] = true
			blk[i/3*3+j/3][val] = true

		}
	}

	return true

}

// @lc code=end

