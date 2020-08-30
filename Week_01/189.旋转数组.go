/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
func rotate(nums []int, k int) {

	if k == 0 {
		return
	}

	k = k % len(nums)

	var target int
	var curr int
	var next int

	times := 0

	for i := 0; times < len(nums); i++ {

		j := i
		curr = nums[i]
		for {
			target = (j + k) % len(nums)

			next = nums[target]

			nums[target] = curr

			j = target
			curr = next

			times++

			if j == i { //封环
				break
			}

		}

	}

}

// @lc code=end

//三次翻转法
func rotateReverse(nums []int, k int) {

	t := k % len(nums)

	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, t-1)
	reverse(nums, t, len(nums)-1)
}

func reverse(nums []int, start int, end int) {

	var temp int
	for start < end {

		temp = nums[start]
		nums[start] = nums[end]
		nums[end] = temp

		start++
		end--

	}
}

