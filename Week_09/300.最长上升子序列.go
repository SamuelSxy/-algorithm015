/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长上升子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {

	res := 0
	tails := make([]int, len(nums))

	for _, num := range nums {

		l, r := 0, res

		for l < r {
			m := (l + r) >> 1

			if tails[m] < num {
				l = m + 1
			} else {
				r = m
			}
		}

		tails[l] = num

		if res == r {
			res++
		}

	}

	return res

}

// @lc code=end

func dp(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))

	dp[0] = 1
	max := 1
	for i := 1; i < len(nums); i++ {

		dp[i] = 1

		for j := 0; j < i; j++ {

			if nums[i] > nums[j] {
				dp[i] = getMax(dp[i], dp[j]+1)
			}

		}

		max = getMax(max, dp[i])

	}

	return max
}

func getMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}

