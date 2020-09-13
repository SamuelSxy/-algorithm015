/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	res := [][]int{}

	recursion([]int{}, nums, &res)

	return res

}

func recursion(ans []int, nums []int, res *[][]int) {

	if len(nums) == 0 {
		//@fixbug：切片底层数据共用【详见47题，这里可以通过】
		ret := make([]int, len(ans))
		copy(ret, ans)
		*res = append(*res, ret)
		return
	}

	for k, v := range nums {
		newAns := append(ans, v)

		newNums := append([]int{}, nums[:k]...)
		newNums = append(newNums, nums[k+1:]...)

		recursion(newAns, newNums, res)
	}

	return

}

// @lc code=end

