# 19. Remove Nth Node From End

<br>

---

<br>

link: https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/

<br>

## Thinking

I did this before, and it was remarkable about using 2 pointer to solve remove element from behind of linked-list.

I read a book about MySQL, when it comes to InnoDB's buffer pool LRU (Least Recently Used). 
there is a cache with linked-list data structure. and the hot page always stay left side (front).
otherwise they will be moved to right side. if the cache size is over filled, innoDB gonna remove some cache from the end of cache.

that is typical __Remove Nth Node From End__ I guess. very elegant.

<br>

Using 2 pointer:

* Pointer-A start from head of the linked-list
* Pointer-B start from `n`th node (`n` step forward than Pointer-A)
* Moving both pointer forward until Pointer-B reach the end.
* Connect the Pointer-A the the Pointer-A's (next: deleted) -> (next) 

<br>

Let's do this.

## Coding

```go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		// because: 1 <= n <= sz, so we can just return nil
		return nil
	}

	// make 2 pointer.
	pointerA := head
	pointerB := head
	for range n {
		pointerB = pointerB.Next
	}

	// let pointer-B reach the end
	for pointerB.Next != nil {
		pointerA = pointerA.Next
		pointerB = pointerB.Next
	}

	pointerA.Next = pointerA.Next.Next

	return head
}
```

I did this and I find out a problem, if the head been removed, it will be wrong.
So we need to figure out how to prevent this happen.

I'm think about add a self-made head into the input linked-list.

and I will return the header.next, because this header is self-made, it won't be removed by this problem.

<br>

```go
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		// because: 1 <= n <= sz, so we can just return nil
		return nil
	}

	cusHeader := &ListNode{
		Val:  -1,
		Next: head,
	}

	// make 2 pointer.
	pointerA := cusHeader
	pointerB := cusHeader
	for range n {
		pointerB = pointerB.Next
	}

	// let pointer-B reach the end
	for pointerB.Next != nil {
		pointerA = pointerA.Next
		pointerB = pointerB.Next
	}

	pointerA.Next = pointerA.Next.Next

	return cusHeader.Next
}
```

<br>

![1.png](imgs/1.png)