package main

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
		for j := 0; j < len(b.board[i]); j++ {
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
	if i >= 0 && i < len(b.board) && j >= 0 && j < len(b.board[i]) {
		return b.board[i][j]
	}
	return -1
}
