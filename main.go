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
	AdjecentMines int
}

type Board struct {
	Rows  int
	Cols  int
	Cells [][]Cell
}
