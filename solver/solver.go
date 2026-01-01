package solver

import (
	"sudoku/utils"
	"sudoku/validator"
)

// Solve attempts to solve the sudoku board using backtracking
// Returns true if solved successfully, false if unsolvable
// Modifies the board in-place
func Solve(board *utils.Board) bool {
	// Find the next empty cell (value = 0)
	row, col := utils.FindEmptyCell(board)

	// Base case: no empty cells means board is complete
	if row == -1 {
		return true
	}

	// Try placing numbers 1-9
	for num := 1; num <= 9; num++ {
		// Check if this number is valid at this position
		if validator.IsValid(board, row, col, num) {
			// Place the number
			board[row][col] = num

			// Recursively attempt to solve the rest of the board
			if Solve(board) {
				return true // Solution found!
			}

			// Backtrack: remove the number and try next
			board[row][col] = 0
		}
	}

	// All numbers failed - this path is a dead end
	return false
}
