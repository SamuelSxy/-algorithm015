/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	hashMap := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		hashMap[s[i]]++
		hashMap[t[i]]--
	}

	for _, v := range hashMap {
		if v != 0 {
			return false
		}
	}

	return true

}

// @lc code=end

