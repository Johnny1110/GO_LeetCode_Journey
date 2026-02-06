package top_k_frequent_elements

import (
	"container/heap"
)

type IntMinHeap struct {
	data       []int
	numFreqMap map[int]int
}

func (h *IntMinHeap) Len() int {
	return len(h.data)
}

// this is min-heap
func (h *IntMinHeap) Less(i, j int) bool {
	ip := h.numFreqMap[h.data[i]]
	jp := h.numFreqMap[h.data[j]]
	return ip < jp
}

func (h *IntMinHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *IntMinHeap) Push(x interface{}) {
	h.data = append(h.data, x.(int))
}

func (h *IntMinHeap) Pop() interface{} {
	x := h.data[h.Len()-1]
	h.data = h.data[:h.Len()-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	// key: num val: frequency
	numFreqMap := make(map[int]int)
	for _, num := range nums {
		numFreqMap[num]++
	}

	// create a min int heap
	H := &IntMinHeap{
		data:       make([]int, 0),
		numFreqMap: numFreqMap,
	}

	// push all numFreqMap into
	for num := range numFreqMap {

		heap.Push(H, num)
		if H.Len() > k {
			heap.Pop(H)
		}
	}

	return H.data
}
