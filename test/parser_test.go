package test

import (
	"sudoku/parser"
	"testing"
)

// TestParseArgs_ValidInput tests parsing a valid sudoku
func TestParseArgs_ValidInput(t *testing.T) {
	args := []string{
		".96.4...1",
		"1...6...4",
		"5.481.39.",
		"..795..43",
		".3..8....",
		"4.5.23.18",
		".1.63..59",
		".59.7.83.",
		"..359...7",
	}

	board, err := parser.ParseArgs(args)
	// Should not error
	if err != nil {
		t.Errorf("ParseArgs() unexpected error: %v", err)
	}

	// Verify some key positions
	expectedValues := []struct {
		row, col int
		value    int
	}{
		{0, 0, 0}, // '.' → 0
		{0, 1, 9}, // '9' → 9
		{0, 2, 6}, // '6' → 6
		{0, 3, 0}, // '.' → 0
		{0, 4, 4}, // '4' → 4
		{1, 0, 1}, // '1' → 1
		{4, 1, 3}, // '3' → 3
		{8, 8, 7}, // '7' → 7
	}

	for _, expected := range expectedValues {
		actual := board[expected.row][expected.col]
		if actual != expected.value {
			t.Errorf("board[%d][%d] = %d, expected %d",
				expected.row, expected.col, actual, expected.value)
		}
	}
}

// TestParseArgs_TooFewArguments tests with less than 9 arguments
func TestParseArgs_TooFewArguments(t *testing.T) {
	testCases := []struct {
		name     string
		argCount int
	}{
		{"No arguments", 0},
		{"One argument", 1},
		{"Three arguments", 3},
		{"Eight arguments", 8},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			args := make([]string, tc.argCount)
			for i := 0; i < tc.argCount; i++ {
				args[i] = "123456789"
			}

			_, err := parser.ParseArgs(args)

			if err == nil {
				t.Errorf("ParseArgs() expected error for %d arguments, got nil", tc.argCount)
				return
			}

			expectedError := "Error: Invalid number of arguments"
			if err.Error() != expectedError {
				t.Errorf("ParseArgs() error = %q, expected %q", err.Error(), expectedError)
			}
		})
	}
}

// TestParseArgs_TooManyArguments tests with more than 9 arguments
func TestParseArgs_TooManyArguments(t *testing.T) {
	testCases := []struct {
		name     string
		argCount int
	}{
		{"Ten arguments", 10},
		{"Fifteen arguments", 15},
		{"Twenty arguments", 20},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			args := make([]string, tc.argCount)
			for i := 0; i < tc.argCount; i++ {
				args[i] = "123456789"
			}

			_, err := parser.ParseArgs(args)

			if err == nil {
				t.Errorf("ParseArgs() expected error for %d arguments, got nil", tc.argCount)
				return
			}

			expectedError := "Error: Invalid number of arguments"
			if err.Error() != expectedError {
				t.Errorf("ParseArgs() error = %q, expected %q", err.Error(), expectedError)
			}
		})
	}
}

// TestParseArgs_InvalidRowLength tests rows with wrong length
func TestParseArgs_InvalidRowLength(t *testing.T) {
	testCases := []struct {
		name      string
		rowLength int
		rowIndex  int
	}{
		{"First row too short", 8, 0},
		{"Middle row too short", 5, 4},
		{"Last row too short", 7, 8},
		{"First row too long", 10, 0},
		{"Middle row too long", 12, 4},
		{"Empty row", 0, 3},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			args := make([]string, 9)
			for i := 0; i < 9; i++ {
				if i == tc.rowIndex {
					// Create row with wrong length
					args[i] = ""
					for j := 0; j < tc.rowLength; j++ {
						args[i] += "1"
					}
				} else {
					args[i] = "123456789"
				}
			}

			_, err := parser.ParseArgs(args)

			if err == nil {
				t.Errorf("ParseArgs() expected error for row length %d, got nil", tc.rowLength)
				return
			}

			expectedError := "Error: Invalid row length"
			if err.Error() != expectedError {
				t.Errorf("ParseArgs() error = %q, expected %q", err.Error(), expectedError)
			}
		})
	}
}

// TestParseArgs_InvalidCharacters tests invalid characters
func TestParseArgs_InvalidCharacters(t *testing.T) {
	testCases := []struct {
		name     string
		char     byte
		position int
	}{
		{"Contains '0'", '0', 0},
		{"Contains lowercase 'a'", 'a', 4},
		{"Contains uppercase 'X'", 'X', 8},
		{"Contains space", ' ', 5},
		{"Contains '#'", '#', 2},
		{"Contains '@'", '@', 7},
		{"Contains '-'", '-', 3},
		{"Contains '+'", '+', 1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			args := make([]string, 9)
			for i := 0; i < 9; i++ {
				if i == 0 {
					// Create row with invalid character
					row := "123456789"
					args[i] = row[:tc.position] + string(tc.char) + row[tc.position+1:]
				} else {
					args[i] = "123456789"
				}
			}

			_, err := parser.ParseArgs(args)

			if err == nil {
				t.Errorf("ParseArgs() expected error for character '%c', got nil", tc.char)
				return
			}

			expectedError := "Error: Invalid character"
			if err.Error() != expectedError {
				t.Errorf("ParseArgs() error = %q, expected %q", err.Error(), expectedError)
			}
		})
	}
}

