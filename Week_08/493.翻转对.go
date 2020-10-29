/*
 * @lc app=leetcode.cn id=493 lang=golang
 *
 * [493] 翻转对
 */

//@todo 树状数组

// @lc code=start

func reversePairs(nums []int) int {
	//归并
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, l, r int) int {

	if l >= r {
		return 0
	}

	m := (l + r) >> 1

	count := mergeSort(nums, l, m) + mergeSort(nums, m+1, r)
	cache := make([]int, r-l+1)

	i, k := l, 0

	for j, idx := m+1, l; j <= r; j++ {
		for ; i <= m && nums[i] < nums[j]; i++ {
			// 合并左侧数组
			cache[k] = nums[i]
			k++
		}

		// 合并右侧数组
		cache[k] = nums[j]
		k++

		//查询i < j 且 nums[i] > 2*nums[j]
		for idx <= m && (nums[idx]+1)>>1 <= nums[j] {
			idx++
		}

		count += m - idx + 1
	}

	copy(nums[l+k:], nums[i:m+1])
	copy(nums[l:], cache[:k])

	return count

}

// @lc code=end

