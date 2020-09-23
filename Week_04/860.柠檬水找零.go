/*
 * @lc app=leetcode.cn id=860 lang=golang
 *
 * [860] 柠檬水找零
 */

// @lc code=start
func lemonadeChange(bills []int) bool {

	cash := make(map[int]int)

	for _, money := range bills {

		cash[money]++

		if money == 20 && cash[10] > 0 {
			cash[10]--
			money = 10
		}

		if money >= 5 && cash[5] >= (money/5-1) {
			cash[5] -= (money/5 - 1)
			money = 5
		}

		if money == 5 {
			continue
		}

		return false

	}

	return true

}

// @lc code=end

