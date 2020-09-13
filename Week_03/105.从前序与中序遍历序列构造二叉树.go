/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
 */

// @lc code=start
//@todo 迭代
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{preorder[0], nil, nil}

	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] { //确定中序遍历数组中根的index
			break
		}
	}

	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])

	return root

}

// @lc code=end

