/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {

	l, r := 0, len(height)-1

	maxL, maxR := 0, 0

	res := 0

	for l < r {
		if height[l] < height[r] {
			if height[l] > maxL {
				maxL = height[l]
			} else {
				res += maxL - height[l]
			}

			l++
		} else {
			if height[r] > maxR {
				maxR = height[r]
			} else {
				res += maxR - height[r]
			}

			r--
		}

	}

	return res

}

// @lc code=end

