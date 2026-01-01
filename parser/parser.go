package parser

import (
	"errors"
	"sudoku/utils"
)

// ParseArgs converts command-line arguments into a Sudoku board
// Returns error if input is invalid
func ParseArgs(args []string) (utils.Board, error) {
	// Step 1: Validate argument count
	if len(args) != 9 {
		return utils.Board{}, errors.New("Error: Invalid number of arguments")
	}

	// Step 2: Create empty board
	board := utils.NewBoard()

	// Step 3: Parse each argument (row)
	for row := 0; row < 9; row++ {
		// Step 4: Validate row length
		if len(args[row]) != 9 {
			return utils.Board{}, errors.New("Error: Invalid row length")
		}

		// Step 5: Parse each character (column)
		for col := 0; col < 9; col++ {
			char := args[row][col]
			// Step 6: Validate character
			if !isValidChar(char) {
				return utils.Board{}, errors.New("Error: Invalid character")
			}

			// Step 7: Convert and store
			board[row][col] = utils.CharToInt(char)
		}
	}

	return board, nil
}

// isValidChar check if a character is valid ('.' or '1-9')
func isValidChar(c byte) bool {
	return c == '.' || (c >= '1' && c <= '9')
}
