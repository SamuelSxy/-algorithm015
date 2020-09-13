/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
 */

// @lc code=start
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	hash := map[int]*TreeNode{}
	isIn := map[int]bool{}

	dfs(root, &hash)

	for p != nil {
		isIn[p.Val] = true
		p = hash[p.Val]
	}

	for q != nil {
		if isIn[q.Val] {
			return q
		}
		q = hash[q.Val]
	}

	return nil

}

func dfs(node *TreeNode, hash *map[int]*TreeNode) {
	if node == nil {
		return
	}

	mymap := *hash

	if node.Left != nil {
		mymap[node.Left.Val] = node
		*hash = mymap
		dfs(node.Left, hash)
	}

	if node.Right != nil {
		mymap[node.Right.Val] = node
		*hash = mymap
		dfs(node.Right, hash)
	}
}

// @lc code=end

func recursion(root, p, q *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	inL := lowestCommonAncestor(root.Left, p, q)
	inR := lowestCommonAncestor(root.Right, p, q)

	if inL != nil && inR != nil {
		return root
	}

	if inL != nil {
		return inL
	}

	if inR != nil {
		return inR
	}

	return nil

}

