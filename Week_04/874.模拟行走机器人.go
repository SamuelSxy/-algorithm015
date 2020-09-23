/*
 * @lc app=leetcode.cn id=874 lang=golang
 *
 * [874] 模拟行走机器人
 */

// @lc code=start

func robotSim(commands []int, obstacles [][]int) int {

	pos := [2]int{0, 0}
	dict := 1
	max := 0

	obMap := make(map[[2]int]bool)

	for _, ob := range obstacles { //hashmap记录障碍物，减少遍历
		obMap[[2]int{ob[0], ob[1]}] = true
	}

	for _, comd := range commands {

		if comd < 0 { //转向
			dict = turn(dict, comd)
		} else { //移动
			pos = walk(dict, comd, pos, obMap)
			temp := pos[0]*pos[0] + pos[1]*pos[1]

			if temp > max {
				max = temp
			}
		}
	}

	return max

}

func turn(dict, comd int) int {

	if comd == -2 {
		dict -= 1
	}

	if comd == -1 {
		dict += 1
	}

	if dict <= 0 {
		dict += 4
	}

	if dict > 4 {
		dict -= 4
	}

	return dict

}

func walk(dict, comd int, pos [2]int, obMap map[[2]int]bool) [2]int {

	target := pos

	for i := 0; i < comd; i++ {
		if dict == 1 { //上北
			target[1]++
		}

		if dict == 2 { //右东
			target[0]++
		}

		if dict == 3 { //下南
			target[1]--
		}

		if dict == 4 { //左西
			target[0]--
		}

		if obMap[target] {
			return pos
		}

		pos = target

	}

	return pos

}

// @lc code=end

