/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */

// @lc code=start
func longestValidParentheses(s string) int {

	dp := make([]int, len(s))
	max := 0

	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else {
				//s[i]对应的'('坐标为(i-dp[i]+1)
				//s[i-1]为 i-1-dp[i-1]+1 => i-dp[i-1]
				//则此时需要判断的为 i-dp[i-1]-1
				if i-dp[i-1] >= 1 && s[i-dp[i-1]-1] == '(' {
					//dp[i-1]+2为当前长度
					//dp[i-dp[i-1]-2]为之前一段
					if i-dp[i-1] >= 2 {
						dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]
					} else {
						dp[i] = dp[i-1] + 2
					}

				}
			}

			if dp[i] > max {
				max = dp[i]
			}
		}
	}

	return max

}

// @lc code=end