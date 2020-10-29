/*
 * @lc app=leetcode.cn id=52 lang=golang
 *
 * [52] N皇后 II
 */

// @lc code=start
var ret int

func totalNQueens(n int) int {
	ret = 0

	queens := make([]int, n) //每行皇后所在的位置

	for i := 0; i < n; i++ {
		queens[i] = -1
	}

	NS, NW, SE := map[int]bool{}, map[int]bool{}, map[int]bool{}

	backtrack(queens, n, 0, NS, NW, SE)

	return ret

}

func backtrack(queens []int, n int, row int, NS, NW, SE map[int]bool) {

	if row == n {
		ret++
		return
	}

	for i := 0; i < n; i++ {

		//当前行肯定没有

		if NS[i] { //纵向（上北下南）有皇后
			continue
		}

		if NW[row-i] { //西北斜线有皇后
			continue
		}

		if SE[row+i] { //东南斜线有皇后
			continue
		}

		queens[row] = i
		NS[i] = true
		NW[row-i] = true
		SE[row+i] = true

		backtrack(queens, n, row+1, NS, NW, SE)

		queens[row] = -1

		delete(NS, i)
		delete(NW, row-i)
		delete(SE, row+i)

	}

}

// @lc code=end

