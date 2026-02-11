# 79. Word Search

<br>

---

<br>

link: https://leetcode.com/problems/word-search/

<br>

Given an `m x n` grid of characters board and a string word, return true if word exists in the grid.

The word can be constructed from letters of sequentially adjacent cells,
where adjacent cells are horizontally or vertically neighboring.

The same letter cell may not be used more than once.

<br>
<br>

## Thinking - Backtracking

When it comes to backtracking, I always think about decision tree.

When I iterate to a new letter here is 2 options:

Does this cell match word[matchIndex]?

* Yes → explore 4 directions deeper
* No → return false (backtrack automatically)

backtracking return: 
* return true when `matchIndex == len(word)`
* return false if can't match word's index in current (row, col).

We also need a `[][]bool` to mark which position already visited.



<br>

## Coding - Backtracking

```go

```

<br>
<br>