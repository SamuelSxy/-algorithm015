/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {

	i := 0
	for {

		if i >= len(nums)-1 || i+nums[i] >= len(nums)-1 { //本次跳跃或者下次跳跃的超过终点
			return true
		}

		if nums[i] == 0 {
			return false
		}

		max := 0
		for j := 1; j <= nums[i]; j++ {
			next := i + j
			if next+nums[next] >= max+nums[max] { //选取下一跳最远的index作为本次跳跃目标
				max = next
			}
		}

		i = max

	}

	return false

}

// @lc code=end

