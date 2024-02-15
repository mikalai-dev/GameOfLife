package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

type Board [][]bool

func clearScreen() {
	cmd := exec.Command("clear") // For Linux/Unix
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func MakeBoard(rows, cols int) Board {
	board := make(Board, rows)
	for i := range board {
		board[i] = make([]bool, cols)
		for j := range board[i] {
			board[i][j] = rand.Intn(2) == 0
		}
	}
	return board
}

func (b Board) Print() {
	clearScreen()
	for _, row := range b {
		for _, cell := range row {
			if cell {
				fmt.Print("█ ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
	time.Sleep(time.Millisecond * 150)
}

// NextState calculates the next state of the board
func (b Board) NextState() Board {
	rows, cols := len(b), len(b[0])
	newBoard := MakeBoard(rows, cols)
	for i := range b {
		for j := range b[i] {
			neighbors := b.countNeighbours(i, j)
			if b[i][j] && (neighbors == 2 || neighbors == 3) {
				newBoard[i][j] = true
			} else if !b[i][j] && neighbors == 3 {
				newBoard[i][j] = true
			} else {
				newBoard[i][j] = false
			}
		}
	}
	return newBoard
}

func (b Board) countNeighbours(row, col int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			newRow, newCol := row+i, col+j
			if newRow >= 0 && newRow < len(b) && newCol >= 0 && newCol < len(b[0]) && b[newRow][newCol] {
				count++
			}
		}
	}
	return count
}

func (b Board) countCells() int {
	count := 0
	for _, row := range b {
		for _, value := range row {
			if value {
				count++
			}
		}
	}
	return count
}

func main() {
	rows, cols, steps := 20, 20, 100
	board := MakeBoard(rows, cols)

	initial_cells_count := board.countCells()

	for i := 0; i < steps; i++ {
		board.Print()
		board = board.NextState()
	}

	final_cells_count := board.countCells()

	fmt.Println("Cells initially: ", initial_cells_count)
	fmt.Println("Cells in the end: ", final_cells_count)
}
