/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {

	dp := make([][]int, len(word1)+1)

	for i := 0; i < len(word1)+1; i++ {
		for j := 0; j < len(word2)+1; j++ {
			if j == 0 {
				dp[i] = make([]int, len(word2)+1)
				dp[i][j] = i
				continue
			}

			if i == 0 {
				dp[i][j] = j
				continue
			}

			N := dp[i-1][j] + 1
			W := dp[i][j-1] + 1
			NW := dp[i-1][j-1] + 1

			if word1[i-1] == word2[j-1] {
				NW = NW - 1
			}

			dp[i][j] = min(min(N, W), NW)

		}
	}

	return dp[len(word1)][len(word2)]

}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// @lc code=end

