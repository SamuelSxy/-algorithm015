/*
 * @lc app=leetcode.cn id=552 lang=golang
 *
 * [552] 学生出勤记录 II
 */

// @lc code=start
func checkRecord(n int) int {

	M := 1000000007

	lenth := n + 1

	if lenth < 6 {
		lenth = 6
	}

	dp := make([]int, lenth)

	dp[0] = 1
	dp[1] = 2
	dp[2] = 4
	dp[3] = 7

	for i := 4; i <= n; i++ {
		dp[i] = ((2*dp[i-1])%M + (M - dp[i-4])) % M
	}

	sum := dp[n]

	for i := 1; i <= n; i++ {
		sum += (dp[i-1] * dp[n-i]) % M
	}
	return (int)(sum % M)

}

// @lc code=end

