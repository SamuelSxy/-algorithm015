/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {

	if n < 3 {
		return n
	}

	//f(n)=f(n-1)+f(n-2)
	pre2, pre1, curr := 1, 2, 0

	for i := 3; i <= n; i++ {

		curr = pre2 + pre1
		pre2 = pre1
		pre1 = curr

	}

	return curr

}

// @lc code=end

func climbStairs_0901(n int) int {

	//第i-2，i-1，i步
	//如果i=0，[0 0 1] 1 2 3 5...:
	pre2, pre1, curr := 0, 0, 1

	for i := 1; i <= n; i++ {
		pre2 = pre1
		pre1 = curr
		curr = pre2 + pre1
	}

	return curr

}