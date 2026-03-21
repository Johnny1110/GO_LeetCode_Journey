package rottingoranges

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

	var hasFresh func(grid [][]int, row, col int) bool
	hasFresh = func(grid [][]int, row, col int) bool {
		if row >= len(grid) || row < 0 || col >= len(grid[0]) || col < 0 {
			return false
		}

		if grid[row][col] == 1 {
			return true
		}

		return false
	}

	sec := 0

	for queue.Len() > 0 {
		sec++
		thisSecLen := queue.Len()

		for range thisSecLen {
			row, col := queue.Pop()

			// top
			if hasFresh(grid, row-1, col) {
				grid[row-1][col] = 2
				queue.Push(row-1, col)
				rottenCount++
			}
			// bot
			if hasFresh(grid, row+1, col) {
				grid[row+1][col] = 2
				queue.Push(row+1, col)
				rottenCount++
			}
			// left
			if hasFresh(grid, row, col-1) {
				grid[row][col-1] = 2
				queue.Push(row, col-1)
				rottenCount++
			}
			// right
			if hasFresh(grid, row, col+1) {
				grid[row][col+1] = 2
				queue.Push(row, col+1)
				rottenCount++
			}
		}

	}

	if rottenCount < orangeCount {
		return -1
	}

	return sec - 1
}
