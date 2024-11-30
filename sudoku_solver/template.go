package sudoku_solver

import "fmt"

// solveSudoku solves the Sudoku puzzle in-place using backtracking.
func solveSudoku(board [][]byte) {

	var backtracking func() bool
	// Define backtracking function to solve the board
	backtracking = func() bool {
		// 1. Iterate through all slot
		for x := 0; x < 9; x++ {
			for y := 0; y < 9; y++ {
				// 2. Try to find next '.'
				if board[x][y] == '.' {
					// 3. Iterate thorough '1' ~ '9'
					for num := byte('1'); num <= byte('9'); num++ {
						// 4. Check is valid (row, col, subBox)
						if isValid(board, num, x, y) {
							board[x][y] = num
							// 5. If next backtracking (next slot) unable to fill out num, rollback board[x][y] = '.'
							if !backtracking() {
								board[x][y] = '.'
							} else {
								return true
							}
						}
					}
					// 6. If all nums ('1' ~ '9 ) can not fit in current board[x][y], return false, redo previous step
					return false
				}
			}
		}
		// 7. All slots already iterated.
		return true
	}

	backtracking()

}

func isValid(board [][]byte, num byte, x int, y int) bool {
	// 1. Calculate subBox x, y
	subBoxStartX := (x / 3) * 3
	subBoxStartY := (y / 3) * 3

	// 2. Iterate current row, col, subBox
	for idx := 0; idx < 9; idx++ {

		// 3. Check num is valid or not
		if board[x][idx] == num ||
			board[idx][y] == num ||
			board[subBoxStartX+idx%3][subBoxStartY+idx/3] == num {
			return false
		}
	}
	return true
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
