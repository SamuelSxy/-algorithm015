/*
 * @lc app=leetcode.cn id=529 lang=golang
 *
 * [529] 扫雷游戏
 */

// @lc code=start
var rePos = [8][2]int{{-1, 1}, {0, 1}, {1, 1}, {-1, 0}, {1, 0}, {-1, -1}, {0, -1}, {1, -1}}

func updateBoard(board [][]byte, click []int) [][]byte {

	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
	} else {
		dfs(board, click)
	}

	return board

}

func dfs(board [][]byte, pos []int) {

	if board[pos[0]][pos[1]] != 'E' { //已经被标记
		return
	}

	count := 0

	//搜索周围8个坐标有几颗雷，并且标记
	for _, p := range rePos {
		target := []int{pos[0] + p[0], pos[1] + p[1]}

		if target[0] < 0 || target[1] < 0 || target[0] >= len(board) || target[1] >= len(board[0]) {
			continue
		}

		if board[target[0]][target[1]] == 'M' || board[target[0]][target[1]] == 'X' {
			count++
		}
	}

	if count > 0 { //标记数字后，停止递归（示例中的E）
		board[pos[0]][pos[1]] = byte(count + '0')
		return
	}

	board[pos[0]][pos[1]] = 'B'

	//如果不是数字，则继续打开周围的格子
	for _, p := range rePos {
		target := []int{pos[0] + p[0], pos[1] + p[1]}

		if target[0] < 0 || target[1] < 0 || target[0] >= len(board) || target[1] >= len(board[0]) {
			continue
		}

		dfs(board, target)

	}

}

// @lc code=end

