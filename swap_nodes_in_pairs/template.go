package swap_nodes_in_pairs

import (
	"fmt"
)

// ListNode Go Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	secNode := head.Next
	swap(head)
	return secNode
}

func swap(firstNode *ListNode) {
	secNode := firstNode.Next
	if secNode == nil {
		return
	}

	thirdNode := secNode.Next

	secNode.Next = firstNode

	if thirdNode == nil {
		firstNode.Next = nil
		return
	}

	forthNode := thirdNode.Next

	if forthNode == nil {
		firstNode.Next = thirdNode
		return
	}

	firstNode.Next = forthNode
	swap(thirdNode)
}

// Go call this func in main.go
func Go() {
	testNode1 := &ListNode{1, nil}
	//testNode1 := &ListNode{1, &ListNode{2, nil}}
	//testNode1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	//testNode1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	//testNode1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}
	ans1 := swapPairs(testNode1)

	for ans1 != nil {
		fmt.Println(ans1.Val)
		ans1 = ans1.Next
	}
}
