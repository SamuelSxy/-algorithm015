/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */

// @lc code=start
func numDecodings(s string) int {

	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	pre, curr := 1, 1

	for i := 1; i < len(s); i++ {
		temp := curr
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' { //10,20
				curr = pre
				continue
			} else { //00
				return 0
			}
		} else {

			if s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6') { //10-26
				curr = curr + pre
			}
		}

		pre = temp
	}

	return curr

}

// @lc code=end

func numDecodings(s string) int {

	if s[0] == '0' || len(s) == 0 {
		return 0
	}

	dp := make([]int, len(s))

	dp[0] = 1

	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				if i > 1 {
					dp[i] = dp[i-2]
				} else {
					dp[i] = 1
				}

			} else {
				return 0
			}
			continue
		}

		if s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6' && s[i] >= '1') {
			if i > 1 {
				dp[i] = dp[i-1] + dp[i-2]
			} else {
				dp[i] = dp[i-1] + 1
			}
			continue
		}

		dp[i] = dp[i-1]
	}

	return dp[len(s)-1]

}

