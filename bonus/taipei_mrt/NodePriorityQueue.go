package taipei_mrt

import "container/heap"

type Item struct {
	node     int
	priority int
	index    int
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// The priority queue is a min-heap, so Less should return true
	// If the priority of item at index a is less than the priority at index b.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	// Swap the items at indices i and j.
	// Update their indices accordingly.
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	// Push a new item onto the heap.
	n := len(*pq)
	item := x.(Item)
	item.index = n
	*pq = append(*pq, item)
	heap.Fix(pq, item.index)
}

func (pq *PriorityQueue) Pop() interface{} {
	// Pop the item with the highest priority (lowest priority value).
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = Item{} // Avoid memory leak
	*pq = old[0 : n-1]
	//heap.Fix(pq, item.index)
	return item
}

func (pq *PriorityQueue) update(item *Item, node, priority int) {
	// Update the priority and node of an item in the heap.
	item.node = node
	item.priority = priority
	heap.Fix(pq, item.index)
}
