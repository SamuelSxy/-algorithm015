/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}

	res := append(inorderTraversal(root.Left), root.Val)
	res = append(res, inorderTraversal(root.Right)...)

	return res

}

// @lc code=end

func stack(root *TreeNode) []int {

	stack := []*TreeNode{}
	res := []int{}

	var node *TreeNode

	for {

		if root != nil {
			stack = append(stack, root) //入栈
			root = root.Left
			continue
		}

		if len(stack) > 0 {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1] //出栈
			res = append(res, node.Val)
			root = node.Right
			continue
		}

		break

	}

	return res

}

