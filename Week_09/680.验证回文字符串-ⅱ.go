/*
 * @lc app=leetcode.cn id=680 lang=golang
 *
 * [680] 验证回文字符串 Ⅱ
 */

// @lc code=start
func validPalindrome(s string) bool {
	return test(s, false)
}

func test(s string, isDeleted bool) bool {

	l, r := 0, len(s)-1

	for l < r {
		if s[l] == s[r] { //符合回文
			l++
			r--
		} else { //不符合，删除l或者r

			if isDeleted {
				return false
			}

			return test(s[l+1:r+1], true) || test(s[l:r], true)
		}
	}

	return true

}

// @lc code=end

