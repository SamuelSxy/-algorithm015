import "sort"

/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {

	res := [][]int{}
	used := make([]bool, len(nums))
	sort.Ints(nums) //排序排序排序！
	recursion([]int{}, used, &res, nums)

	return res

}

func recursion(ans []int, used []bool, res *[][]int, nums []int) {

	if len(nums) == len(ans) {
		// 切片底层公用数据，所以要copy[46竟然通过了...]
		ret := make([]int, len(ans))
		copy(ret, ans)
		*res = append(*res, ret)
		return
	}

	for i := 0; i < len(nums); i++ {

		if used[i] {
			continue
		}

		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}

		newAns := append(ans, nums[i])

		newUsed := make([]bool, len(used))
		copy(newUsed, used)
		newUsed[i] = true

		recursion(newAns, newUsed, res, nums)
	}

	return

}

// @lc code=end

//[3,3,0,3]