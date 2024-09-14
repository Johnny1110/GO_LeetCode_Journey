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
	if head.Next == nil {
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
		return
	}

	forthNode := thirdNode.Next

	if forthNode == nil {
		return
	}

	firstNode.Next = forthNode
	fmt.Println("first: ", firstNode.Val)
	fmt.Println("sec: ", secNode.Val)
	fmt.Println("third: ", thirdNode.Val)
	fmt.Println("forth: ", forthNode.Val)
	fmt.Println("first.next: ", firstNode.Next.Val)
	fmt.Println("sec.next: ", secNode.Next.Val)
	fmt.Println("third.next: ", thirdNode.Next.Val)
	fmt.Println("forth.next: ", forthNode.Next.Val)
	swap(thirdNode)
}

// Go call this func in main.go
func Go() {
	testNode1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	ans := swapPairs(testNode1)

	fmt.Println(ans.Val)
	fmt.Println(ans.Next.Val)
	fmt.Println(ans.Next.Next.Val)
	fmt.Println(ans.Next.Next.Next.Val)
	fmt.Println(ans.Next.Next.Next.Next.Val)

}
