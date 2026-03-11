package nqueens

func solveNQueens(n int) [][]string {
	result := [][]string{}
	chessbaord := NewChessBoard(n)

	var backtracking func(row int)
	backtracking = func(row int) {
		if row == n {
			result = append(result, chessbaord.Strings())
			return
		}

		for col := 0; col < n; col++ {

			if !chessbaord.IsSafe(row, col) {
				continue
			}

			// update state
			chessbaord.PutQueen(row, col)
			// go next row
			backtracking(row + 1)
			// row back state
			chessbaord.RemoveQueen(row, col)
		}
	}

	backtracking(0)

	return result
}

// ---------------------------------------------------

type ChessBoard struct {
	state        [][]bool
	vertical     []bool
	diagonal     map[int]bool
	antiDiagonal map[int]bool
}

func NewChessBoard(n int) *ChessBoard {

	state := make([][]bool, n)
	for i := range n {
		state[i] = make([]bool, n)
	}

	return &ChessBoard{
		state:        state,
		vertical:     make([]bool, n),
		diagonal:     make(map[int]bool),
		antiDiagonal: make(map[int]bool),
	}
}

func (b ChessBoard) IsSafe(row, col int) bool {
	// check vertical
	if b.vertical[col] {
		return false
	}

	// check digonal
	if b.diagonal[row-col] {
		return false
	}

	// check anti-digonal
	if b.antiDiagonal[row+col] {
		return false
	}

	return true
}

func (b *ChessBoard) PutQueen(row, col int) {
	b.state[row][col] = true
	b.vertical[col] = true
	b.diagonal[row-col] = true
	b.antiDiagonal[row+col] = true
}

func (b *ChessBoard) RemoveQueen(row, col int) {
	b.state[row][col] = false
	b.vertical[col] = false
	b.diagonal[row-col] = false
	b.antiDiagonal[row+col] = false
}

func (b ChessBoard) Strings() []string {
	n := len(b.state)

	result := make([]string, n)

	for i := 0; i < n; i++ {

		tmp := make([]uint8, n)

		for j := 0; j < n; j++ {
			if b.state[i][j] {
				tmp[j] = 'Q'
			} else {
				tmp[j] = '.'
			}
		}

		result[i] = string(tmp)
	}

	return result
}
