/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
//先二分确定所在行,再二分确定所在列 logm+logn=log(mn)
func searchMatrix(matrix [][]int, target int) bool {

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	//二分确定所在行
	low, high, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1

	var midR, midC int

	for low <= high {
		midR = (low + high) / 2

		if midR == 0 && target <= matrix[midR][right] {
			break
		}

		if midR > 0 && matrix[midR-1][right] < target && target <= matrix[midR][right] {
			break
		}

		if matrix[midR][right] <= target { //下半区
			low = midR + 1
		} else { //上半区，且不含当前行
			high = midR - 1
		}
	}

	//二分确定所在列
	for left <= right {

		midC = (left + right) / 2

		if matrix[midR][midC] == target {
			return true
		}

		if matrix[midR][midC] <= target {
			left = midC + 1
		} else {
			right = midC - 1
		}

	}

	return false

}

// @lc code=end

//@todo 官方题解

