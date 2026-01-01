package main

import (
	"fmt"
	"os"
	"sudoku/parser"
	"sudoku/solver"
	"sudoku/utils"
)

func main() {
	// Get command-line arguments
	args := os.Args[1:]

	// Parse the arguments into a board
	board, err := parser.ParseArgs(args)
	if err != nil {
		fmt.Println("Error")
		return
	}

	// Attempt to solve sudoku
	if !solver.Solve(&board) {
		fmt.Println("Error")
		return
	}

	// Print the solved board
	utils.PrintBoard(&board)
}
