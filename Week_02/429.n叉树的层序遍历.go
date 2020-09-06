/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N叉树的层序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {

	res := [][]int{}

	if root == nil {
		return res
	}

	dfs(root, 0, &res)

	return res

}

func dfs(node *Node, level int, res *[][]int) {

	if node == nil {
		return
	}

	if len(*res) == level {
		*res = append(*res, []int{})
	}

	result := *res
	result[level] = append(result[level], node.Val)

	for _, child := range node.Children {
		dfs(child, level+1, res)
	}
}

// @lc code=end

func bfs(root *Node) [][]int {

	res := [][]int{}

	if root == nil {
		return res
	}

	currList := []*Node{}
	currList = append(currList, root)

	for len(currList) > 0 {
		nextList := []*Node{}
		level := []int{}

		for _, node := range currList {
			level = append(level, node.Val)
			nextList = append(nextList, node.Children...)
		}

		res = append(res, level)

		currList = nextList

	}

	return res

}

func queue(root *Node) [][]int {

	res := [][]int{}

	if root == nil {
		return res
	}

	queue := []*Node{}
	queue = append(queue, root)

	for len(queue) > 0 {
		level := []int{}
		size := len(queue)
		for size > 0 {
			node := queue[0]
			queue = queue[1:]

			size--

			level = append(level, node.Val)
			queue = append(queue, node.Children...)
		}

		res = append(res, level)

	}

	return res

}

