/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {

	ret := 0

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {

			if grid[row][col] == '0' { //!!!byte
				continue
			}

			ret++

			dfs(grid, row, col)
		}
	}

	return ret

}

func dfs(grid [][]byte, row, col int) {

	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[row]) || grid[row][col] == '0' {
		return
	}

	if grid[row][col] == '1' {
		grid[row][col] = '0'
	}

	dfs(grid, row-1, col)
	dfs(grid, row+1, col)
	dfs(grid, row, col-1)
	dfs(grid, row, col+1)

}

// @lc code=end

