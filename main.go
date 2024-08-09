package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Cell struct {
	IsMine        bool
	IsRevealed    bool
	IsFlagged     bool
	AdjacentMines int
}

type Board struct {
	Rows  int
	Cols  int
	Cells [][]Cell
}

func NewBoard(rows, cols, mines int) *Board {
	board := &Board{
		Rows:  rows,
		Cols:  cols,
		Cells: make([][]Cell, rows),
	}

	for i := range board.Cells {
		board.Cells[i] = make([]Cell, cols)
	}

	placeMines(board, mines)
	calculateAdjacentMines(board)

	return board
}

func placeMines(board *Board, mines int) {
	rand.Seed(time.Now().UnixNano())
	for mines > 0 {
		r := rand.Intn(board.Rows)
		c := rand.Intn(board.Cols)
		if !board.Cells[r][c].IsMine {
			board.Cells[r][c].IsMine = true
			mines--
		}
	}
}

func calculateAdjacentMines(board *Board) {
	directions := []struct{ dr, dc int }{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for r := 0; r < board.Rows; r++ {
		for c := 0; c < board.Cols; c++ {
			if board.Cells[r][c].IsMine {
				continue
			}

			for _, d := range directions {
				nr, nc := r+d.dr, c+d.dc
				if nr >= 0 && nr < board.Rows && nc >= 0 && nc < board.Cols && board.Cells[nr][nc].IsMine {
					board.Cells[r][c].AdjacentMines++
				}
			}
		}
	}
}

func (b *Board) Reveal(r, c int) bool {
	if r < 0 || r >= b.Rows || c < 0 || c >= b.Cols || b.Cells[r][c].IsRevealed {
		return false
	}

	b.Cells[r][c].IsRevealed = true

	if b.Cells[r][c].IsMine {
		endGame(b)
		return true
	}

	if b.Cells[r][c].AdjacentMines == 0 {
		directions := []struct{ dr, dc int }{
			{-1, -1}, {-1, 0}, {-1, 1},
			{0, -1}, {0, 1},
			{1, -1}, {1, 0}, {1, 1},
		}

		for _, d := range directions {
			b.Reveal(r+d.dr, c+d.dc)
		}
	}

	return false
}

func printBoard(board *Board) {
	for _, row := range board.Cells {
		for _, cell := range row {
			if cell.IsRevealed {
				if cell.IsMine {
					fmt.Print("* ")
				} else {
					fmt.Printf("%d ", cell.AdjacentMines)
				}
			} else if cell.IsFlagged {
				fmt.Print("F ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}

func checkWin(board *Board) bool {
	for _, row := range board.Cells {
		for _, cell := range row {
			if !cell.IsMine && !cell.IsRevealed {
				return false
			}
		}
	}
	return true
}

func endGame(board *Board) {
	fmt.Println("Game Over! You hit a mine.")
	for r := 0; r < board.Rows; r++ {
		for c := 0; c < board.Cols; c++ {
			board.Cells[r][c].IsRevealed = true
		}
	}
	printBoard(board)
}

func main() {
	rows, cols, mines := 10, 10, 10
	board := NewBoard(rows, cols, mines)

	var r, c int
	for {
		printBoard(board)
		fmt.Print("Enter row and column to reveal (e.g., '3 4'): ")
		fmt.Scan(&r, &c)

		if board.Reveal(r, c) {
			break
		}

		// Check for win condition
		if checkWin(board) {
			fmt.Println("Congratulations! You've cleared the minefield.")
			printBoard(board)
			break
		}
	}
}
