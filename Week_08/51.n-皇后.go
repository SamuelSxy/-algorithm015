/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51]=true

 baN 皇后
*/

// @lc code=start
var ret [][]string

func solveNQueens(n int) [][]string {

	ret = [][]string{}

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
		ret = append(ret, boardFormat(queens, n))
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

func boardFormat(queens []int, n int) []string {
	board := []string{}

	for i := 0; i < n; i++ {

		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}

		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board

}

// @lc code=end

