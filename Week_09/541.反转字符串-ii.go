/*
 * @lc app=leetcode.cn id=541 lang=golang
 *
 * [541] 反转字符串 II
 */

// @lc code=start
func reverseStr(s string, k int) string {

	str := []byte(s)

	l := 0

	for l < len(s) {

		r := l + k - 1

		if r >= len(str) {
			r = len(str) - 1
		}

		nextL := l + 2*k

		for r > l {
			str[l], str[r] = str[r], str[l]
			l++
			r--
		}

		l = nextL

	}

	return string(str)

}

// @lc code=end

// "abcdefg"\n2

