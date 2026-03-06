package top_k_frequent_elements

import "container/heap"

func topKFrequent(nums []int, k int) []int {

	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	H := &MinHeap{
		data:   []interface{}{},
		weight: freqMap,
	}

	for num, _ := range freqMap {

		heap.Push(H, num)
		if H.Len() > k {
			H.Pop()
		}
	}

	result := make([]int, k)
	for range k {
		result = append(result, heap.Pop(H).(int))
	}

	return result
}

type MinHeap struct {
	data   []interface{}
	weight map[int]int
}

func (h MinHeap) Len() int {
	return len(h.data)
}

// i priority > j
func (h MinHeap) Less(i, j int) bool {
	ival := h.data[i].(int)
	jval := h.data[j].(int)
	return h.weight[ival] < h.weight[jval]
}

func (h MinHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *MinHeap) Push(val interface{}) {
	h.data = append(h.data, val)
}

func (h *MinHeap) Pop() interface{} {
	ret := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	return ret
}
