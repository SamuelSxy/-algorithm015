/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start

//hash表
//时间：O(n)
//空间：O(n)
func twoSum(nums []int, target int) []int {

	hash := make(map[int]int)

	for i, num := range nums {
		if j, ok := hash[target-num]; ok {
			return []int{j, i}
		}

		if _, ok := hash[num]; !ok {
			hash[num] = i
		}
	}

	return []int{}

}

// @lc code=end

//时间：O(n^2)
//空间：O(1)
func brute(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}
