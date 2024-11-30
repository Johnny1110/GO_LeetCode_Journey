package sudoku_solver_enhancement

import "fmt"

// solveSudoku solves the Sudoku puzzle in-place using backtracking.
func solveSudoku(board [][]byte) {
	// TODO
}

// Debugging helper function to print the board
func debugBoard(board [][]byte) {
	for _, row := range board {
		fmt.Println(string(row))
	}
	fmt.Println()
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
