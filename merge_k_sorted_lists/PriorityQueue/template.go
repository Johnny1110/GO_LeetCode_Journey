package merge_k_sorted_lists_pq

import (
	"container/heap"
	"fmt"
)

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// PriorityQueue implements a min-heap for ListNode pointers
type PriorityQueue []*ListNode

// Len returns the length of the priority queue
func (pq PriorityQueue) Len() int { return len(pq) }

// Less is the comparator for ListNode, it prioritizes smaller values
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Val < pq[j].Val }

// Swap swaps two elements in the priority queue
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

// Push pushes an element into the priority queue
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*ListNode))
}

// Pop removes and returns the smallest element from the priority queue
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {
	// init a pq
	pq := &PriorityQueue{}
	heap.Init(pq)

	// put all linked list head in to heap
	for _, list := range lists {
		if list != nil {
			heap.Push(pq, list)
		}
	}

	// init a dummy node to build result
	dummy := &ListNode{}
	currentNode := dummy

	for pq.Len() > 0 {
		smallestNode := heap.Pop(pq).(*ListNode)
		currentNode.Next = smallestNode
		currentNode = smallestNode

		if smallestNode.Next != nil {
			// because all given nodes are sorted, so if we added a node to dummy's next, we can put node.next to pq
			heap.Push(pq, smallestNode.Next)
		}
	}
	return dummy.Next
}

// Helper function to create a linked list from a slice of integers
func createLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, num := range nums[1:] {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return head
}

// Helper function to print the linked list
func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}

// Go call this func in main.go
func Go() {
	// Create multiple sorted linked lists
	list1 := createLinkedList([]int{1, 4, 5})
	list2 := createLinkedList([]int{1, 3, 4})
	list3 := createLinkedList([]int{2, 6})

	// Merge all k sorted lists
	lists := []*ListNode{list1, list2, list3}
	mergedList := mergeKLists(lists)

	// Print the merged list
	printLinkedList(mergedList)
}
