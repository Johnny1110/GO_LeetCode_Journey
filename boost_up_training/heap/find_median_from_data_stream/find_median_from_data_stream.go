package find_median_from_data_stream

import (
	"container/heap"
	"math"
)

type IntHeap struct {
	data  []int
	isMin bool
}

func NewIntHeap(min bool) *IntHeap {
	return &IntHeap{
		data:  make([]int, 0),
		isMin: min,
	}
}

func (h IntHeap) Peek() int {
	if len(h.data) == 0 {
		return math.MinInt
	}

	return h.data[0]
}

func (h IntHeap) Len() int {
	return len(h.data)
}

// Less i's priority is larger than j ?
func (h IntHeap) Less(i, j int) bool {
	if h.isMin {
		return h.data[i] < h.data[j]
	} else {
		return h.data[i] > h.data[j]
	}
}

func (h *IntHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *IntHeap) Push(x interface{}) {
	h.data = append(h.data, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	x := h.data[h.Len()-1]
	h.data = h.data[:h.Len()-1]
	return x
}

type MedianFinder struct {
	leftHeap  *IntHeap
	rightHeap *IntHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		leftHeap:  NewIntHeap(false), // max-heap
		rightHeap: NewIntHeap(true),  // min-heap
	}
}

func (this *MedianFinder) AddNum(num int) {
	// Step 1: Push to left (maxHeap)
	heap.Push(this.leftHeap, num)

	// Step 2: Balance values — move largest of left to right
	heap.Push(this.rightHeap, heap.Pop(this.leftHeap))

	// Step 3: Balance sizes — left should be >= right
	if this.rightHeap.Len() > this.leftHeap.Len() {
		heap.Push(this.leftHeap, heap.Pop(this.rightHeap))
	}
}

func (this *MedianFinder) FindMedian() float64 {

	if this.leftHeap.Len() == this.rightHeap.Len() {
		left := this.leftHeap.Peek()
		right := this.rightHeap.Peek()
		return float64(left+right) / 2
	} else {
		return float64(this.leftHeap.Peek())
	}
}
