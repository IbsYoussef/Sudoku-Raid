package test

import (
	"sudoku/utils"
	"sudoku/validator"
	"testing"
)

// TestIsValid_EmptyBoard verifies that all numbers 1-9 are valid
// at any position on an empty board
func TestIsValid_EmptyBoard(t *testing.T) {
	board := utils.NewBoard()

	// Test various positions across the board
	testCases := []struct {
		row, col, num int
	}{
		{0, 0, 1}, // Top-left corner
		{4, 4, 5}, // Center
		{8, 8, 9}, // Bottom-right corner
		{2, 7, 3}, // Random position
		{6, 1, 7}, // Random position
	}

	for _, tc := range testCases {
		if !validator.IsValid(&board, tc.row, tc.col, tc.num) {
			t.Errorf("IsValid(%d, %d, %d) = false on empty board, expected true",
				tc.row, tc.col, tc.num)
		}
	}
}

// TestIsValid_RowConflict verifies that placing a duplicate number
// in the same row is correctly detected as invalid
func TestIsValid_RowConflict(t *testing.T) {
	board := utils.NewBoard()

	// Place 5 at position (0, 3)
	board[0][3] = 5

	// Attempt to place 5 in other positions in row 0
	testCases := []struct {
		col int
	}{
		{0}, {1}, {2}, {4}, {5}, {6}, {7}, {8},
	}

	for _, tc := range testCases {
		if validator.IsValid(&board, 0, tc.col, 5) {
			t.Errorf("IsValid(0, %d, 5) = true, expected false (row conflict at col 3)",
				tc.col)
		}
	}

	// Verify other numbers are still valid in row 0
	for num := 1; num <= 9; num++ {
		if num == 5 {
			continue
		}
		if !validator.IsValid(&board, 0, 0, num) {
			t.Errorf("IsValid(0, 0, %d) = false, expected true (no conflict)", num)
		}
	}
}

// TestIsValid_ColumnConflict verifies that placing a duplicate number
// in the same column is correctly detected as invalid
func TestIsValid_ColumnConflict(t *testing.T) {
	board := utils.NewBoard()

	// Place 7 at position (4, 2)
	board[4][2] = 7

	// Attempt to place 7 in other positions in column 2
	testCases := []struct {
		row int
	}{
		{0}, {1}, {2}, {3}, {5}, {6}, {7}, {8},
	}

	for _, tc := range testCases {
		if validator.IsValid(&board, tc.row, 2, 7) {
			t.Errorf("IsValid(%d, 2, 7) = true, expected false (column conflict at row 4)",
				tc.row)
		}
	}

	// Verify other numbers are still valid in column 2
	for num := 1; num <= 9; num++ {
		if num == 7 {
			continue
		}
		if !validator.IsValid(&board, 0, 2, num) {
			t.Errorf("IsValid(0, 2, %d) = false, expected true (no conflict)", num)
		}
	}
}

// TestIsValid_BoxConflict verifies that placing a duplicate number
// in the same 3x3 box is correctly detected as invalid
func TestIsValid_BoxConflict(t *testing.T) {
	board := utils.NewBoard()

	// Place 9 at position (1, 1) in box 0 (top-left)
	board[1][1] = 9

	// Attempt to place 9 in other positions within box 0 (rows 0-2, cols 0-2)
	testCases := []struct {
		row, col int
	}{
		{0, 0},
		{0, 1},
		{0, 2},
		{1, 0},
		{1, 2}, // Skip (1,1) as it's occupied
		{2, 0},
		{2, 1},
		{2, 2},
	}

	for _, tc := range testCases {
		if validator.IsValid(&board, tc.row, tc.col, 9) {
			t.Errorf("IsValid(%d, %d, 9) = true, expected false (box conflict at (1,1))",
				tc.row, tc.col)
		}
	}

	// Verify 9 is valid in different boxes
	if !validator.IsValid(&board, 0, 3, 9) {
		t.Errorf("IsValid(0, 3, 9) = false, expected true (box 1, different from box 0)")
	}

	if !validator.IsValid(&board, 3, 0, 9) {
		t.Errorf("IsValid(3, 0, 9) = false, expected true (box 3, different from box 0)")
	}
}

