package number_of_islands

type Context struct {
	visited   [][]bool
	grid      [][]byte
	maxRowCnt int
	maxColCnt int
	landCnt   int
}

func (c *Context) findOne() {
	c.landCnt++
}

func (c *Context) markAsVisited(i int, j int) {
	c.visited[i][j] = true
}

func newContext(grid [][]byte) *Context {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}
	return &Context{
		visited:   visited,
		grid:      grid,
		maxRowCnt: len(grid) - 1,
		maxColCnt: len(grid[0]) - 1,
		landCnt:   0,
	}
}

func numIslands(grid [][]byte) int {
	context := newContext(grid)

	// Iterate through the map:
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {

			if grid[i][j] == '1' && !context.visited[i][j] {
				// perform DFS recursive explore >>>
				exploreIsland(context, i, j)
				// <<< perform DFS recursive explore
				context.findOne()
			} else {
				continue
			}
		}
	}

	return context.landCnt
}

func exploreIsland(context *Context, i, j int) {
	if i < 0 || i > context.maxRowCnt || j < 0 || j > context.maxColCnt {
		// reached the edge of grid, just return.
		return
	}

	if context.grid[i][j] == '0' || context.visited[i][j] {
		// return if reach the water cell or already visited.
		return
	}

	context.markAsVisited(i, j)
	// up
	exploreIsland(context, i-1, j)
	// down
	exploreIsland(context, i+1, j)
	// left
	exploreIsland(context, i, j-1)
	// right
	exploreIsland(context, i, j+1)
}
