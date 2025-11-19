package number_of_islands

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	// define DFS func
	var dfs func(row, col int)
	dfs = func(row, col int) {
		// combined boundary check
		if row < 0 || row >= rows || col < 0 || col >= cols ||
			visited[row][col] || grid[row][col] != '1' {
			return
		}

		visited[row][col] = true

		// explore all 4 directions
		dfs(row+1, col)
		dfs(row-1, col)
		dfs(row, col-1)
		dfs(row, col+1)
	}

	// perform
	islands := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				dfs(i, j)
				islands++
			}
		}
	}

	return islands
}
