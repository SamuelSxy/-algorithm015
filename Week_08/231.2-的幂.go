/*
 * @lc app=leetcode.cn id=231 lang=golang
 *
 * [231] 2的幂
 */

// @lc code=start
func isPowerOfTwo(n int) bool {

	ret := false

	if n != 0 && n&(n-1) == 0 {
		ret = true
	}

	return ret

}

// @lc code=end

