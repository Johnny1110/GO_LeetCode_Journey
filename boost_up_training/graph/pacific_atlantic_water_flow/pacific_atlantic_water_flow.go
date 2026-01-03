package pacific_atlantic_water_flow

func pacificAtlantic(heights [][]int) [][]int {
	// Step-1. create data structure to mark which cell can flow water to target Pacific or target Atlantic.
	pacific := make([][]bool, len(heights))
	for i := range pacific {
		pacific[i] = make([]bool, len(heights[i]))
	}

	atlantic := make([][]bool, len(heights))
	for i := range atlantic {
		atlantic[i] = make([]bool, len(heights[i]))
	}

	rowLen := len(heights)
	colLen := len(heights[0])

	// Step-2. Write the DFS Function
	// - row, col — current cell position
	// - prevHeight — the height of the cell we came from (to check if we can flow "uphill")
	// - reachable — which mark table to update (pacific or atlantic)
	// Logic inside:
	// 1. Boundary check — is row, col within the grid?
	// 2. Already visited? — if reachable[row][col] is true, skip
	// 3. Height check — can we move here? (heights[row][col] >= prevHeight)
	// 4. Mark as reachable
	// 5. Explore all 4 directions (up, down, left, right)
	var dfs func(row, col, prevHeight int, reachable [][]bool)
	dfs = func(row, col, prevHeight int, reachable [][]bool) {
		// 1. Boundary check (if outside the grid, skip)
		if row < 0 || row > rowLen-1 || col < 0 || col > colLen-1 {
			return
		}

		// 2. Already visited? — if reachable[row][col] is true, skip
		if reachable[row][col] {
			return
		}

		// 3. Height check — can we move here? (heights[row][col] >= prevHeight)
		currentHeight := heights[row][col]
		if currentHeight >= prevHeight {
			// 4. Mark as reachable
			reachable[row][col] = true
		} else {
			return
		}

		// 5. Explore all 4 directions (up, down, left, right)
		dfs(row-1, col, currentHeight, reachable) // up
		dfs(row+1, col, currentHeight, reachable) // down
		dfs(row, col-1, currentHeight, reachable) // left
		dfs(row, col+1, currentHeight, reachable) // right
	}

	// Step-3: Call DFS from All Edge Cells
	// - from Pacific edges:
	// TOP:
	for i := range colLen {
		dfs(0, i, 0, pacific)
	}
	// Left:
	for i := range rowLen {
		dfs(i, 0, 0, pacific)
	}

	// - from Atlantic edges:
	// Bottom:
	for i := range colLen {
		dfs(rowLen-1, i, 0, atlantic)
	}

	// Right:
	for i := range rowLen {
		dfs(i, colLen-1, 0, atlantic)
	}

	// Step-4: Find the Intersection
	result := [][]int{}

	for i := range rowLen {
		for j := range colLen {
			if atlantic[i][j] && pacific[i][j] {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}
