/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
func preorder(root *Node) []int {

	if root == nil {
		return []int{}
	}

	res := []int{root.Val}

	for _, node := range root.Children {
		res = append(res, preorder(node)...)
	}

	return res

}

// @lc code=end

func stack(root *Node) []int {

	res := []int{}

	if root == nil {
		return res
	}

	var node *Node

	stack := []*Node{}
	stack = append(stack, root)

	for len(stack) > 0 {

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, node.Val)

		if node.Children == nil {
			continue
		}

		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}

	}

	return res

}

