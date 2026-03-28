# 25. Reverse Nodes in k-Group

<br>

---

<br>

## Coding

```go
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{}

	groups := []*Group{}

	ptrA, ptrB := head, head

	for head != nil {

		for range k - 1 {
			if ptrB == nil {
				break
			}
			ptrB = ptrB.Next
		}

		if ptrB != nil {
			head = ptrB.Next
			ptrB.Next = nil
		} else {
			head = nil
		}

		group := &Group{
			head: ptrA,
			tail: ptrB,
			flag: ptrB != nil,
		}

		ptrA = head
		ptrB = head

		groups = append(groups, group)
	}

	reverseGroup(groups)

	ptr := dummy
	for _, group := range groups {
		ptr.Next = group.head
		ptr = group.tail
	}

	return dummy.Next
}

func reverseGroup(groups []*Group) {

	for _, group := range groups {
		if !group.flag || group.head == group.tail {
			return
		}

		// reverse
		curr := group.head.Next
		prev := group.head

		prev.Next = nil

		for curr != nil {
			tmp := curr.Next

			curr.Next = prev
			prev = curr
			curr = tmp
		}

		group.head, group.tail = group.tail, group.head
	}
}

type Group struct {
	head *ListNode
	tail *ListNode
	flag bool
}
```

<br>
<br>

## Time & Space Complexity

```
Assume: assume n = length of input NodeList

Time: O(n) -> group by k takes O(n), reverseGroup takes O(n), concat result takes O(n/k), overall it is O(n).

Space: O(n) -> I using a *Group[] to store group info.
```

<br>

## Claude Fedback

### Space Complexity: O(n/k), not O(n)

Your groups slice holds ⌈n/k⌉ Group structs, and each struct is constant size (two pointers + a bool). So the auxiliary space is O(n/k).

You wrote O(n), which is only true in the worst case when k=1. But since k is a parameter, the more precise answer is O(n/k). Interviewers care about this distinction because the classic in-place solution achieves O(1) space by rewiring pointers as it goes, without storing groups at all.

### The Bigger Theory Question

Your approach is essentially "partition, transform, merge" — you break the list into chunks, operate on each chunk, then stitch them back. This is a clean mental model and makes the code easy to reason about, but it trades space for clarity.

The O(1) space approach instead maintains a pointer to the node before each group and rewires in place. The key insight is that you only need to track the previous group's tail to connect it to the current group's new head after reversal — no array needed. It's the same reversal logic you already wrote, just done inline during a single pass. That would be worth exploring as a follow-up if you want to practice the in-place pointer manipulation pattern.

<br>
<br>

## Revise

```go
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	dummy := &ListNode{}
	ptrH := dummy
	ptrA, ptrB := head, head

	for head != nil {

		for range k - 1 {
			if ptrB == nil {
				break
			}
			ptrB = ptrB.Next
		}

		if ptrB != nil {
			head = ptrB.Next
			ptrB.Next = nil

			// reverse
			curr := ptrA.Next
			prev := ptrA
			prev.Next = nil

			for curr != nil {
				tmp := curr.Next

				curr.Next = prev
				prev = curr
				curr = tmp
			}

			ptrH.Next = ptrB
			ptrH = ptrA

			ptrA = head
			ptrB = head

		} else {
			head = nil
			ptrH.Next = ptrA
		}

	}

	return dummy.Next
}
```

Now Space Complexity is O(1)