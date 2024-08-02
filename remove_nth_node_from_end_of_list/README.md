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





