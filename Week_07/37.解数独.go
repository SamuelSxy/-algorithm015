/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 */

// @lc code=start
func solveSudoku(board [][]byte) {

	var row, col [9][9]bool //行与列
	var block [3][3][9]bool //块
	var spaces [][2]int     //空格坐标

	for x, line := range board {
		for y, val := range line {

			if val == '.' {
				//记录空格坐标
				spaces = append(spaces, [2]int{x, y})
			} else {

				digit := val - '1' //byte0

				//标记【行列块】中已经使用该数字
				row[x][digit] = true
				col[y][digit] = true
				block[x/3][y/3][digit] = true

			}

		}
	}

	var dfs func(int) bool
	dfs = func(pos int) bool {

		if pos == len(spaces) {
			return true
		}

		x, y := spaces[pos][0], spaces[pos][1] //取空格坐标

		for digit := byte(0); digit < 9; digit++ {

			if !row[x][digit] && !col[y][digit] && !block[x/3][y/3][digit] {

				row[x][digit] = true
				col[y][digit] = true
				block[x/3][y/3][digit] = true

				board[x][y] = digit + '1' //toString

				if dfs(pos+1) == true { //递归无false
					return true
				}

				// 回溯
				row[x][digit] = false
				col[y][digit] = false
				block[x/3][y/3][digit] = false

				//理论需要擦除答案。
				// 但这里用spaces控制空格是否填写，所以可以不用回溯
				// board[i][j] = '.'

			}

		}

		return false

	}

	dfs(0)

}

// @lc code=end

