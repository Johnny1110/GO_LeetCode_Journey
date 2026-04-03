# 79. Word Search

<br>

---

<br>

Given an `m x n` grid of characters `board` and a string `word`, return `true` if `word` exists in the grid.

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are **horizontally or vertically** neighboring. The **same letter cell may not be used more than once**.

<br>

**Example 1:**

```
Input: board = [["A","B","C","E"],
                ["S","F","C","S"],
                ["A","D","E","E"]], word = "ABCCED"
Output: true
```

**Example 2:**

```
Input: board = [["A","B","C","E"],
                ["S","F","C","S"],
                ["A","D","E","E"]], word = "SEE"
Output: true
```

**Example 3:**

```
Input: board = [["A","B","C","E"],
                ["S","F","C","S"],
                ["A","D","E","E"]], word = "ABCB"
Output: false  (cannot reuse the same cell)
```

**Constraints:**
- `m == board.length`, `n == board[i].length`
- `1 <= m, n <= 6`
- `1 <= word.length <= 15`
- `board` and `word` consist of only lowercase and uppercase English letters

<br>

---

<br>

## 💡 Hints

<details>
<summary>Hint 1 - Where to start?</summary>

Try every cell as a potential starting point. If `board[r][c] == word[0]`, kick off a DFS from there.

</details>

<details>
<summary>Hint 2 - What makes this backtracking?</summary>

You need to **mark** a cell as visited while exploring a path, then **unmark** it when you backtrack — so other paths can reuse it.

</details>

<details>
<summary>Hint 3 - Termination conditions</summary>

- ✅ Base case (success): `index == len(word)` — you've matched all characters
- ❌ Base case (fail): out of bounds, cell already visited, or `board[r][c] != word[index]`

</details>

<br>

---

<br>

## Coding

```go
// TODO
```

<br>

## Time & Space Complexity

```
Assume: m = board rows, n = board cols, L = word length

Time: ?

Space: ?
```

<br>

## AI Feedback

> Fill in after solving and self-assessing complexity.
