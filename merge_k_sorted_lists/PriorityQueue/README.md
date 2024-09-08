# Priority Queue Solution


<br>

---

<br>

make a priority queue.

put all node to PriorityQueue and then link them 1 by 1.

make a PriorityQueue in golang:

```go
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
```


