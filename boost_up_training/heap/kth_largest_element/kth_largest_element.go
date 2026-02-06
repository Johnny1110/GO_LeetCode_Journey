package kth_largest_element

import (
	"container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	// min-heap: smaller value = higher priority
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]     // grab last element
	*h = old[0 : n-1] // shrink slice
	return x
}

func findKthLargest(nums []int, k int) int {
	H := IntHeap(make([]int, 0))

	for _, num := range nums {
		heap.Push(&H, num)
		if H.Len() > k {
			heap.Pop(&H)
		}
	}

	// Kth Largest element will sit on the top.
	return H[0]
}
