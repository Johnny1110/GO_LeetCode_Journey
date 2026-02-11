package n_queens

type NQueensContext struct {
	vertical     map[int]bool
	diagonal     map[int]bool
	antiDiagonal map[int]bool

	currentState [][]bool
	result       [][]string
	n            int
}

func (c *NQueensContext) collectCurrentState2Result() {
	collected := make([]string, c.n)

	for i := 0; i < c.n; i++ {

		buf := make([]uint8, c.n)
		for j := 0; j < c.n; j++ {
			if c.currentState[i][j] {
				buf[j] = 'Q'
			} else {
				buf[j] = '.'
			}
		}

		collected[i] = string(buf)
	}

	c.result = append(c.result, collected)
}

func (c *NQueensContext) tryPutQueen(row, col int) bool {
	// do safe zone check
	if c.vertical[col] || c.diagonal[row-col] || c.antiDiagonal[row+col] {
		return false // col already taken
	}
	// mark taken
	c.vertical[col] = true
	c.diagonal[row-col] = true
	c.antiDiagonal[row+col] = true
	// update state
	c.currentState[row][col] = true

	return true
}

func (c *NQueensContext) undo(row, col int) {
	// update state
	c.currentState[row][col] = false
	// mark taken
	c.vertical[col] = false
	c.diagonal[row-col] = false
	c.antiDiagonal[row+col] = false
}

func NewNQueensContext(n int) *NQueensContext {
	currentState := make([][]bool, n)
	for i := range n {
		currentState[i] = make([]bool, n)
	}

	return &NQueensContext{
		vertical:     make(map[int]bool),
		diagonal:     make(map[int]bool),
		antiDiagonal: make(map[int]bool),
		currentState: currentState,
		result:       make([][]string, 0),
		n:            n,
	}
}

func solveNQueens(n int) [][]string {
	// init current state
	checker := NewNQueensContext(n)
	backtracking(checker, 0)
	return checker.result
}

func backtracking(context *NQueensContext, row int) {
	if row == context.n {
		// collect currentState to result
		context.collectCurrentState2Result()
		return
	}

	for col := 0; col < context.n; col++ {
		if ok := context.tryPutQueen(row, col); ok {
			// go deeper
			backtracking(context, row+1)
			// undo state
			context.undo(row, col)
		}
	}
}
