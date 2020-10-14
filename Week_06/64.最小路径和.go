/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

// @lc code=start
func minPathSum(grid [][]int) int {

	//上/左
	for i := 0; i < len(grid); i++ { //行遍历
		for j := 0; j < len(grid[0]); j++ { //列遍历

			if i == 0 && j == 0 { //没有上&左
				continue
			}

			if i == 0 { //没有上
				grid[i][j] += grid[i][j-1]
				continue
			}

			if j == 0 { //没有左
				grid[i][j] += grid[i-1][j]
				continue
			}

			//上和左选择最短
			grid[i][j] += getMin(grid[i][j-1], grid[i-1][j])

		}
	}

	return grid[len(grid)-1][len(grid[0])-1]

}

func getMin(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// @lc code=end

