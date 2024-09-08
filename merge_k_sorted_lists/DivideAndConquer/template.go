package merge_k_sorted_lists

import "go_leetcode_journey/common"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	return nil
}

// Go call this func in main.go
func Go() {
	common.Assert_answer(mergeKLists(nil), nil)
}
