package find_median_from_data_stream

import "container/heap"

type MedianFinder struct {
	left  *Heap[int]
	right *Heap[int]
}

func Constructor() MedianFinder {
	return MedianFinder{
		left:  NewHeap(func(a, b int) bool { return a > b }),
		right: NewHeap(func(a, b int) bool { return a < b }),
	}
}

func (this *MedianFinder) AddNum(num int) {
	this.left.HPush(num)
	this.right.HPush(this.left.HPop())

	if this.right.Len() > this.left.Len()+1 {
		// balance
		this.left.HPush(this.right.HPop())
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.right.Len() == 0 {
		return 0
	} else if this.left.Len() == this.right.Len() {
		return (float64(this.left.HPeek()) + float64(this.right.HPeek())) / 2
	} else {
		return float64(this.right.HPeek())
	}
}

// =============================================

type Heap[T any] struct {
	data []T
	less func(i, j T) bool
}

func NewHeap[T any](less func(i, j T) bool) *Heap[T] {
	return &Heap[T]{
		data: []T{},
		less: less,
	}
}

func (h Heap[T]) Len() int           { return len(h.data) }
func (h Heap[T]) Less(i, j int) bool { return h.less(h.data[i], h.data[j]) }
func (h Heap[T]) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }

func (h *Heap[T]) Push(val any) {
	h.data = append(h.data, val.(T))
}

func (h *Heap[T]) Pop() any {
	ret := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	return ret
}

func (h Heap[T]) Peek() any {
	return h.data[0]
}

// wrapper

func (h *Heap[T]) HPush(val T) {
	heap.Push(h, val)
}

func (h *Heap[T]) HPop() T {
	return heap.Pop(h).(T)
}

func (h *Heap[T]) HPeek() T {
	return h.Peek().(T)
}
