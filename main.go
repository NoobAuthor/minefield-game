package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Create a Struct for the Game Board
// Define a struct to represent the ame board, cells, and their states.

type Cell struct {
	IsMine        bool
	IsReavealed   bool
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
