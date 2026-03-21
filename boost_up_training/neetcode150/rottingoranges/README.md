994. Rotting Oranges 

<br>

---

<br>

## Coding

```go
type Queue [][]int

func (q Queue) Len() int {
	return len(q)
}

func (q *Queue) Push(row, col int) {
	*q = append(*q, []int{row, col})
}

func (q *Queue) Pop() (int, int) {
	node := (*q)[0]
	(*q) = (*q)[1:]
	return node[0], node[1]
}

func hasFresh(grid [][]int, row, col int) bool {
	if row >= len(grid) || row < 0 || col >= len(grid[0]) || col < 0 {
		return false
	}

	if grid[row][col] == 1 {
		return true
	}

	return false
}

func orangesRotting(grid [][]int) int {

	queue := Queue([][]int{})
	orangeCount, rottenCount := 0, 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] != 0 {
				orangeCount++
			}

			if grid[row][col] == 2 {
				rottenCount++
				queue.Push(row, col)
			}
		}
	}

	if orangeCount == 0 {
		return 0
	}

	sec := 0

	dirs := [4][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	for queue.Len() > 0 {
		sec++
		thisSecLen := queue.Len()

		for range thisSecLen {
			row, col := queue.Pop()

			for _, dir := range dirs {
				nr, nc := row+dir[0], col+dir[1]
				if hasFresh(grid, nr, nc) {
					grid[nr][nc] = 2
					queue.Push(nr, nc)
					rottenCount++
				}

			}
		}

	}

	if rottenCount < orangeCount {
		return -1
	}

	return sec - 1
}
```

<br>

## Time & Space Complexity

```
Assume: n = amount of grid total slot count, h = total rotten orange

Time: O(n)

Space: O(h)
```

## AI feedback

Time is correct — every cell is visited at most once, so `O(n)`.

Space needs a correction. The queue doesn't stay at size `h`. As BFS runs, every fresh orange that gets rotted is pushed onto the queue. In the worst case (imagine a grid where every cell is a fresh orange except one corner that's rotten), nearly all `n` cells pass through the queue. So the worst-case space is `O(n)`, not `O(h)`.

The initial rotten count `h` only describes the queue's starting size, not its peak size during execution.

```
Assume: n = amount of grid total slot count.

Time: O(n)

Space: O(n)
```