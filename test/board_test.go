package test

import (
	"sudoku/utils"
	"testing"
)

// TestNewBoard verifies that a new board is created with all zeros
func TestNewBoard(t *testing.T) {
	board := utils.NewBoard()

	// Check that all cells are 0 (empty)
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] != 0 {
				t.Errorf("NewBoard() cell (%d, %d) = %d, expected 0",
					row, col, board[row][col])
			}
		}
	}
}

// TestCharToInt verifies character to integer conversion
func TestCharToInt(t *testing.T) {
	// Test cases: input character → expected output
	tests := []struct {
		input    byte
		expected int
	}{
		{'.', 0}, // Empty cell
		{'1', 1}, // Digit 1
		{'2', 2}, // Digit 2
		{'5', 5}, // Digit 5
		{'9', 9}, // Digit 9
	}

	// Run all test cases
	for _, test := range tests {
		result := utils.CharToInt(test.input)
		if result != test.expected {
			t.Errorf("CharToInt(%c) = %d, expected %d",
				test.input, result, test.expected)
		}
	}
}

// TestFindEmptyCell verifies finding empty cells
func TestFindEmptyCell(t *testing.T) {
	// Test 1: Board with empty cell at (0, 0)
	t.Run("EmptyCell at start", func(t *testing.T) {
		board := utils.NewBoard()
		row, col := utils.FindEmptyCell(&board)

		if row != 0 || col != 0 {
			t.Errorf("FindEmptyCell() = (%d, %d), expected (0, 0)", row, col)
		}
	})

	// Test 2: Board with first empty cell in middle
	t.Run("EmptyCell in middle", func(t *testing.T) {
		board := utils.NewBoard()

		// Fill first row and part of second
		for col := 0; col < 9; col++ {
			board[0][col] = col + 1
		}
		board[1][0] = 1
		board[1][1] = 2

		// First empty cell should be at (1, 2)
		row, col := utils.FindEmptyCell(&board)

		if row != 1 || col != 2 {
			t.Errorf("FindEmptyCell() = (%d, %d), expected (1, 2)", row, col)
		}
	})

	// Test 3: Completely filled board
	t.Run("No empty cells", func(t *testing.T) {
		board := utils.NewBoard()

		// Fill entire board with numbers
		for row := 0; row < 9; row++ {
			for col := 0; col < 9; col++ {
				board[row][col] = (row+col)%9 + 1
			}
		}

		// Should return (-1, -1) for complete board
		row, col := utils.FindEmptyCell(&board)

		if row != -1 || col != -1 {
			t.Errorf("FindEmptyCell() = (%d, %d), expected (-1, -1)", row, col)
		}
	})
}

// TestPrintBoard verifies board printing format
func TestPrintBoard(t *testing.T) {
	t.Run("Simple 1-9 pattern", func(t *testing.T) {
		board := utils.NewBoard()

		// Fill board with a simple pattern
		for row := 0; row < 9; row++ {
			for col := 0; col < 9; col++ {
				board[row][col] = (row+col)%9 + 1
			}
		}

		// Expected output
		expected := "1 2 3 4 5 6 7 8 9\n" +
			"2 3 4 5 6 7 8 9 1\n" +
			"3 4 5 6 7 8 9 1 2\n" +
			"4 5 6 7 8 9 1 2 3\n" +
			"5 6 7 8 9 1 2 3 4\n" +
			"6 7 8 9 1 2 3 4 5\n" +
			"7 8 9 1 2 3 4 5 6\n" +
			"8 9 1 2 3 4 5 6 7\n" +
			"9 1 2 3 4 5 6 7 8\n" +
			"\n"

		// Capture actual output
		actual := captureOutput(func() {
			utils.PrintBoard(&board)
		})

		// Compare
		if actual != expected {
			t.Errorf("PrintBoard() output mismatch")
			t.Logf("Expected:\n%s", expected)
			t.Logf("Got:\n%s", actual)
		}
	})

	t.Run("Empty board (all zeros)", func(t *testing.T) {
		board := utils.NewBoard()

		// Expected: all zeros
		expected := "0 0 0 0 0 0 0 0 0\n" +
			"0 0 0 0 0 0 0 0 0\n" +
			"0 0 0 0 0 0 0 0 0\n" +
			"0 0 0 0 0 0 0 0 0\n" +
			"0 0 0 0 0 0 0 0 0\n" +
			"0 0 0 0 0 0 0 0 0\n" +
			"0 0 0 0 0 0 0 0 0\n" +
			"0 0 0 0 0 0 0 0 0\n" +
			"0 0 0 0 0 0 0 0 0\n" +
			"\n"

		actual := captureOutput(func() {
			utils.PrintBoard(&board)
		})

		if actual != expected {
			t.Errorf("PrintBoard() empty board output mismatch")
			t.Logf("Expected:\n%s", expected)
			t.Logf("Got:\n%s", actual)
		}
	})

	t.Run("Solved sudoku example", func(t *testing.T) {
		board := utils.Board{
			{3, 9, 6, 2, 4, 5, 7, 8, 1},
			{1, 7, 8, 3, 6, 9, 5, 2, 4},
			{5, 2, 4, 8, 1, 7, 3, 9, 6},
			{2, 8, 7, 9, 5, 1, 6, 4, 3},
			{9, 3, 1, 4, 8, 6, 2, 7, 5},
			{4, 6, 5, 7, 2, 3, 9, 1, 8},
			{7, 1, 2, 6, 3, 8, 4, 5, 9},
			{6, 5, 9, 1, 7, 4, 8, 3, 2},
			{8, 4, 3, 5, 9, 2, 1, 6, 7},
		}

		expected := "3 9 6 2 4 5 7 8 1\n" +
			"1 7 8 3 6 9 5 2 4\n" +
			"5 2 4 8 1 7 3 9 6\n" +
			"2 8 7 9 5 1 6 4 3\n" +
			"9 3 1 4 8 6 2 7 5\n" +
			"4 6 5 7 2 3 9 1 8\n" +
			"7 1 2 6 3 8 4 5 9\n" +
			"6 5 9 1 7 4 8 3 2\n" +
			"8 4 3 5 9 2 1 6 7\n" +
			"\n"

		actual := captureOutput(func() {
			utils.PrintBoard(&board)
		})

		if actual != expected {
			t.Errorf("PrintBoard() solved sudoku output mismatch")
			t.Logf("Expected:\n%s", expected)
			t.Logf("Got:\n%s", actual)
		}
	})
}
