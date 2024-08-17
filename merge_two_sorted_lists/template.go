package merge_two_sorted_lists

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var initNode = &ListNode{Val: -999, Next: nil}
	recursiveMergeNode(initNode, list1, list2)
	return initNode.Next
}

func recursiveMergeNode(node *ListNode, node1 *ListNode, node2 *ListNode) {
	if node1 == nil && node2 == nil {
		return
	}

	if node1 == nil {
		node.Next = node2
		node2 = node2.Next
		node = node.Next
	} else if node2 == nil {
		node.Next = node1
		node1 = node1.Next
		node = node.Next
	} else if node1.Val == node2.Val {
		node.Next = node1
		node1 = node1.Next
		node = node.Next

		node.Next = node2
		node2 = node2.Next
		node = node.Next
	} else if node1.Val > node2.Val {
		node.Next = node2
		node2 = node2.Next
		node = node.Next
	} else {
		node.Next = node1
		node1 = node1.Next
		node = node.Next
	}
	recursiveMergeNode(node, node1, node2)
}

// Go call this func in main.go
func Go() {
	nodeList1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	nodeList2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}
	nodeResult := mergeTwoLists(nodeList1, nodeList2)

	for nodeResult.Next != nil {
		fmt.Println(nodeResult.Next)
		nodeResult = nodeResult.Next
	}
}
