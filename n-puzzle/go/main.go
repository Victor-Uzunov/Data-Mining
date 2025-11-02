package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

// readInput reads the puzzle input from stdin
func readInput() (int, int, [][]int) {
	var N, l int
	fmt.Scan(&N)
	fmt.Scan(&l)

	boardDim := int(math.Sqrt(float64(N + 1)))
	boardArr := make([][]int, boardDim)
	for i := range boardArr {
		boardArr[i] = make([]int, boardDim)
	}

	for i := 0; i < boardDim; i++ {
		for j := 0; j < boardDim; j++ {
			fmt.Scan(&boardArr[i][j])
		}
	}

	return N, l, boardArr
}

// printSolution prints the solution steps
func printSolution(steps []Direction, measureTime bool, start time.Time) {
	fmt.Println(len(steps))
	for i := len(steps) - 1; i >= 0; i-- {
		if measureTime {
			elapsedMs := int(time.Since(start).Milliseconds())
			fmt.Printf("# TIMES_MS: alg=%d\n", elapsedMs)
		}
		fmt.Println(steps[i])
	}
}

func main() {
	measureOnly := os.Getenv("FMI_TIME_ONLY") != ""
	start := time.Now()

	// Read input
	N, l, boardArr := readInput()

	// Initialize puzzle components
	startBoard := NewBoard(boardArr)
	goalBoard := NewGoalBoard(startBoard.GetBoardDim())
	goalBoard.GenerateGoalBoard(l)

	// Create solver and check solvability
	solver := NewSolver(startBoard, goalBoard)
	if !solver.IsSolvable(N) {
		fmt.Println("-1")
		return
	}

	// Solve the puzzle
	solver.IDAStar(startBoard)

	// Print solution
	printSolution(solver.steps, measureOnly, start)
}
