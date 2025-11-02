package main

import "math"

// Solver implements the IDA* algorithm
type Solver struct {
	board     *Board
	goalBoard *GoalBoard
	steps     []Direction
}

// NewSolver creates a new solver
func NewSolver(board *Board, goalBoard *GoalBoard) *Solver {
	return &Solver{
		board:     board,
		goalBoard: goalBoard,
		steps:     []Direction{},
	}
}

// GetHeuristic calculates the Manhattan distance heuristic
func (s *Solver) GetHeuristic(board [][]int) int {
	manhattanSum := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != s.goalBoard.GetBoardElement(i, j) && board[i][j] != 0 {
				goalX, goalY := s.goalBoard.GetElementCoord(board[i][j])
				manhattanSum += abs(i-goalX) + abs(j-goalY)
			}
		}
	}
	return manhattanSum
}

// Search performs the recursive search
func (s *Solver) Search(board *Board, g, threshold int) int {
	evalFunc := g + s.GetHeuristic(board.board)
	if evalFunc > threshold {
		return evalFunc
	}
	if s.goalBoard.IsGoal(board.board) {
		return -1
	}

	min := math.MaxInt32

	for _, nextState := range board.GenerateAllChildren() {
		temp := s.Search(nextState, g+1, threshold)
		if temp == -1 {
			s.steps = append(s.steps, nextState.direction)
			return -1
		}
		if temp < min {
			min = temp
		}
	}

	return min
}

// IDAStar runs the IDA* algorithm
func (s *Solver) IDAStar(startBoard *Board) {
	threshold := s.GetHeuristic(startBoard.board)

	for {
		temp := s.Search(startBoard, 0, threshold)
		if temp == -1 {
			return
		}
		threshold = temp
	}
}

// IsSolvable checks if the puzzle is solvable
func (s *Solver) IsSolvable(N int) bool {
	boardDim := s.board.GetBoardDim()
	arr := make([]int, N+1)
	idx := 0
	zeroRow := 0

	for i := 0; i < boardDim; i++ {
		for j := 0; j < boardDim; j++ {
			currentElement := s.board.GetBoardElement(i, j)
			arr[idx] = currentElement
			idx++
			if currentElement == 0 {
				zeroRow = i
			}
		}
	}

	numberOfInversions := 0
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] != 0 && arr[j] != 0 && arr[i] > arr[j] {
				numberOfInversions++
			}
		}
	}

	if boardDim%2 != 0 {
		return numberOfInversions%2 == 0
	}
	return (numberOfInversions+zeroRow)%2 != 0
}
