
/*
 * @lc app=leetcode.cn id=621 lang=golang
 *
 * [621] 任务调度器
 */

// @lc code=start
func leastInterval(tasks []byte, n int) int {

	counts := [26]int{}

	for _, v := range tasks {
		counts[v-'A']++
	}

	//获取最多的任务
	max := 0
	maxCount := 0
	for _, num := range counts {
		if num > max {
			max = num
			maxCount = 1
			continue
		}

		if num == max {
			maxCount++
		}
	}

	return getMax(len(tasks), (n+1)*(max-1)+maxCount)

}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

