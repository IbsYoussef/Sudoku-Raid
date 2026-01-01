package validator

import "sudoku/utils"

// IsValid checks if placing num at (row, col) is valid
// Returns true if placement follows all Sudoku rules
func IsValid(board *utils.Board, row, col, num int) bool {
	// Check all three rules
	return isRowValid(board, row, num) &&
		isColValid(board, col, num) &&
		isBoxValid(board, row, col, num)
}

// isRowValid checks if num already exists in the row
func isRowValid(board *utils.Board, row, num int) bool {
	for col := 0; col < 9; col++ {
		if board[row][col] == num {
			return false // Number already in row!
		}
	}
	return true // Number not found in row
}

// isColValid checks if num already exists in the column
func isColValid(board *utils.Board, col, num int) bool {
	for row := 0; row < 9; row++ {
		if board[row][col] == num {
			return false // Number already in column
		}
	}
	return true // Number not found in column
}

// isBoxValid checks if num already exists in the 3x3 box
func isBoxValid(board *utils.Board, row, col, num int) bool {
	// Find the starting position of the 3x3 box
	boxRow := (row / 3) * 3
	boxCol := (col / 3) * 3

	// Check all 9 cells in the box
	for r := boxRow; r < boxRow+3; r++ {
		for c := boxCol; c < boxCol+3; c++ {
			if board[r][c] == num {
				return false // Number already in box
			}
		}
	}
	return true
}
