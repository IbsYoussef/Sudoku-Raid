package utils

import "fmt"

// Board represents a 9x9 Sudoku grid
// 0 = empty cell, 1-9 = filled cell
type Board [9][9]int

// Newboard creates a new empty Sudoku Board
func NewBoard() Board {
	var board Board
	return board
}

// CharToInt converts a character ('1' - '9' or '.') to an integer
func CharToInt(c byte) int {
	if c == '.' {
		return 0
	}
	return int(c - '0')
}

// PrintBoard prints the Sudoku board in the required format
// Each row on a new line, numbers seperated by spaces
// Final empty line at the end
func PrintBoard(board *Board) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			fmt.Print(board[row][col])
			if col < 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// FindEmptyCell returns the coordinates of the next empty cell (value = 0)
// Returns (-1, -1) if no empty cells exist (board is complete)
func FindEmptyCell(board *Board) (int, int) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == 0 {
				return row, col
			}
		}
	}
	return -1, -1
}
