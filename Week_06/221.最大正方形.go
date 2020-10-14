/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] 最大正方形
 */

// @lc code=start
func maximalSquare(matrix [][]byte) int {

	dp := make([][]int, len(matrix))

	ret := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {

			if j == 0 {
				dp[i] = make([]int, len(matrix[0]))
			}

			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				}

				ret = max(ret, dp[i][j])
			}

		}
	}

	return ret * ret

}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end

