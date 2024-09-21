package reverse_nodes_in_k_Group

import "fmt"

// ListNode Go Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{}
	recursion(dummy, head, k)
	return dummy.Next
}

// assume there are 3 nodes: A, B, C.
func recursion(lastGroupLastOne, currentGroupFirstOne *ListNode, k int) {
	// 1. find Node C:
	nodeC := currentGroupFirstOne
	for i := range k {
		if nodeC == nil {
			lastGroupLastOne.Next = currentGroupFirstOne
			return
		}
		if i < k-1 {
			nodeC = nodeC.Next
		}
	}

	// 2. keep C's next as D
	nodeD := nodeC.Next

	// 3. keep A for next recursion call `lastGroupLastOne`
	nodeA := currentGroupFirstOne

	// 4. reverse K Nodes "A->B->C" => "A<-B<-C"
	reverseKNode(currentGroupFirstOne, currentGroupFirstOne.Next, k-1)

	// 5. link last group last one to nodeC
	lastGroupLastOne.Next = nodeC

	// 6. recursion call
	recursion(nodeA, nodeD, k)

}

func reverseKNode(first, second *ListNode, k int) {
	if k == 0 {
		return
	}
	third := second.Next
	second.Next = first
	reverseKNode(second, third, k-1)

}

// Go call this func in main.go
func Go() {
	testNode1 := &ListNode{1, nil}
	//testNode1 := &ListNode{1, &ListNode{2, nil}}
	//testNode1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	//testNode1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	//testNode1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}

	ans1 := reverseKGroup(testNode1, 1)

	//fmt.Println(ans1.Val)
	//fmt.Println(ans1.Next.Val)
	//fmt.Println(ans1.Next.Next.Val)
	//fmt.Println(ans1.Next.Next.Next.Val)
	//fmt.Println(ans1.Next.Next.Next.Next.Val)
	//fmt.Println(ans1.Next.Next.Next.Next.Next.Val)

	for ans1 != nil {
		fmt.Println(ans1.Val)
		ans1 = ans1.Next
	}
}
