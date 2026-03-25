# 21. Merge Two Sorted Lists

<br>

---

<br>

## Coding

```go
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
```

<br>
<br>

## Time & Space Complexity

```
Assume: n = length of (list1 & list2)

Time: O(2n * log n), each NodeList ned 1 pop() and 1 push(), both action need O(log n)

Space: O(n)
```

<br>
<br>

## AI Feedback

Using a Min-Heap (Priority Queue) is a clever way to handle $k$ lists, for exactly two lists, it’s actually less efficient than a simple pointer-based approach.

<br>

### Time Complexity

Your estimate of $O(2n \log n)$ is slightly off in the base of the logarithm.

* Heap Size: In your code, the heap never contains more than 2 elements at any given time (one from list1 and one from list2).

* Heap Operations: The cost of Push and Pop in a heap of size $k$ is $O(\log k)$. Since $k=2$ here, these operations are effectively $O(\log 2)$, which is constant time $O(1)$.

* Total Iterations: You iterate once for every node in both lists. If $N$ is the total number of nodes in both lists, you do $N$ pops and $N$ pushes.

* Refined Complexity: Since $\log 2$ is a constant, your time complexity is actually $O(N)$, where $N$ is the total number of nodes.

> Note: If you were merging $k$ lists, the complexity would be $O(N \log k)$.


<br>

### Space Complexity

* Heap Space: Your heap stores at most 2 nodes. This is $O(1)$ (constant space).