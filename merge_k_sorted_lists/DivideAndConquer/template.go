package merge_k_sorted_lists_dac

import (
	"fmt"
)

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	return mergeLists(lists, 0, len(lists)-1)
}

func mergeLists(lists []*ListNode, left int, right int) *ListNode {
	if left == right {
		return lists[left]
	}

	mid := (left + right) / 2

	// divide problem
	l1 := mergeLists(lists, left, mid)
	l2 := mergeLists(lists, mid+1, right)

	return mergeTwoLists(l1, l2)
}

// Go problem-021: https://leetcode.com/problems/merge-two-sorted-lists/description/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	currentNode := dummy
	recursiveMergeTwoLists(currentNode, l1, l2)
	return dummy.Next
}

func recursiveMergeTwoLists(node *ListNode, l1 *ListNode, l2 *ListNode) {
	if l1 == nil {
		node.Next = l2
		return
	}
	if l2 == nil {
		node.Next = l1
		return
	}

	if l1.Val < l2.Val {
		node.Next = l1
		l1 = l1.Next
	} else {
		node.Next = l2
		l2 = l2.Next
	}
	recursiveMergeTwoLists(node.Next, l1, l2)
}

// for test data:
// Helper function to create a linked list from a slice
func createList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	current := head
	for _, val := range arr[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

// Helper function to print the linked list
func printList(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val)
		if node.Next != nil {
			fmt.Print("->")
		}
		node = node.Next
	}
	fmt.Println()
}

// Go call this func in main.go
func Go() {
	// Example 1
	lists := []*ListNode{
		createList([]int{1, 4, 5}),
		createList([]int{1, 3, 4}),
		createList([]int{2, 6}),
	}
	mergedList := mergeKLists(lists)
	printList(mergedList)

	// Example 2
	lists = []*ListNode{}
	mergedList = mergeKLists(lists)
	printList(mergedList)

	// Example 3
	lists = []*ListNode{createList([]int{})}
	mergedList = mergeKLists(lists)
	printList(mergedList)
}
