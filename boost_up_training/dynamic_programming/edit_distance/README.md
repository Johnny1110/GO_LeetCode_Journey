# 72. Edit Distance

<br>

---

<br>

link: https://leetcode.com/problems/edit-distance/description/

<br>

## Thinking

This one is pretty, and I have no clue how to do it, the top said using 2D array as DP.

Then HOW to define the 2D array dp?

<br>

This is a classic alo [Levenshtein distance](https://zh.wikipedia.org/zh-tw/%E8%90%8A%E6%96%87%E6%96%AF%E5%9D%A6%E8%B7%9D%E9%9B%A2)

> Levenshtein distance: 指兩個字串之間，由一個轉成另一個所需的最少編輯操作次數。
> Allowed operations: replace, insert, delete
> 俄羅斯科學家弗拉基米爾·萊文斯坦在1965年提出這個概念
> 應用: DNA分析, 拼寫檢查, 語音辨識, 抄襲偵測

<br>
<br>

### The Key Insight

* When comparing two strings, you're really making decisions character by character from some position.
* Imagine you're looking at word1[i] and word2[j]. You have to ask: "What's the minimum cost to transform word1[0...i] into word2[0...j]?"

<br>

**Define DP:**

* `dp[i][j]`: minimum cost to transform `word1[0...i]` into `word2[0...j]`.

<br>

**Example**

```
word1 = "horse", word2 = "ros"

i: 0~5
j: 0~3
```

create DP:

```json
[
  [?, ?, ?, ?],
  [?, ?, ?, ?],
  [?, ?, ?, ?],
  [?, ?, ?, ?],
  [?, ?, ?, ?],
  [?, ?, ?, ?]
]
```

<br>

init DP:

* `dp[0][0] = 0` -> empty to empty
* `dp[i][0] = i` -> delete all i chars
* `dp[0][j] = j` -> inert all j chars

```json
[
	[0, 1, 2, 3],
    [1, ?, ?, ?],
    [2, ?, ?, ?],
    [3, ?, ?, ?],
    [4, ?, ?, ?],
    [5, ?, ?, ?]
]
```

**Explain**:

* `dp[0][1]` -> from word1 0~0 to word2 0~1: empty -> "r" (insert "r" so result is 1)
* `dp[0][2]` -> from word1 0~0 ti word2 0~2: empty -> "ro" ("insert "ro" so result is 2) 
* `dp[0][3]` -> from word1 0~0 ti word2 0~3: empty -> "ros" ("insert "ros" so result is 3)

<br>

* `dp[1][0]` -> from word1 0~1 to word2 0~0: "h" -> empty (remove "h" so result is 1)
* `dp[2][0]` -> from word1 0~2 to word2 0~0: "ho" -> empty (remove "ho" so result is 2)
* `dp[3][0]` -> from word1 0~3 to word2 0~0: "hou" -> empty (remove "hou" so result is 3)
* `dp[4][0]` -> from word1 0~4 to word2 0~0: "hous" -> empty (remove "hous" so result is 4)
* `dp[5][0]` -> from word1 0~5 to word2 0~0: "house" -> empty (remove "house" so result is 5)

<br>

Define Transform Formula:

example: "cat" to "cu"

* If both chars matched: `word1[i-1] == word2[j-1]`: 
  * `dp[i][j] = dp[i-1][j-1]` No operation needed, The cost is the same as dp[i-1][j-1] like "ca" to "cu" = 0 + 1 = 1

* If both chars mismatch: `word1[i-1] != word2[j-1]` choose one of the operations:
  * Replace: `dp[i][j] = dp[i-1][j-1] + 1` (step of "ca" to "c" add 1 replace step)
  * Insert:  `dp[i][j] = dp[i][j-1] + 1` (step of "cat" to "c" add 1 insert step)
  * Delete:  `dp[i][j] = dp[i-1][j] + 1` (step of "ca" to "cu" and +1 delete step)

<br>
<br>

## Coding

```go
func minDistance(word1 string, word2 string) int {
    
}
```