/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	interval := 1

	for interval < len(lists) {

		for i := 0; i+interval < len(lists); i += interval * 2 {
			lists[i] = merge2Lists(lists[i], lists[i+interval])
		}

		interval *= 2 // Double the interval for the next round
	}

	return lists[0]
}

func merge2Lists(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	pointer := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			pointer.Next = list1
			pointer = list1
			list1 = list1.Next
		} else {
			pointer.Next = list2
			pointer = list2
			list2 = list2.Next
		}
	}

	if list1 != nil {
		pointer.Next = list1
	}

	if list2 != nil {
		pointer.Next = list2
	}

	return dummy.Next
}

// @lc code=end

