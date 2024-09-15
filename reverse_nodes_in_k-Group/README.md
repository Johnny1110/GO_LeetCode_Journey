# Reverse Nodes in k-Group

<br>

---

<br>

## Desc

Given the head of a linked list, reverse the nodes of the list k at a time, and return the modified list.

k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.

You may not alter the values in the list's nodes, only nodes themselves may be changed.

<br>

Example 1:

![1](https://assets.leetcode.com/uploads/2020/10/03/reverse_ex1.jpg)

```
Input: head = [1,2,3,4,5], k = 2
Output: [2,1,4,3,5]
```

<br>

Example 2:

![2](https://assets.leetcode.com/uploads/2020/10/03/reverse_ex2.jpg)

```
Input: head = [1,2,3,4,5], k = 3
Output: [3,2,1,4,5]
```

<br>

Constraints:
```
The number of nodes in the list is n.
1 <= k <= n <= 5000
0 <= Node.val <= 1000
```

<br>

Follow-up: Can you solve the problem in O(1) extra memory space?

<br>
<br>

## Topic

* Linked List
* Recursion

<br>

## Thinking

![S__11771908.jpg](imgs/S__11771908.jpg)

* init a dummy head node, link to head input.
* input to recursive func with dummy node.
* implements a swap k node func like input `1-2-3` swap `3-2-1`
* swap dummy's next with k, linked dummy's next to new sorted link head
* linked new sort linked last node to next step's head
* recursive call.