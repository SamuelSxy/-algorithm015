/*
 * @lc app=leetcode.cn id=917 lang=golang
 *
 * [917] 仅仅反转字母
 */

// @lc code=start
func reverseOnlyLetters(S string) string {

	buf := []byte(S)

	l, r := 0, len(S)-1

	for l < r {
		if !isLetter(buf[l]) {
			l++
			continue
		}

		if !isLetter(buf[r]) {
			r--
			continue
		}

		buf[l], buf[r] = buf[r], buf[l]
		l++
		r--
	}

	return string(buf)
}

func isLetter(ch byte) bool {

	if (ch <= 'z' && ch >= 'a') || (ch <= 'Z' && ch >= 'A') {
		return true
	}

	return false
}

// @lc code=end

