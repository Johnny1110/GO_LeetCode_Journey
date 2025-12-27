# 417. Pacific Atlantic Water Flow

<br>

---

<br>

link: https://leetcode.com/problems/pacific-atlantic-water-flow/description/

<br>

---

<br>

## Thinking

### Constraints:

```
m == heights.length
n == heights[r].length
1 <= m, n <= 200
0 <= heights[r][c] <= 105
```

### Topics:

* Array
* Depth-First Search
* Breadth-First Search
* Matrix

<br>

### Deconstruct the Problem

Check every cells, explore to **Pacific (Up, Left) and Atlantic (Down , Right)**.

If can reach the 1 edge of Pacific and 1 edge of Atlantic, that means we found one of the results.

<br>

### Refine the Approach

__Start from each cell → Can I reach Pacific? Can I reach Atlantic?__

This means for each cell, you might do two separate explorations. With an `m × n` grid, that's potentially expensive.

**Reverse the Direction**

Instead of asking:

> "From this cell, can water flow to the ocean?"

What if you asked:

> "From the ocean, which cells can water flow from?"


* Water flows from higher or equal height to lower or equal height
* If we reverse the perspective, you're asking: **starting from the ocean edge, which cells can reach this ocean (going uphill or flat)?**

<br>

So, I guess, we should make 2 DFS searching, starting from Pacific side and Atlantic side.

Using a data structure to mark which cell can flow water to target Pacific or target Atlantic,

and finally, return cascade of 2 marks table.

data structure:

```go
pacific := make([][]bool, m)
atlantic := make([][]bool, m)
```

<br>

---

<br>

## Coding - DFS

```go
func pacificAtlantic(heights [][]int) [][]int {
    
}
```

<br>

---

<br>

## Coding - BFS

```go
func pacificAtlantic(heights [][]int) [][]int {
    
}
```