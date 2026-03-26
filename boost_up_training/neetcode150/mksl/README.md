# 23. Merge k Sorted Lists

<br>

---

<br>

## Coding

```go
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	}

	H := &MinHeap{
		data: []*ListNode{},
	}

	for _, head := range lists {
		if head != nil {
			heap.Push(H, head)
		}
	}

	dummy := &ListNode{}
	pointer := dummy

	for H.Len() > 0 {
		node := heap.Pop(H).(*ListNode)

		pointer.Next = node
		pointer = node
		tmp := node.Next
		node.Next = nil

		if tmp != nil {
			heap.Push(H, tmp)
		}
	}

	return dummy.Next

}

type MinHeap struct {
	data []*ListNode
}

func (h MinHeap) Len() int           { return len(h.data) }
func (h MinHeap) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h MinHeap) Less(i, j int) bool { return h.data[i].Val < h.data[j].Val }

func (h *MinHeap) Pop() interface{} {
	ret := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	return ret
}

func (h *MinHeap) Push(node interface{}) {
	h.data = append(h.data, node.(*ListNode))
}
```

<br>
<br>

## Time & Space Complexity

```
Assume: n = length of nodelist, m = all nodes count

Time: 0(m * 2 * log n) -> loop m times, each time need 2 log n operations (Heap Pop and Heap Push).

Space: O(n) max Heap size is n. 
```