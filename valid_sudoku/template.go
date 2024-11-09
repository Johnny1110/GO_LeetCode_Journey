package valid_sudoku

import (
	"go_leetcode_journey/common"
)

type ByteSet struct {
	context map[byte]int
}

func (b *ByteSet) Add(input byte) bool {
	if input == '.' {
		return true
	}
	if b.context == nil {
		b.context = make(map[byte]int)
	}
	if _, exists := b.context[input]; exists {
		return false
	} else {
		b.context[input] = 1
		return true
	}
}

func calculateSubBoxIdx(x, y int) int {
	// Minimize x, y
	x = x / 3
	y = y / 3
	// calculate seq num
	return x*3 + y
}

func isValidSudoku(board [][]byte) bool {

	// init 9*9 Maps (9 rows, 9 cells, 9 sub-box).
	rowSets := make([]ByteSet, 9)
	cellSets := make([]ByteSet, 9)
	subBoxSets := make([]ByteSet, 9)

	for x, ray := range board {
		for y, value := range ray {
			if !rowSets[x].Add(value) {
				return false
			}
			if !cellSets[y].Add(value) {
				return false
			}

			subBoxIdx := calculateSubBoxIdx(x, y)
			if !subBoxSets[subBoxIdx].Add(value) {
				return false
			}
		}
	}

	return true
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

	common.Assert_answer(isValidSudoku(board1), true)

	board2 := [][]byte{
		{'6', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	common.Assert_answer(isValidSudoku(board2), false)

	//fmt.Println(calculateSubBoxIdx(0, 0))
	//fmt.Println(calculateSubBoxIdx(0, 1))
	//fmt.Println(calculateSubBoxIdx(0, 2))
	//fmt.Println(calculateSubBoxIdx(0, 3))
	//fmt.Println(calculateSubBoxIdx(0, 4))
	//fmt.Println(calculateSubBoxIdx(0, 5))
	//fmt.Println(calculateSubBoxIdx(0, 6))
	//fmt.Println(calculateSubBoxIdx(0, 7))
	//fmt.Println(calculateSubBoxIdx(0, 8))
}
