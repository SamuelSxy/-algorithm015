/*
 * @lc app=leetcode.cn id=647 lang=golang
 *
 * [647] 回文子串
 */

// @lc code=start
func countSubstrings(s string) int {

	t := "$#"

	//马拉车
	for i := 0; i < len(s); i++ {
		t += string(s[i]) + "#"
	}

	t += "!"

	mem := make([]int, len(t))
	iMax, rMax := 0, 0
	ans := 0

	for i := 1; i < len(t)-1; i++ {
		if i <= rMax {
			mem[i] = min(rMax-i+1, mem[2*iMax-i])
		} else {
			mem[i] = 1
		}

		//中心拓展
		for t[i+mem[i]] == t[i-mem[i]] {
			mem[i]++
		}

		if i+mem[i]-1 > rMax {
			iMax = i
			rMax = i + mem[i] - 1
		}

		ans += mem[i] >> 1
	}

	return ans

}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// @lc code=end

