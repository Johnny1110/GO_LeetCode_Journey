package merge_k_sorted_lists

import "container/heap"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	H := Heap([]any{})
	dummyHead := &ListNode{}

	for _, head := range lists {
		if head != nil {
			heap.Push(&H, head)
		}
	}

	tmp := dummyHead
	for H.Len() > 0 {
		node := heap.Pop(&H).(*ListNode)
		tmp.Next = node
		tmp = node

		if node.Next != nil {
			heap.Push(&H, node.Next)
		}
	}

	return dummyHead.Next
}

// ============================================

type Heap []any

func (h Heap) Len() int           { return len(h) }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h Heap) Less(i, j int) bool { return h[i].(*ListNode).Val < h[j].(*ListNode).Val }

func (h *Heap) Push(node any) {
	*h = append(*h, node)
}

func (h *Heap) Pop() any {
	ret := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return ret
}
