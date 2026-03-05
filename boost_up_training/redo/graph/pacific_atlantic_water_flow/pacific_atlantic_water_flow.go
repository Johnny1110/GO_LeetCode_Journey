package pacific_atlantic_water_flow

type State int

const (
	UNVISITED State = iota
	COMPLETED
)

func pacificAtlantic(heights [][]int) [][]int {
	result := [][]int{}

	m, n := len(heights), len(heights[0])

	fromTopLeftState := make([][]State, m)
	fromBotRightState := make([][]State, m)
	for row := 0; row < m; row++ {
		fromTopLeftState[row] = make([]State, n)
		fromBotRightState[row] = make([]State, n)
	}

	var dfs func(prevHeight, row, col int, state [][]State)
	dfs = func(prevHeight, row, col int, state [][]State) {
		if row >= m || row < 0 || col >= n || col < 0 {
			return
		}

		if state[row][col] == COMPLETED {
			return
		}

		thisHeight := heights[row][col]
		if thisHeight >= prevHeight {
			state[row][col] = COMPLETED

			// keep explore
			dfs(thisHeight, row+1, col, state)
			dfs(thisHeight, row-1, col, state)
			dfs(thisHeight, row, col+1, state)
			dfs(thisHeight, row, col-1, state)
		}
	}

	// iterate from left and right
	for idx := 0; idx < m; idx++ {
		dfs(0, idx, 0, fromTopLeftState)
		dfs(0, idx, n-1, fromBotRightState)
	}

	// iterate from top and bot
	for idx := 0; idx < n; idx++ {
		dfs(0, 0, idx, fromTopLeftState)
		dfs(0, m-1, idx, fromBotRightState)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if fromTopLeftState[i][j] == COMPLETED && fromBotRightState[i][j] == COMPLETED {
				result = append(result, []int{i, j})
			}
		}
	}

	return result

}
