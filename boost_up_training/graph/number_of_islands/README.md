# 200. Number of Islands

<br>

---

<br>

link: https://leetcode.com/problems/number-of-islands

<br>
<br>

## Thinking

Given an `m x n` 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.


<br>

As I know, this is the first problem of graph series. and it marked topic with __BFS/DFS__.

<br>

### why this is a graph problem:

1. Connected Components: An island is essentially a group of connected `1`s (land cells). In graph theory, we call these "connected components" - subgraph where every node can reach every other node within that subgraph.
2. Implicit Graph Structure: The 2D grid forms an implicit graph where:
   * Each cell is a node.
   * Adjacent cells (up, down, left right) form edges.
   * We only care about edges between land cells (`1`s).

<br>

### Core Insight

The key insight is: counting islands = counting connected components of `1`s

<br>

* If we stand on any land cell (`1`), we can walk to all connected land cells without crossing water zone.
* All cells we can reach form one island.
* Once we've visited an entire island, any unvisited `1` must belong to a different island.

<br>

### Approach

**Approach-1: Depth-First Search (DFS)**: The Exploration Approach

Imagine we are explorer, when we find a new land:

1. Mark the current position as "visited".
2. Explore in all 4 directions recursively.
3. Keep going until we hit water or hit boundaries.

<br>

**Why DFS?**

* DFS naturally explores an entire connected component before moving on.
* It uses the call stack to remember where we've been
* Perfect for "flood fill" type problems where we want to mark/process an entire region.

<br>
<br>

**Approach-2: Breadth-First Search (BFS)**: The Layer-by-Layer Approach

instead of diving depp immediately, we explore layer by layer:

1. Start from a land cell.
2. Visit all immediate neighbors first.
3. Then visit neighbors of neighbors.
4. Continue until the entire island is mapped.

<br>

**Why BFS?**

* BFS also explores entire connected components
* Uses a queue instead of recursion
* Better if we care about "Distance" or "Short Path" (not needed by this problem)

<br>
<br>

### Build Graph Problem Intuition

When you see a grid problem, ask yourself:

1. What are the nodes? (each cell)
2. What are the edges? (adjacency between land cells)
3. What am I counting/finding? (number of separate component)
4. Do I need to track visited nodes? (almost always yes in grid problems)
5. DFS or BFS? (Both work for connected components; DFS is often simpler)

<br>
<br>

### Common Patterns This Problem Teaches

* **Flood Fill Pattern**: Starting from a point, "fill" all connected cells.
* **Connected Components**: Count separate groups in a graph.
* **Grid as Graph**: 2D arrays often represent implicit graphs.
* **Marking Visited**: Modify in-place or use separate visited set.

<br>
<br>

## Coding - DFS

<br>

1. I want to create a 2D bool array as "visited" land mark array.
2. Iterate through the map:
   * If the node is a __land cell__ `&&` __not been visited__:
     * Start DFS recursive call to explore all the connected component from current node, mark al of them as "visited" (in 2D bool array).
     * After the explore call is over, count the discovered island + 1.
   * If the node is __water cell__ `||` __been visited__:
     * Skip to next cell.
* Return the discovered island count as result.

<br>

```go
func numIslands(grid [][]byte) int {
    
}
```

<br>
<br>

## Coding - BFS

```go
func numIslands(grid [][]byte) int {
    
}
```
