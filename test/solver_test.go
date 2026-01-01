package test

import (
	"sudoku/solver"
	"sudoku/utils"
	"testing"
)

// TestSolve_EmptyBoard verifies that the solver can solve a completely empty board
func TestSolve_EmptyBoard(t *testing.T) {
	board := utils.NewBoard()

	// An empty board should be solvable (has many solutions)
	if !solver.Solve(&board) {
		t.Errorf("Solve() false on empty board, expected true")
	}

	// Verify the board is now complete (no zeros)
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == 0 {
				t.Errorf("Board still has a empty cell at (%d, %d) after solve", row, col)
			}

			// Verify each number is 1-9
			num := board[row][col]
			if num < 1 || num > 9 {
				t.Errorf("Board has invalid number at %d at (%d, %d)", num, row, col)
			}
		}
	}
}

// TestSolve_ExamplePuzzle verifies expected output of sample solution
func TestSolve_ExamplePuzzle(t *testing.T) {
	// Example puzzle from sample outputs
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

	// Expected solution from the exercise
	expected := utils.Board{
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

	// Solve the puzzle
	if !solver.Solve(&board) {
		t.Fatalf("Solve() = false, expected true (solvable puzzle)")
	}

	// Compare with expected solution
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] != expected[row][col] {
				t.Errorf("Board[%d][%d] = %d, expected = %d",
					row, col, board[row][col], expected[row][col])
			}
		}
	}
}

// TestSolve_AlreadySolved verifies that the solver correctly handles a board that is already complete
func TestSolve_AlreadySolved(t *testing.T) {
	// Already solved
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

	// Make a copy to compare later
	original := board

	// Should return true immediately (no empty cells)
	if !solver.Solve(&board) {
		t.Errorf("Solve () = false on already solved board, expected true")
	}

	// Board should remain unchanged
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] != original[row][col] {
				t.Errorf("Board was modified at (%d, %d), expected unchanged", row, col)
			}
		}
	}
}

// TestSolve_InvalidPuzzle verifies that the solver correctly returns false for an unsolvable puzzle (no valid solution exists)
func TestSolve_InvalidPuzzle(t *testing.T) {
	// A carefully constructed puzzle with no solution
	// This board has valid initial state but cannot be completed
	board := utils.Board{
		{5, 1, 6, 8, 4, 9, 7, 3, 2},
		{3, 0, 7, 6, 0, 5, 0, 0, 0},
		{8, 0, 9, 7, 0, 0, 0, 6, 5},
		{1, 3, 5, 0, 6, 0, 9, 0, 7},
		{4, 7, 2, 5, 9, 1, 0, 0, 6},
		{9, 6, 8, 3, 7, 0, 0, 5, 0},
		{2, 5, 3, 1, 8, 6, 0, 7, 4},
		{6, 8, 4, 2, 0, 7, 5, 0, 0},
		{7, 9, 1, 0, 5, 0, 6, 0, 8},
	}

	// Should return false (unsolvable)
	if solver.Solve(&board) {
		t.Errorf("Solve() = true on unsolvable puzzle, expected false")
	}
}

// TestSolve_VeryHardPuzzle verifies the solver can handle extremely difficult puzzles that require extensive backtracking
func TestSolve_VeryHardPuzzle(t *testing.T) {
	// One of the hardest known sudoku puzzles
	// Known as "Al Escargot" - created by Arto Inkala
	board := utils.Board{
		{1, 0, 0, 0, 0, 7, 0, 9, 0},
		{0, 3, 0, 0, 2, 0, 0, 0, 8},
		{0, 0, 9, 6, 0, 0, 5, 0, 0},
		{0, 0, 5, 3, 0, 0, 9, 0, 0},
		{0, 1, 0, 0, 8, 0, 0, 0, 2},
		{6, 0, 0, 0, 0, 4, 0, 0, 0},
		{3, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 4, 0, 0, 0, 0, 0, 0, 7},
		{0, 0, 7, 0, 0, 0, 3, 0, 0},
	}

	// This puzzle should still be solvable (but takes more time)
	if !solver.Solve(&board) {
		t.Errorf("Solve() = false on very hard puzzle, expected true")
	}

	// Verify solution is complete (no zeros remain)
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == 0 {
				t.Errorf("Board still has empty cell at (%d, %d)", row, col)
			}
		}
	}

	// Verify each number is in valid range
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			num := board[row][col]
			if num < 1 || num > 9 {
				t.Errorf("Board has invalid number %d at (%d, %d)", num, row, col)
			}
		}
	}
}