// TestParseArgs_AllDots tests board with all empty cells
func TestParseArgs_AllDots(t *testing.T) {
	args := []string{
		".........",
		".........",
		".........",
		".........",
		".........",
		".........",
		".........",
		".........",
		".........",
	}

	board, err := parser.ParseArgs(args)
	if err != nil {
		t.Errorf("ParseArgs() unexpected error: %v", err)
	}

	// All cells should be 0
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] != 0 {
				t.Errorf("board[%d][%d] = %d, expected 0", row, col, board[row][col])
			}
		}
	}
}

// TestParseArgs_MixedValidChars tests mix of dots and numbers
func TestParseArgs_MixedValidChars(t *testing.T) {
	args := []string{
		"1........",
		".2.......",
		"..3......",
		"...4.....",
		"....5....",
		".....6...",
		"......7..",
		".......8.",
		"........9",
	}

	board, err := parser.ParseArgs(args)
	if err != nil {
		t.Errorf("ParseArgs() unexpected error: %v", err)
	}

	// Check diagonal (should be 1-9)
	for i := 0; i < 9; i++ {
		expected := i + 1
		actual := board[i][i]
		if actual != expected {
			t.Errorf("board[%d][%d] = %d, expected %d", i, i, actual, expected)
		}
	}

	// Check that other cells are 0 (empty)
	// Test a few non-diagonal cells
	testCells := []struct {
		row, col int
	}{
		{0, 1},
		{0, 8},
		{1, 0},
		{1, 2},
		{4, 3},
		{4, 6},
		{8, 0},
		{8, 7},
	}

	for _, cell := range testCells {
		if board[cell.row][cell.col] != 0 {
			t.Errorf("board[%d][%d] = %d, expected 0",
				cell.row, cell.col, board[cell.row][cell.col])
		}
	}
}

// TestParseArgs_AllNumbers tests board with all cells filled
func TestParseArgs_AllNumbers(t *testing.T) {
	args := []string{
		"123456789",
		"456789123",
		"789123456",
		"234567891",
		"567891234",
		"891234567",
		"345678912",
		"678912345",
		"912345678",
	}

	board, err := parser.ParseArgs(args)
	if err != nil {
		t.Errorf("ParseArgs() unexpected error: %v", err)
	}

	// Verify a few specific cells
	if board[0][0] != 1 {
		t.Errorf("board[0][0] = %d, expected 1", board[0][0])
	}
	if board[0][8] != 9 {
		t.Errorf("board[0][8] = %d, expected 9", board[0][8])
	}
	if board[8][8] != 8 {
		t.Errorf("board[8][8] = %d, expected 8", board[8][8])
	}

	// Check no zeros (all filled)
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == 0 {
				t.Errorf("board[%d][%d] = 0, expected non-zero", row, col)
			}
		}
	}
}

// TestParseArgs_EdgeCases tests various edge cases
func TestParseArgs_EdgeCases(t *testing.T) {
	t.Run("All 9s", func(t *testing.T) {
		args := make([]string, 9)
		for i := 0; i < 9; i++ {
			args[i] = "999999999"
		}

		board, err := parser.ParseArgs(args)
		if err != nil {
			t.Errorf("ParseArgs() unexpected error: %v", err)
		}

		// All cells should be 9
		for row := 0; row < 9; row++ {
			for col := 0; col < 9; col++ {
				if board[row][col] != 9 {
					t.Errorf("board[%d][%d] = %d, expected 9", row, col, board[row][col])
				}
			}
		}
	})

	t.Run("All 1s", func(t *testing.T) {
		args := make([]string, 9)
		for i := 0; i < 9; i++ {
			args[i] = "111111111"
		}

		board, err := parser.ParseArgs(args)
		if err != nil {
			t.Errorf("ParseArgs() unexpected error: %v", err)
		}

		// All cells should be 1
		for row := 0; row < 9; row++ {
			for col := 0; col < 9; col++ {
				if board[row][col] != 1 {
					t.Errorf("board[%d][%d] = %d, expected 1", row, col, board[row][col])
				}
			}
		}
	})

	t.Run("Alternating dots and numbers", func(t *testing.T) {
		args := []string{
			"1.2.3.4.5",
			".6.7.8.9.",
			"1.2.3.4.5",
			".6.7.8.9.",
			"1.2.3.4.5",
			".6.7.8.9.",
			"1.2.3.4.5",
			".6.7.8.9.",
			"1.2.3.4.5",
		}

		board, err := parser.ParseArgs(args)
		if err != nil {
			t.Errorf("ParseArgs() unexpected error: %v", err)
		}

		// Check pattern on first row
		expected := []int{1, 0, 2, 0, 3, 0, 4, 0, 5}
		for col := 0; col < 9; col++ {
			if board[0][col] != expected[col] {
				t.Errorf("board[0][%d] = %d, expected %d",
					col, board[0][col], expected[col])
			}
		}
	})
}
