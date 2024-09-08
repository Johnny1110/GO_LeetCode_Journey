# Merge k Sorted Lists

<br>

---

<br>

## Desc

You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.


<br>

Example 1:

```
Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]

Explanation: The linked-lists are:
[
1->4->5,
1->3->4,
2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6
```
Example 2:

```
Input: lists = []
Output: []

```

Example 3:

```
Input: lists = [[]]
Output: []
```

<br>

Constraints:

```
k == lists.length
0 <= k <= 104
0 <= lists[i].length <= 500
-104 <= lists[i][j] <= 104
lists[i] is sorted in ascending order.
The sum of lists[i].length will not exceed 104.
```

<br>
<br>

## Topic

* Linked List
* [Divide and Conquer](https://www.bilibili.com/video/BV1854y1R7VG/?spm_id_from=333.337.search-card.all.click&vd_source=9780a181ac9f1fee5f680f255ee5bc73)
* Heap (Priority Queue)
* [Merge Sort](https://www.bilibili.com/video/BV1Hf4y1j7DA/?spm_id_from=333.337.search-card.all.click&vd_source=9780a181ac9f1fee5f680f255ee5bc73)

<br>

Approach:

1. [DivideAndConquer](DivideAndConquer) Solution.
2. [PriorityQueue](PriorityQueue) Solution.
3. [MergeSort](MergeSort) Solution.
