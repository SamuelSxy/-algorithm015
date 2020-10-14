/*
 * @lc app=leetcode.cn id=312 lang=golang
 *
 * [312] 戳气球
 */

// 参考题解：
// https://leetcode-cn.com/problems/burst-balloons/solution/dong-tai-gui-hua-tao-lu-jie-jue-chuo-qi-qiu-wen-ti/

// @lc code=start
func maxCoins(nums []int) int {

	n := len(nums)

	//“你可以假设 nums[-1] = nums[n] = 1，但注意它们不是真实存在的所以并不能被戳破。”
	//左右补1，即：1,nums[0]...nums[n-1],1
	//len(points)=len(nums)+2：points[0]....ponts[n+1]
	points := make([]int, n+2)
	points[0], points[n+1] = 1, 1
	for i := 1; i < n+1; i++ {
		points[i] = nums[i-1]
	}

	//dp[a][b]表示一个开区间(a,b)戳破所有气球的最高分
	//(0,0)->(n+1,n+1)
	dp := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		dp[i] = make([]int, n+2)
	}

	//从对角线，自下而上，自左而右（不包含对角线）
	for l := n + 1; l >= 0; l-- { //自下而上
		for r := l + 1; r <= n+1; r++ { //从对角线，自左而右
			for m := l + 1; m <= r-1; m++ { //最后戳破的气球，不含头尾(l,r)=>[l+1,r-1]
				//nums[left] * nums[mid] * nums[right]
				sum := dp[l][m] + dp[m][r] + points[l]*points[m]*points[r]
				dp[l][r] = max(dp[l][r], sum)
			}
		}
	}

	return dp[0][n+1]

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

