# 295. Find Median from Data Stream

<br>

---

<br>

link: https://leetcode.com/problems/find-median-from-data-stream/description/

<br>
<br>

## Thinking

### Topic:

* Two Pointers
* Design
* Sorting
* Heap (Priority Queue)
* Data Stream

<br>

### Approach we can try:

1. Insertion Sort — easiest to understand
2. Two Heaps — the "aha!" solution, most elegant
3. Balanced BST — More complex to implement


<br>

---

<br>

## Insertion Sort

### The Idea

Keep a sorted slice. When a new number comes:

1. Find the correct position (binary search)
2. Insert it there

**How do you find the insertion position efficiently?**

I think we can use a array to store the income data as data stream. 
and we're using binary search to find the target position index.

**After finding the position, how do you insert in Go?**

Just like regular array insert into the middle of the data array, and move every right element forward to right side.

**What's the complexity of each operation?**

Find the position index is O(Log N), insert action is O(N)

<br>
<br>

### Coding

```go

```
