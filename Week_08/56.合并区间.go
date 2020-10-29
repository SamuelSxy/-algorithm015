import "sort"

/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {

	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ret := [][]int{intervals[0]}

	for i := 0; i < len(intervals); i++ {

		if intervals[i][0] <= ret[len(ret)-1][1] {
			if ret[len(ret)-1][1] < intervals[i][1] {
				ret[len(ret)-1][1] = intervals[i][1]
			}
		} else {
			ret = append(ret, intervals[i])
		}
	}

	return ret

}

// @lc code=end