// TestIsValid_AllBoxes verifies box validation works correctly
// across all nine 3x3 boxes in the sudoku grid
func TestIsValid_AllBoxes(t *testing.T) {
	board := utils.NewBoard()

	// Place number 3 in the center of each of the 9 boxes
	boxPlacements := []struct {
		row, col int
		boxNum   int
	}{
		{1, 1, 0}, // Box 0 (top-left)
		{1, 4, 1}, // Box 1 (top-middle)
		{1, 7, 2}, // Box 2 (top-right)
		{4, 1, 3}, // Box 3 (middle-left)
		{4, 4, 4}, // Box 4 (center)
		{4, 7, 5}, // Box 5 (middle-right)
		{7, 1, 6}, // Box 6 (bottom-left)
		{7, 4, 7}, // Box 7 (bottom-middle)
		{7, 7, 8}, // Box 8 (bottom-right)
	}

	for _, placement := range boxPlacements {
		board[placement.row][placement.col] = 3
	}

	// Verify 3 is invalid in all other positions within each box
	for _, placement := range boxPlacements {
		boxRow := (placement.row / 3) * 3
		boxCol := (placement.col / 3) * 3

		for r := boxRow; r < boxRow+3; r++ {
			for c := boxCol; c < boxCol+3; c++ {
				// Skip the cell where 3 is already placed
				if r == placement.row && c == placement.col {
					continue
				}

				if validator.IsValid(&board, r, c, 3) {
					t.Errorf("IsValid(%d, %d, 3) = true in box %d, expected false",
						r, c, placement.boxNum)
				}
			}
		}
	}
}

// TestIsValid_MultipleConflicts verifies that the validator correctly
// detects when a number violates multiple rules simultaneously
func TestIsValid_MultipleConflicts(t *testing.T) {
	board := utils.NewBoard()

	// Place 4 at (0, 0)
	board[0][0] = 4

	// Test (0, 1): shares row 0 AND box 0 with (0, 0)
	if validator.IsValid(&board, 0, 1, 4) {
		t.Errorf("IsValid(0, 1, 4) = true, expected false (row + box conflict)")
	}

	// Test (1, 0): shares column 0 AND box 0 with (0, 0)
	if validator.IsValid(&board, 1, 0, 4) {
		t.Errorf("IsValid(1, 0, 4) = true, expected false (column + box conflict)")
	}

	// Place 4 at (0, 5) in a different box but same row
	board[0][5] = 4

	// Test (0, 3): shares row 0 with both (0, 0) and (0, 5)
	if validator.IsValid(&board, 0, 3, 4) {
		t.Errorf("IsValid(0, 3, 4) = true, expected false (row conflict)")
	}
}

// TestIsValid_PartiallyFilledBoard verifies validation on a realistic
// partially-filled sudoku puzzle from the exercise examples
func TestIsValid_PartiallyFilledBoard(t *testing.T) {
	// Use the example sudoku from the exercise
	board := utils.Board{
		{0, 9, 6, 0, 4, 0, 0, 0, 1},
		{1, 0, 0, 0, 6, 0, 0, 0, 4},
		{5, 0, 4, 8, 1, 0, 3, 9, 0},
		{0, 0, 7, 9, 5, 0, 0, 4, 3},
		{0, 3, 0, 0, 8, 0, 0, 0, 0},
		{4, 0, 5, 0, 2, 3, 0, 1, 8},
		{0, 1, 0, 6, 3, 0, 0, 5, 9},
		{0, 5, 9, 0, 7, 0, 8, 3, 0},
		{0, 0, 3, 5, 9, 0, 0, 0, 7},
	}

	testCases := []struct {
		row, col, num int
		expected      bool
		reason        string
	}{
		// Valid placements (empty cells with no conflicts)
		{0, 0, 3, true, "empty cell, no conflicts"},
		{1, 1, 7, true, "empty cell, no conflicts"},
		{4, 2, 1, true, "empty cell, no conflicts"},

		// Invalid placements - row conflicts
		{0, 0, 9, false, "9 already in row 0 at column 1"},
		{1, 1, 1, false, "1 already in row 1 at column 0"},

		// Invalid placements - column conflicts
		{0, 0, 5, false, "5 already in column 0 at row 2"},
		{1, 1, 9, false, "9 already in column 1 at row 0"},

		// Invalid placements - box conflicts
		{0, 0, 6, false, "6 already in box 0 at position (0,2)"},
		{1, 2, 9, false, "9 already in box 0 at position (0,1)"},
	}

	for _, tc := range testCases {
		result := validator.IsValid(&board, tc.row, tc.col, tc.num)
		if result != tc.expected {
			t.Errorf("IsValid(%d, %d, %d) = %v, expected %v (%s)",
				tc.row, tc.col, tc.num, result, tc.expected, tc.reason)
		}
	}
}

