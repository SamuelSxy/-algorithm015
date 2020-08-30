/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {

	j := 0

	for i := 1; i < len(nums); i++ {

		if nums[i] == nums[i-1] {
			continue
		}

		j++
		nums[j] = nums[i]

	}

	return j + 1

}

// @lc code=end

