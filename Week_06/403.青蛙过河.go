/*
 * @lc app=leetcode.cn id=403 lang=golang
 *
 * [403] 青蛙过河
 */
// @todo时间优化
// @lc code=start
func canCross(stones []int) bool {

	dp := make(map[int]map[int]int) //记录下一跳的距离(k),同一个石头可能有多个hop（来源不同）

	for _, idx := range stones {
		dp[idx] = make(map[int]int, 0)
	}

	dp[0][0] = 0

	for i := 0; i < len(stones); i++ {
		for _, hop := range dp[stones[i]] {
			for nextHop := hop - 1; nextHop <= hop+1; nextHop++ {
				if nextHop <= 0 {
					continue
				}

				if _, ok := dp[stones[i]+nextHop]; !ok { //没有这块石头
					continue
				}

				// dp[stones[i]+nextHop] = append(dp[stones[i]+nextHop], nextHop)
				//考虑去重，修改为map
				dp[stones[i]+nextHop][nextHop] = nextHop

			}
		}
	}

	return len(dp[stones[len(stones)-1]]) > 0

}

// @lc code=end

