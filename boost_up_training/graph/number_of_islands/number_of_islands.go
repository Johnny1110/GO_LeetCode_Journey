package number_of_islands

type Queue struct {
	pipe [][]int
}

func (q *Queue) Push(row, col int) {
	q.pipe = append(q.pipe, []int{row, col})
}

func (q *Queue) Pop() (int, int, bool) {
	if len(q.pipe) == 0 {
		return 0, 0, false
	}
	box := q.pipe[0]
	q.pipe = q.pipe[1:]
	return box[0], box[1], true
}

func NewQueue() *Queue {
	queue := &Queue{pipe: [][]int{}}
	return queue
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// define BFS closure func
	bfs := func(startRow, startCol int) {
		queue := NewQueue()
		queue.Push(startRow, startCol)
		visited[startRow][startCol] = true

		// loop to explore land cells from queue.
		for {
			row, col, ok := queue.Pop()
			if !ok {
				break
			}

			// check all 4 neighbors
			for _, direction := range directions {
				newRow, newCol := row+direction[0], col+direction[1]

				// check boundaries & conditions
				if newRow >= 0 && newRow < rows && // in map
					newCol >= 0 && newCol < cols && // in map
					grid[newRow][newCol] == '1' && // is land cell
					!visited[newRow][newCol] { // not been visited
					visited[newRow][newCol] = true // mark before queuing
					queue.Push(newRow, newCol)
				}
			}
		}
	}

	// using BFS to explore whole island
	island := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				bfs(i, j)
				island++
			}
		}
	}

	return island
}
