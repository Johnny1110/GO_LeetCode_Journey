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
// ListNode Go Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	// TODO
}
```