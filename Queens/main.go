package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solveNQueens(n int) [][]string {
	var results [][]string // final output

	// make board
	board := make([][]rune, n)
	for i := range board {
		board[i] = make([]rune, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	// recursive anonymous function
	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			var solution []string
			for _, r := range board {
				solution = append(solution, string(r))
			}
			results = append(results, solution)
			return
		}
		for col := 0; col < n; col++ {
			if isSafe(board, row, col, n) {
				board[row][col] = 'Q'
				backtrack(row + 1)
				board[row][col] = '.' // backtrack
			}
		}
	}
	backtrack(0)
	return results
}

func isSafe(board [][]rune, row, col, n int) bool {
	// Check the column
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	// Check the upper-left diagonal
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	// Check the upper-right diagonal
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	return true
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <n>")
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		fmt.Println("Invalid input. Please enter a positive integer.")
		os.Exit(1)
	}

	solutions := solveNQueens(n)
	for _, solution := range solutions {
		fmt.Println(strings.Join(solution, "\n"))
		fmt.Println()
	}
}
