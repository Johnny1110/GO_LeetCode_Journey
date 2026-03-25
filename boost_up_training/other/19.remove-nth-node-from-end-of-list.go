/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{}
	dummy.Next = head

	// main process
	pointerA, pointerB := dummy, dummy
	for range n {
		pointerB = pointerB.Next
	}

	for pointerB.Next != nil {
		pointerA = pointerA.Next
		pointerB = pointerB.Next
	}

	deleteNode := pointerA.Next
	pointerA.Next = deleteNode.Next

	return dummy.Next
}

// @lc code=end

