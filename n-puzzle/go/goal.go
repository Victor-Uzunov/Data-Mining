package main

// GoalBoard represents the goal state
type GoalBoard struct {
	goalBoard [][]int
	zeroX     int
	zeroY     int
}

// NewGoalBoard creates a new goal board
func NewGoalBoard(boardDim int) *GoalBoard {
	goalBoard := make([][]int, boardDim)
	for i := range goalBoard {
		goalBoard[i] = make([]int, boardDim)
	}
	return &GoalBoard{goalBoard: goalBoard}
}

// GenerateGoalBoard generates the goal board configuration
func (g *GoalBoard) GenerateGoalBoard(l int) {
	number := 1
	currentNumberCount := 0

	for i := 0; i < len(g.goalBoard); i++ {
		for j := 0; j < len(g.goalBoard[i]); j++ {
			if currentNumberCount == l {
				g.goalBoard[i][j] = 0
				g.zeroX = i
				g.zeroY = j
				currentNumberCount++
			} else {
				g.goalBoard[i][j] = number
				number++
				currentNumberCount++
			}
		}
	}

	if l == -1 {
		g.goalBoard[len(g.goalBoard)-1][len(g.goalBoard)-1] = 0
		g.zeroX = len(g.goalBoard) - 1
		g.zeroY = len(g.goalBoard) - 1
	}
}

// GetElementCoord returns the coordinates of a specific element
func (g *GoalBoard) GetElementCoord(element int) (int, int) {
	for i := 0; i < len(g.goalBoard); i++ {
		for j := 0; j < len(g.goalBoard[i]); j++ {
			if g.goalBoard[i][j] == element {
				return i, j
			}
		}
	}
	return -1, -1
}

// IsGoal checks if the given board matches the goal state
func (g *GoalBoard) IsGoal(board [][]int) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != g.goalBoard[i][j] {
				return false
			}
		}
	}
	return true
}

// GetBoardElement returns the element at position (i, j)
func (g *GoalBoard) GetBoardElement(i, j int) int {
	if i >= 0 && i < len(g.goalBoard) && j >= 0 && j < len(g.goalBoard[i]) {
		return g.goalBoard[i][j]
	}
	return -1
}
