/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */
//@todo 非递归（字典序法）
// @lc code=start
func combine(n int, k int) [][]int {

	ans := [][]int{}

	recursion(n, k, 1, []int{}, &ans)

	return ans
}

func recursion(n, k, curr int, res []int, ans *[][]int) {

	if len(res)+(n-curr+1) < k { //当前元素后续的所有节点数小于k
		return
	}

	if len(res) == k {
		ret := make([]int, k)
		copy(ret, res)
		*ans = append(*ans, ret)
		return
	}

	newRes := append(res, curr)

	curr++

	recursion(n, k, curr, newRes, ans) //将当前节点加入答案[1,2]=>[1,2,3]
	recursion(n, k, curr, res, ans)    //不将当前节点加入答案[1]=>[1,3]
}

// @lc code=end

