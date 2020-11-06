import "strings"

/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] 同构字符串
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {

	for i := 0; i < len(s); i++ {
		//思路一致，但是使用strings包，学习了
		if strings.Index(s, string(s[i])) != strings.Index(t, string(t[i])) {
			return false
		}
	}
	return true

}

// @lc code=end

func isIsomorphic(s string, t string) bool {

	dictS, dictT := make(map[byte]int), make(map[byte]int) //记录字母首次出现的位置

	for i := 0; i < len(s); i++ {

		if _, ok := dictS[s[i]]; !ok {
			dictS[s[i]] = i
		}

		if _, ok := dictT[t[i]]; !ok {
			dictT[t[i]] = i
		}

		if dictS[s[i]] != dictT[t[i]] { //首次出现的位置不同，则不同结构
			return false
		}

	}

	return true

}

