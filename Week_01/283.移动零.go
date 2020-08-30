/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {

	j, temp := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}

		temp = nums[j]
		nums[j] = nums[i]
		nums[i] = temp

		j++

	}

}

// @lc code=end

