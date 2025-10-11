# Reorder List

<br>

---

<br>

link: https://leetcode.com/problems/reorder-list/description/

<br>

## Thinking

<br>

![1.jpg](imgs/1.jpg)

Using 2 pointer, one start from head another from the tail. but it's not a Doubly Linked List, so I can't use a pointer backward from the tail.

<br>

**Topic**

* Linked List
* Two Pointers
* Stack
* Recursion

<br>

Stack is interesting, and me think how does that work in this problem.

Stack's feature is fist in last out.

What if I put all node in to it?

<br>

ex: [1,2,3,4,5]

```go
|5|
|4|
|3|
|2|
|1|
```

<br>

I think that's means I can do backward by stack.

<br>
<br>

## Coding

```go
type ListNode struct {
	Val  int
	Next *ListNode
}

type ListNodeStack struct {
	container []*ListNode
}

func NewListNodeStack() *ListNodeStack {
	return &ListNodeStack{
		container: make([]*ListNode, 0),
	}
}

func (s *ListNodeStack) Push(node *ListNode) {
	s.container = append(s.container, node)
}

func (s *ListNodeStack) Pop() (*ListNode, error) {
	length := len(s.container)
	if length == 0 {
		return nil, fmt.Errorf("stack is empty")
	}

	node := s.container[length-1]
	s.container = s.container[:length-1]
	return node, nil
}

func (s *ListNodeStack) isEmpty() bool {
	return len(s.container) == 0
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// put all node(except head) into stack
	stack := NewListNodeStack()
	tmpPointer := head.Next
	for tmpPointer != nil {
		stack.Push(tmpPointer)
		tmpPointer = tmpPointer.Next
	}

	pointerA := head
	pointerB, _ := stack.Pop() // tail

	for (pointerA != pointerB) && (pointerA.Next != pointerB) {
		fmt.Println("pointerA:", pointerA, "pointerB:", pointerB)
		if stack.isEmpty() {
			break
		}

		CPA := pointerA
		CPB := pointerB

		pointerB, _ = stack.Pop()
		pointerA = pointerA.Next

		CPA.Next = CPB
		CPB.Next = pointerA
	}

	pointerB.Next = nil
}
```

<br>

result:

![2.png](imgs/2.png)

<br>

I passed that problem, and performance is perfect.
But I don't think this is a good way to solve this problem, because I feel like the code I wrote was hard to read and understand.
I should revamp this.

<br>
<br>

### Claude AI:

**The Optimal O(1) Space Approach**

Theory: Break the problem into 3 steps:

* Find the middle (slow/fast pointer)
* Reverse the second half (in-place pointer manipulation)
* Merge two halves (interleave/交錯)

<br>

### Step 1: Find Middle

```
1 → 2 → 3 → 4 → 5
        ↑
      middle
```

Using slow/fast pointers (slow moves 1, fast moves 2), when fast reaches end, slow is at middle.

<br>

### Step 2: Reverse Second Half

```
Before: 3 → 4 → 5
After:  3 ← 4 ← 5
        ↓
       null
```

This is the classic "reverse linked list" pattern—changing direction of pointers in-place.

<br>

### Step 3: Merge

```
List 1: 1 → 2 → 3
List 2: 5 → 4

Result: 1 → 5 → 2 → 4 → 3
```

Alternate picking from each list.

<br>
<br>

### Tips

**The Pattern to Remember** 

Many linked list problems have this pattern:

```
"I need to access from both ends"
    ↓
Option A: Store in auxiliary structure (O(n) space)
Option B: Find middle + reverse half (O(1) space)
```

<br>

## Revamp

### Coding

```go

```