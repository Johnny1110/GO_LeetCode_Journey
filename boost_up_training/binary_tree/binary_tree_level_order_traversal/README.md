# 102. Binary Tree Level Order Traversal

<br>

---

<br>

link: https://leetcode.com/problems/binary-tree-level-order-traversal/description/

> Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).


<br>
<br>

## Thinking

<br>

The Topic say we should using BFS (Breadth-First Search). than let go take some time to study it.

<br>

## BFS

BFS 是「廣度優先搜尋」（Breadth-First Search）的縮寫，是一種用於圖形或樹狀結構的搜尋演算法。
其特點是從起始節點開始，一層一層地向外擴散進行搜尋。 它會先遍歷所有離起始點為1 的節點，再遍歷離起始點為2 的節點，依此類推，直到所有可達的節點都被遍歷。 
BFS 通常使用 __佇列（先進先出)__ 的資料結構來儲存待探索的節點，並用於尋找圖中的最短路徑或最少步驟的問題。


**與 DFS（深度優先搜尋）的比較**:

* BFS 採用「一層一層」的方式遍歷，善於尋找最短路徑，使用佇列。
* DFS 則採取「一條路走到黑」的方式遍歷，善於尋找特定路徑或遍歷所有節點，通常使用堆疊（或遞迴）。

<br>

Now I know about the most important thing about how to using DFS to solve this problem is 
using queue to store the next level nodes.


I think I can try like make a queue to do like.

* put head node into that queue.
* pop a node form that queue, put that node's left and right son node into queue
* pop a node from that queue again, then this time and next time popped node is the first layer's nodes.
* then we can do iterate by each node order by layer-level then left to right.

<br>

**The problem is How to know the layer?**

Now we know how to iterate through the tree structure by each layer. but we still have no clue how to identify the different layer.



<br>
<br>

## Coding