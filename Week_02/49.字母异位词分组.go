/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {

	group := make(map[[26]int][]string)

	for _, word := range strs {
		temp := [26]int{}
		for i := 0; i < len(word); i++ {
			temp[int(word[i])-97]++
		}

		group[temp] = append(group[temp], word)

	}

	res := make([][]string, 0, len(group))
	for _, v := range group {
		res = append(res, v)
	}

	return res

}

// @lc code=end

