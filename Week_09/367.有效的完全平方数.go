/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 */

// @lc code=start
func isPerfectSquare(num int) bool {

	if num < 2 {
		return true
	}

	l, r := 2, num

	for l <= r {

		mid := (l + r) / 2
		res := mid * mid

		if res == num {
			return true
		}

		if res < num {
			l = mid + 1
		}

		if res > num {
			r = mid - 1
		}

	}

	return false
}

// @lc code=end

