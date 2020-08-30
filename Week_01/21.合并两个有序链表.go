/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	var head, master, slave, temp *ListNode

	master = &ListNode{}
	master.Val = ^int(^uint(0) >> 1)
	master.Next = l1

	slave = l2

	head = master

	for slave != nil {
		if master.Next == nil {
			master.Next = slave
			return head.Next
		}

		if master.Next.Val > slave.Val {
			temp = master.Next
			master.Next = slave
			slave = temp
		}

		master = master.Next

	}

	return head.Next

}

// @lc code=end

