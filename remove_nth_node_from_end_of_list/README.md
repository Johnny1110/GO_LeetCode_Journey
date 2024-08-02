# Remove Nth Node From End of List

<br>

---

<br>

## Desc

Given the head of a linked list, remove the nth node from the end of the list and return its head.

<br>

![1](https://assets.leetcode.com/uploads/2020/10/03/remove_ex1.jpg)

<br>

Example 1:

```
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
```


Example 2:
```
Input: head = [1], n = 1
Output: []
```
Example 3:
```
Input: head = [1,2], n = 1
Output: [1]
```

Constraints:
```
The number of nodes in the list is sz.
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz
```

<br>
<br>

## Topic

* Linked List
* Two Pointers

<br>

## Thinking

Topic said Linked List and Two Pointers. About Linked List, absolutely right,
but I'm a little bit confused about Two Pointers.

I need think about why should to utilized Two Pointers.

<br>

```go
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	return nil
}
```

So the ListNode is a singly linked list, if I want to find a node from the end 
I should reach to the end first.

I may want to iterate through the ListNode, then put all of their values to the temp ArrayList.
like:

```
ListNode: 1 -> 2 -> 3 -> 4 -> 5

ArrayList: [1, 2, 3, 4, 5]
```

<br>

using ArrayList to find the target values that we need to remove.

```java
targetValue = ArrayList[len(ArrayList) - n]
// so that:
targetValue = ArrayList(3) = 4
targetIndex = 3
```

then:

```java
ListNode[targetIndex-1].Next = ListNode[targetIndex].Next
// so that:
ListNode[2] -> ListNode[4]
```

<br>

all done.

<br>

I don't know how to utilized tow pointer to solve this problem.
I asked ChatGPT for the approach:

<br>

## Two-Pointers Approach

The idea behind the two-pointers technique is to use two pointers that are spaced n nodes apart. This way, when the second pointer reaches the end of the list, the first pointer will be at the node just before the one you want to remove.

Here's how you can implement it:

Initialize two pointers: fast and slow, both pointing to the head of the list.
Move fast pointer n steps ahead: This ensures that the gap between fast and slow is n nodes.
Move both pointers together: until fast reaches the end of the list. At this point, slow will be just before the node you want to remove.
Adjust the pointers: to remove the nth node from the end.

<br>

```go
package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// Create a dummy node that points to the head of the list
	dummy := &ListNode{Next: head}
	// Initialize two pointers, both pointing to the dummy node
	first, second := dummy, dummy

	// Move first pointer n+1 steps ahead
	for i := 0; i <= n; i++ {
		first = first.Next
	}

	// Move both pointers together until first reaches the end
	for first != nil {
		first = first.Next
		second = second.Next
	}

	// Remove the nth node from the end
	second.Next = second.Next.Next

	// Return the head of the modified list
	return dummy.Next
}

// Helper function to print the list
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

// Example usage
func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	n := 2
	newHead := removeNthFromEnd(head, n)
	printList(newHead) // Output: 1 2 3 5
}
```

Explanation:

* Dummy Node: A dummy node is used to handle edge cases, such as removing the first node.
* Two Pointers: first is moved n + 1 steps ahead so that the distance between first and second is n nodes.
* Move Together: Both pointers are moved together until first reaches the end. This ensures that second is just before the node to be removed.
* Remove Node: second.Next is adjusted to skip the node to be removed.

This approach has a time complexity of O(sz) and a space complexity of O(1), making it efficient and elegant.





