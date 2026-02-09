# 23. Merge k Sorted Lists

<br>

---

<br>

link: https://leetcode.com/problems/merge-k-sorted-lists/description/

<br>
<br>

## Thinking

In previous practice, I tried this problem with divide & conqueror. now we should using `Heap` to solve this problem.


<br>

What if we just create a Min-Heap and put all node into it. it naturally can pop from min to max.

<br>

Oh wait, I think we can make a max-k length heap.

For example, we have k head in nodeList.

We put each head elements into it as init. the smallest head will be on the top of the heap.

**For loop until heap is empty**

1. Pop smallest node `x` from heap
2. If `x.Next` exists, push `x.Next` into heap
3. Append `x` to `result` linked list
4. Repeat until heap is empty

<br>
<br>

## Coding

```go
import (
	"container/heap"
)

// ListNode Go Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// define Heap
type NodeHeap []*ListNode

func (h NodeHeap) Len() int {
	return len(h)
}

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Less min heap, i's priority > j ?
func (h NodeHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	ret := old[h.Len()-1]
	*h = old[:h.Len()-1]
	return ret
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	dummyHead := &ListNode{}

	// init heap
	H := NodeHeap(make([]*ListNode, 0))
	for _, head := range lists {
		if head != nil {
			heap.Push(&H, head)
		}
	}

	if H.Len() == 0 {
		return nil
	}

	// pop first to link the dummy head
	currentPointer := heap.Pop(&H).(*ListNode)

	dummyHead.Next = currentPointer
	if currentPointer.Next != nil {
		heap.Push(&H, currentPointer.Next)
	}

	for H.Len() != 0 {
		node := heap.Pop(&H).(*ListNode)
		currentPointer.Next = node
		currentPointer = node
		if node.Next != nil {
			heap.Push(&H, node.Next)
		}
	}

	return dummyHead.Next
}
```

Result: 

![1.png](imgs/1.png)