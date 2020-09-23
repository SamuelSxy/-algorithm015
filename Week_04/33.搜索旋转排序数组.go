/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {

	l, r := 0, len(nums)-1

	//二分法，边缘比较
	for l <= r {
		mid := (l + r) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[l] <= nums[mid] { //左侧全部升序		//⚠️此处必须(l,mid]
			if nums[l] <= target && target < nums[mid] { //目标在顺序区域
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else { //右侧全部升序，且目标在顺序区域
			if nums[mid] < target && target <= nums[r] { //目标在顺序区域
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -1

}

// @lc code=end

