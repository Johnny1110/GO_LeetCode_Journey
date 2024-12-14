package sudoku_solver_enhancement

import "fmt"

const boardSize = 9

type board struct {
	board [boardSize][boardSize]int

	usedInRow, usedInCol, usedInSq [boardSize][boardSize + 1]bool // why boardSize+1 ? because it's contains 0~9 total 10 elements
}

func calculateSqlIndex(r, c int) int {
	return r/3*3 + c/3
}

func (b *board) set(r, c, v int) {
	b.board[r][c] = v
	b.usedInRow[r][v] = true
	b.usedInCol[c][v] = true
	b.usedInSq[calculateSqlIndex(r, c)][v] = true
}

func (b *board) reset(r, c int) {
	v := b.board[r][c]
	b.board[r][c] = 0
	b.usedInRow[r][v] = false
	b.usedInCol[c][v] = false
	b.usedInSq[calculateSqlIndex(r, c)][v] = false
}

func (b *board) ableToSet(r, c, v int) bool {
	return !b.usedInRow[r][v] &&
		!b.usedInCol[c][v] &&
		!b.usedInSq[calculateSqlIndex(r, c)][v]
}

func (b *board) save(rawBoard [][]byte) {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			rawBoard[i][j] = byte(b.board[i][j]) + '0'
		}
	}
}

func (b *board) isSet(r, c int) bool {
	return b.board[r][c] != 0
}

func (b *board) printDiagram() {
	for _, row := range b.board {
		fmt.Println(row)
	}
	fmt.Println()
}

func initBoard(input [][]byte) board {
	var b board

	for r := 0; r < boardSize; r++ {
		for c := 0; c < boardSize; c++ {
			if input[r][c] != '.' {
				v := int(input[r][c] - '0')
				b.set(r, c, v)
			}
		}
	}
	return b
}

func solve(r, c int, b *board) bool {
	// all row done. -> final break point.
	if r == boardSize {
		return true
	}
	// current row's all col done. -> move to next row.
	if c == boardSize {
		return solve(r+1, 0, b)
	}

	// current slot has number. -> move to next slot.
	if b.isSet(r, c) {
		return solve(r, c+1, b)
	}

	for i := 1; i <= boardSize; i++ {
		if b.ableToSet(r, c, i) {
			b.set(r, c, i)
			if solve(r, c+1, b) {
				return true // this return will be triggered when reached line 65.
			}
			b.reset(r, c)
		}
	}

	// all 1~9 are not able to fill in
	return false
}

func solveSudoku(rawBoard [][]byte) {
	b := initBoard(rawBoard)
	solve(0, 0, &b)
	b.save(rawBoard)
}

// Go call this func in main.go
func Go() {
	board1 := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	solveSudoku(board1)
	debugBoard(board1)
}

// Debugging helper function to print the board
func debugBoard(board [][]byte) {
	for _, row := range board {
		fmt.Println(string(row))
	}
	fmt.Println()
}
