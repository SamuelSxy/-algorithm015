/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
func findMin(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	l, r := 0, len(nums)-1

	if nums[l] < nums[r] { //完全顺序
		return nums[l]
	}

	for l < r { //左右夹逼至l==r
		// mid := (l + r) / 2
		mid := l + (r-l)/2 //防止越界

		if nums[0] <= nums[mid] { //左侧顺序，最小值在右侧
			l = mid + 1 //l肯定不是最小值，+1
		} else {
			r = mid //r可能为最小值
		}

	}

	return nums[l]

}

// @lc code=end

