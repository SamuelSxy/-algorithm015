/*
 * @lc app=leetcode.cn id=557 lang=golang
 *
 * [557] 反转字符串中的单词 III
 *
 * https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/description/
 *
 * algorithms
 * Easy (73.40%)
 * Likes:    253
 * Dislikes: 0
 * Total Accepted:    100.8K
 * Total Submissions: 137.3K
 * Testcase Example:  `"Let's take LeetCode contest"`
 *
 * 给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
 *
 *
 *
 * 示例：
 *
 * 输入："Let's take LeetCode contest"
 * 输出："s'teL ekat edoCteeL tsetnoc"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。
 *
 *
 */

// @lc code=start
func reverseWords(s string) string {

	buf := []byte(s)

	l, r := 0, 0

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			r = i - 1

			for l < r {
				buf[l], buf[r] = buf[r], buf[l]
				l++
				r--
			}

			l = i + 1
		}

		if i == len(s)-1 {
			r = i

			for l < r {
				buf[l], buf[r] = buf[r], buf[l]
				l++
				r--
			}

			l = i + 1
		}

	}

	return string(buf)
}

// @lc code=end