// TestIsValid_EdgeCases verifies validation at board boundaries
// and special positions
func TestIsValid_EdgeCases(t *testing.T) {
	board := utils.NewBoard()

	t.Run("All numbers 1-9 valid on empty board", func(t *testing.T) {
		// Every number should be placeable at (0,0) on empty board
		for num := 1; num <= 9; num++ {
			if !validator.IsValid(&board, 0, 0, num) {
				t.Errorf("IsValid(0, 0, %d) = false on empty board, expected true", num)
			}
		}
	})

	t.Run("Corner cells validation", func(t *testing.T) {
		// Test all four corners of the board
		corners := []struct {
			row, col int
			name     string
		}{
			{0, 0, "top-left"},
			{0, 8, "top-right"},
			{8, 0, "bottom-left"},
			{8, 8, "bottom-right"},
		}

		for _, corner := range corners {
			if !validator.IsValid(&board, corner.row, corner.col, 5) {
				t.Errorf("IsValid(%d, %d, 5) = false at %s corner, expected true",
					corner.row, corner.col, corner.name)
			}
		}
	})

	t.Run("Center cell validation", func(t *testing.T) {
		// Test the absolute center of the board (box 4, row 4, col 4)
		if !validator.IsValid(&board, 4, 4, 5) {
			t.Errorf("IsValid(4, 4, 5) = false at center, expected true")
		}
	})
}

// TestIsValid_FilledRow verifies that validation correctly handles
// a completely filled row
func TestIsValid_FilledRow(t *testing.T) {
	board := utils.NewBoard()

	// Fill row 5 with all numbers 1-9
	board[5] = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Every number should be invalid at every position in row 5
	for col := 0; col < 9; col++ {
		for num := 1; num <= 9; num++ {
			if validator.IsValid(&board, 5, col, num) {
				t.Errorf("IsValid(5, %d, %d) = true, expected false (row full)", col, num)
			}
		}
	}

	// Numbers should still be valid in other rows
	// Test row 0, column 3 with number 1
	// This avoids row 5 and column 0 (where 1 exists in row 5)
	if !validator.IsValid(&board, 0, 3, 1) {
		t.Errorf("IsValid(0, 3, 1) = false, expected true (different row, no conflicts)")
	}
}

// TestIsValid_FilledColumn verifies that validation correctly handles
// a completely filled column
func TestIsValid_FilledColumn(t *testing.T) {
	board := utils.NewBoard()

	// Fill column 5 with numbers 1-9
	for row := 0; row < 9; row++ {
		board[row][5] = row + 1
	}

	// Every number should be invalid at every position in column 5
	for row := 0; row < 9; row++ {
		for num := 1; num <= 9; num++ {
			if validator.IsValid(&board, row, 5, num) {
				t.Errorf("IsValid(%d, 5, %d) = true, expected false (column full)", row, num)
			}
		}
	}

	// Numbers should still be valid in other columns
	// Test row 3, column 0 with number 1
	// This avoids column 5 and row 0 (where 1 exists in column 5)
	if !validator.IsValid(&board, 3, 0, 1) {
		t.Errorf("IsValid(3, 0, 1) = false, expected true (different column, no conflicts)")
	}
}

// TestIsValid_FilledBox verifies that validation correctly handles
// a completely filled 3x3 box
func TestIsValid_FilledBox(t *testing.T) {
	board := utils.NewBoard()

	// Fill box 0 (top-left, rows 0-2, cols 0-2) with numbers 1-9
	board[0][0], board[0][1], board[0][2] = 1, 2, 3
	board[1][0], board[1][1], board[1][2] = 4, 5, 6
	board[2][0], board[2][1], board[2][2] = 7, 8, 9

	// Every number should be invalid at every position in box 0
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			for num := 1; num <= 9; num++ {
				if validator.IsValid(&board, r, c, num) {
					t.Errorf("IsValid(%d, %d, %d) = true, expected false (box full)", r, c, num)
				}
			}
		}
	}

	// Numbers should still be valid in other boxes
	// Test box 4 (center) at position (3, 3)
	if !validator.IsValid(&board, 3, 3, 1) {
		t.Errorf("IsValid(3, 3, 1) = false, expected true (different box)")
	}
}
