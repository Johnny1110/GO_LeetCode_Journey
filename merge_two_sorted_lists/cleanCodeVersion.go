package merge_two_sorted_lists

func mergeTwoListsV2(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	recursiveMergeNodeV2(current, list1, list2)
	return dummy.Next
}

func recursiveMergeNodeV2(current *ListNode, node1 *ListNode, node2 *ListNode) {
	if node1 == nil {
		current.Next = node2
		return
	}
	if node2 == nil {
		current.Next = node1
		return
	}

	if node1.Val <= node2.Val {
		current.Next = node1
		recursiveMergeNodeV2(current.Next, node1.Next, node2)
	} else {
		current.Next = node2
		recursiveMergeNodeV2(current.Next, node1, node2.Next)
	}
}
