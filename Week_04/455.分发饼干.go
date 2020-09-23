import "sort"

/*
 * @lc app=leetcode.cn id=455 lang=golang
 *
 * [455] 分发饼干
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {

	ans := 0

	sort.Ints(g)
	sort.Ints(s)

	for {
		if len(s) == 0 || len(g) == 0 {
			break
		}

		if s[0] >= g[0] {
			ans++
			g = g[1:]
		}
		s = s[1:]

	}

	return ans

}

// @lc code=end

