/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */
//@todo dp
// @lc code=start
func generateParenthesis(n int) []string {

	res := []string{}

	generate(0, 0, n, "", &res)

	return res

}

func generate(l, r, n int, s string, res *[]string) {

	if l == n && r == n {
		*res = append(*res, s)
		return
	}

	if l < n {
		generate(l+1, r, n, s+"(", res)
	}

	if r < l {
		generate(l, r+1, n, s+")", res)
	}

	return

}

// @lc code=end

