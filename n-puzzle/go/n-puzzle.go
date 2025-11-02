package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

// Direction represents the movement direction
type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
)

func (d Direction) String() string {
	return [...]string{"left", "right", "up", "down"}[d]
}

// Board represents the puzzle board state
type Board struct {
	board       [][]int
	parentBoard *Board
	direction   Direction
}

// NewBoard creates a new board
func NewBoard(initialBoard [][]int) *Board {
	return &Board{
		board: initialBoard,
	}
}

// GetZeroCoord returns the coordinates of the zero tile
func (b *Board) GetZeroCoord() (int, int) {
	for i := 0; i < len(b.board); i++ {
		for j := 0; j < len(b.board); j++ {
			if b.board[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

// copyBoard creates a deep copy of the board
func copyBoard(src [][]int) [][]int {
	dst := make([][]int, len(src))
	for i := range src {
		dst[i] = make([]int, len(src[i]))
		copy(dst[i], src[i])
	}
	return dst
}

// GenerateLeftChildren generates the left move (moving zero right, tile moves left)
func (b *Board) GenerateLeftChildren() *Board {
	zeroX, zeroY := b.GetZeroCoord()
	leftChildren := copyBoard(b.board)

	// left means moving the tile from right to left (zero moves right)
	if zeroY+1 < len(b.board) {
		leftChildren[zeroX][zeroY] = leftChildren[zeroX][zeroY+1]
		leftChildren[zeroX][zeroY+1] = 0

		child := NewBoard(leftChildren)
		child.parentBoard = b
		child.direction = Left
		return child
	}
	return nil
}

// GenerateRightChildren generates the right move (moving zero left, tile moves right)
func (b *Board) GenerateRightChildren() *Board {
	zeroX, zeroY := b.GetZeroCoord()
	rightChildren := copyBoard(b.board)

	// right means moving the tile from left to right (zero moves left)
	if zeroY-1 >= 0 {
		rightChildren[zeroX][zeroY] = rightChildren[zeroX][zeroY-1]
		rightChildren[zeroX][zeroY-1] = 0

		child := NewBoard(rightChildren)
		child.parentBoard = b
		child.direction = Right
		return child
	}
	return nil
}

// GenerateTopChildren generates the up move (moving zero down, tile moves up)
func (b *Board) GenerateTopChildren() *Board {
	zeroX, zeroY := b.GetZeroCoord()
	topChildren := copyBoard(b.board)

	// up means moving the tile from bottom to top (zero moves down)
	if zeroX+1 < len(b.board) {
		topChildren[zeroX][zeroY] = topChildren[zeroX+1][zeroY]
		topChildren[zeroX+1][zeroY] = 0

		child := NewBoard(topChildren)
		child.parentBoard = b
		child.direction = Up
		return child
	}
	return nil
}

// GenerateBottomChildren generates the down move (moving zero up, tile moves down)
func (b *Board) GenerateBottomChildren() *Board {
	zeroX, zeroY := b.GetZeroCoord()
	bottomChildren := copyBoard(b.board)

	// down means moving the tile from top to bottom (zero moves up)
	if zeroX-1 >= 0 {
		bottomChildren[zeroX][zeroY] = bottomChildren[zeroX-1][zeroY]
		bottomChildren[zeroX-1][zeroY] = 0

		child := NewBoard(bottomChildren)
		child.parentBoard = b
		child.direction = Down
		return child
	}
	return nil
}

// GenerateAllChildren generates all possible children boards
func (b *Board) GenerateAllChildren() []*Board {
	var children []*Board

	if left := b.GenerateLeftChildren(); left != nil {
		children = append(children, left)
	}
	if right := b.GenerateRightChildren(); right != nil {
		children = append(children, right)
	}
	if top := b.GenerateTopChildren(); top != nil {
		children = append(children, top)
	}
	if bottom := b.GenerateBottomChildren(); bottom != nil {
		children = append(children, bottom)
	}

	return children
}

// GetBoardDim returns the dimension of the board
func (b *Board) GetBoardDim() int {
	return len(b.board)
}

// GetBoardElement returns the element at position (i, j)
func (b *Board) GetBoardElement(i, j int) int {
	if i >= 0 && i < len(b.board) && j >= 0 && j < len(b.board) {
		return b.board[i][j]
	}
	return -1
}

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
		for j := 0; j < len(g.goalBoard); j++ {
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
		for j := 0; j < len(g.goalBoard); j++ {
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
		for j := 0; j < len(board); j++ {
			if board[i][j] != g.goalBoard[i][j] {
				return false
			}
		}
	}
	return true
}

// GetBoardElement returns the element at position (i, j)
func (g *GoalBoard) GetBoardElement(i, j int) int {
	if i >= 0 && i < len(g.goalBoard) && j >= 0 && j < len(g.goalBoard) {
		return g.goalBoard[i][j]
	}
	return -1
}

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
		for j := 0; j < len(board); j++ {
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
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

	measureOnly := os.Getenv("FMI_TIME_ONLY") != ""
	start := time.Now()

	startBoard := NewBoard(boardArr)
	goalBoard := NewGoalBoard(boardDim)
	goalBoard.GenerateGoalBoard(l)

	solver := NewSolver(startBoard, goalBoard)

	if !solver.IsSolvable(N) {
		fmt.Println("-1")
		return
	}

	solver.IDAStar(startBoard)

	fmt.Println(len(solver.steps))
	for i := len(solver.steps) - 1; i >= 0; i-- {
		if measureOnly {
			elapsedMs := int(time.Since(start).Milliseconds())
			fmt.Printf("# TIMES_MS: alg=%d\n", elapsedMs)
		}
		fmt.Println(solver.steps[i])
	}
}
