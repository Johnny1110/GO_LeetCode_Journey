/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type MinHeap struct {
	data []*ListNode
}

func (h MinHeap) Len() int      { return len(h.data) }
func (h MinHeap) Swap(i, j int) { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h MinHeap) Less(i, j int) bool {
	return h.data[i].Val < h.data[j].Val
}

func (h *MinHeap) Push(node interface{}) {
	h.data = append(h.data, node.(*ListNode))
}

func (h *MinHeap) Pop() interface{} {
	ret := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	return ret
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	dummy := &ListNode{}
	pointer := dummy

	H := &MinHeap{
		data: []*ListNode{},
	}

	heap.Push(H, list1)
	heap.Push(H, list2)

	for H.Len() > 0 {
		node := heap.Pop(H).(*ListNode)
		tmp := node.Next
		node.Next = nil

		pointer.Next = node
		pointer = node

		if tmp != nil {
			heap.Push(H, tmp)
		}
	}

	return dummy.Next
}

// @lc code=end

