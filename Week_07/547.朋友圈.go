/*
 * @lc app=leetcode.cn id=547 lang=golang
 *
 * [547] 朋友圈
 */
//@todo 并查集
// @lc code=start
func findCircleNum(M [][]int) int {

	visted := make(map[int]bool)
	ret := 0

	for i := 0; i < len(M); i++ {
		if !visted[i] {
			ret++
			dfs(visted, M, i)
			// fmt.Println(visted)
		}
	}

	return ret
}

func dfs(visited map[int]bool, M [][]int, idx int) {

	if visited[idx] {
		return
	}

	visited[idx] = true

	for id, isFriend := range M[idx] {
		if isFriend == 1 && !visited[id] {
			dfs(visited, M, id)
		}
	}
}

// @lc code=end

