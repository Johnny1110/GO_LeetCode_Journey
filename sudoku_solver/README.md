# Sudoku Solver

<br>

---

<br>

## Desc

Write a program to solve a Sudoku puzzle by filling the empty cells.

A sudoku solution must satisfy all of the following rules:

Each of the digits 1-9 must occur exactly once in each row.
Each of the digits 1-9 must occur exactly once in each column.
Each of the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
The '.' character indicates empty cells.

<br>

Example 1:

![1](https://upload.wikimedia.org/wikipedia/commons/thumb/f/ff/Sudoku-by-L2G-20050714.svg/250px-Sudoku-by-L2G-20050714.svg.png)

<br>

Input:
```
board = [
    ["5","3",".",".","7",".",".",".","."],
    ["6",".",".","1","9","5",".",".","."],
    [".","9","8",".",".",".",".","6","."],
    ["8",".",".",".","6",".",".",".","3"],
    ["4",".",".","8",".","3",".",".","1"],
    ["7",".",".",".","2",".",".",".","6"],
    [".","6",".",".",".",".","2","8","."],
    [".",".",".","4","1","9",".",".","5"],
    [".",".",".",".","8",".",".","7","9"]
]
```

<br>

Output:
```
[
    ["5","3","4","6","7","8","9","1","2"],
    ["6","7","2","1","9","5","3","4","8"],
    ["1","9","8","3","4","2","5","6","7"],
    ["8","5","9","7","6","1","4","2","3"],
    ["4","2","6","8","5","3","7","9","1"],
    ["7","1","3","9","2","4","8","5","6"],
    ["9","6","1","5","3","7","2","8","4"],
    ["2","8","7","4","1","9","6","3","5"],
    ["3","4","5","2","8","6","1","7","9"]
]
```

<br>

Explanation: The input board is shown above and the only valid solution is shown below:

![2](https://upload.wikimedia.org/wikipedia/commons/thumb/3/31/Sudoku-by-L2G-20050714_solution.svg/250px-Sudoku-by-L2G-20050714_solution.svg.png)

<br>

Constraints:

```
board.length == 9
board[i].length == 9
board[i][j] is a digit or '.'.
It is guaranteed that the input board has only one solution.
```

<br>
<br>

## Topic

* Array
* Hash Table
* Backtracking
* Matrix

<br>
<br>

## Thinking

The point is to understand how to utilize backtracking to solve this problem, but I have no clue how to relate this problem to a backtracking approach.

I asked Chat-GPT for some hint, and I got this:

The Sudoku Solver is well-suited for a backtracking algorithm, which systematically searches for a solution by exploring all possibilities.

Step-by-Step Thought Process:

1. Iterate Through the Grid

   * Loop over all cells in the grid.
   * Identify the empty cells ('.') where you need to place numbers.

2. Try Numbers (1-9)

   * For each empty cell, attempt to place each digit from 1 to 9.
   * Check if placing a particular number in the cell is valid:
     * Does it already exist in the current row?
     * Does it already exist in the current column? 
     * Does it already exist in the corresponding 3x3 sub-box?

3. Validate Placement
   * If the placement is valid, recursively attempt to solve the board with the updated state.

4. Backtrack When Necessary

   * If no number from 1 to 9 works for the current empty cell:
     * Undo the last placement (backtrack).
     * Try the next possibility for the previous cell.

5. Base Case

   * When all cells are filled, and the board is valid, you've found the solution.


### Example Walkthrough

Consider an empty cell at position (row, col):

* Try placing 1.
* Check if 1 is already in the row, column, or 3x3 box.
* If valid, proceed to the next empty cell.
* If no number works, backtrack.

<br>

Now I need to think about how to design the backtracking structure.

<br>

After 3 hrs trying, I finally realized that is too hard to solve this problem by myself.
I decided to follow up Chat-GPT's tutorial, and try it again after code review.

Chat-GPT answer:

```go

// solveSudoku solves the Sudoku puzzle in-place using backtracking.
func solveSudoku(board [][]byte) {
	// Backtracking function to solve the board
	var backtrack func() bool

	backtrack = func() bool {

		for i := 0; i < 9; i++ { // Iterate over rows
			for j := 0; j < 9; j++ { // Iterate over columns

				if board[i][j] == '.' { // Find the first empty cell

					for num := byte('1'); num <= byte('9'); num++ { // Try placing numbers 1 to 9

						if isValid(board, i, j, num) { // Check if the number is valid
							board[i][j] = num // Place the number
							if backtrack() {  // Recur to solve the rest of the board
								return true
							}
							// Backtrack: Remove the number if it leads to an invalid solution
							board[i][j] = '.'
						}
					}

					return false // If no number can be placed, return false
				}
			}
		}
		return true // If all cells are filled correctly, return true
	}

	backtrack() // Start backtracking
}

// isValid checks if placing a number in a specific cell is valid.
func isValid(board [][]byte, row, col int, num byte) bool {
	subBoxRowStart := (row / 3) * 3 // Starting row index of the 3x3 sub-box
	subBoxColStart := (col / 3) * 3 // Starting column index of the 3x3 sub-box

	// Check if the number already exists in the current row, column, or 3x3 sub-box
	for i := 0; i < 9; i++ {
		if board[row][i] == num { // Check row
			return false
		}
		if board[i][col] == num { // Check column
			return false
		}
		if board[subBoxRowStart+(i/3)][subBoxColStart+(i%3)] == num { // Check 3x3 sub-box
			return false
		}
	}
	return true // The number is valid
}
```