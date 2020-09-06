/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
func preorderTraversal(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}

	res := []int{root.Val}
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)

	return res

}

// @lc code=end

func stack(root *TreeNode) []int {

	res := []int{}

	if root == nil {
		return res
	}

	stack := []*TreeNode{}
	var node *TreeNode

	stack = append(stack, root) //根节点入栈

	for len(stack) > 0 {

		node = stack[len(stack)-1]
		res = append(res, node.Val)

		stack = stack[:len(stack)-1] //出栈

		if node.Right != nil { //右节点入栈
			stack = append(stack, node.Right)
		}

		if node.Left != nil { //左节点入栈
			stack = append(stack, node.Left)
		}

	}

	return res

}



