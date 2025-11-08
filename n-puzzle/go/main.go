package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Board [][]int

func findZero(board Board, n int) (int, int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func calculateInversion(board Board, n int) int {
	var flatList []int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] != 0 {
				flatList = append(flatList, board[i][j])
			}
		}
	}

	inversions := 0
	for i := 0; i < len(flatList); i++ {
		for j := i + 1; j < len(flatList); j++ {
			if flatList[i] > flatList[j] {
				inversions++
			}
		}
	}
	return inversions
}

func isSolvable(initial Board, n int) bool {
	inversions := calculateInversion(initial, n)

	if n%2 != 0 {
		return inversions%2 == 0
	} else {
		r, _ := findZero(initial, n)
		rowFromBottom := n - r
		return (inversions+rowFromBottom)%2 != 0
	}
}

func serialize(board Board) string {
	var sb strings.Builder
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			sb.WriteString(strconv.Itoa(board[i][j]))
			sb.WriteString(",")
		}
		sb.WriteString(";")
	}
	return sb.String()
}

func manhattanDistance(initial Board, n int, goalPosMap map[int][2]int) int {
	distance := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			currPos := initial[i][j]
			if currPos == 0 {
				continue
			}
			goalPos := goalPosMap[currPos]
			goalI, goalJ := goalPos[0], goalPos[1]
			distance += abs(i-goalI) + abs(j-goalJ)
		}
	}
	return distance
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func swap(board Board, r, c, nr, nc int) Board {
	newBoard := make(Board, len(board))
	for i := range board {
		newBoard[i] = make([]int, len(board[i]))
		copy(newBoard[i], board[i])
	}
	newBoard[r][c], newBoard[nr][nc] = newBoard[nr][nc], newBoard[r][c]
	return newBoard
}

type SearchResult struct {
	path []string
	f    float64
}

func dfsIterative(board Board, path []string, g int, limit int, n int, target string, goalPosMap map[int][2]int, visited map[string]int) SearchResult {
	boardTuple := serialize(board)
	if prevG, exists := visited[boardTuple]; exists && prevG <= g {
		return SearchResult{nil, math.Inf(1)}
	}

	visited[boardTuple] = g

	h := manhattanDistance(board, n, goalPosMap)
	f := g + h

	if f > limit {
		return SearchResult{nil, float64(f)}
	}

	if boardTuple == target {
		return SearchResult{path, 0}
	}

	i, j := findZero(board, n)
	moves := [][3]interface{}{
		{"up", -1, 0},
		{"down", 1, 0},
		{"left", 0, -1},
		{"right", 0, 1},
	}
	minF := math.Inf(1)

	reserveMove := ""
	if len(path) > 0 {
		lastMove := path[len(path)-1]
		switch lastMove {
		case "up":
			reserveMove = "down"
		case "down":
			reserveMove = "up"
		case "left":
			reserveMove = "right"
		case "right":
			reserveMove = "left"
		}
	}

	for _, move := range moves {
		moveStr := move[0].(string)
		if moveStr == reserveMove {
			continue
		}

		di := move[1].(int)
		dj := move[2].(int)
		ni, nj := i+di, j+dj

		if 0 <= ni && ni < n && 0 <= nj && nj < n {
			newBoard := swap(board, i, j, ni, nj)
			newPath := make([]string, len(path)+1)
			copy(newPath, path)
			newPath[len(path)] = moveStr

			result := dfsIterative(newBoard, newPath, g+1, limit, n, target, goalPosMap, visited)

			if result.path != nil {
				return result
			}

			if result.f < minF {
				minF = result.f
			}
		}
	}

	return SearchResult{nil, minF}
}

func solve(initial Board, goal Board, n int) (interface{}, float64) {
	startTime := time.Now()

	if !isSolvable(initial, n) {
		ms := float64(time.Since(startTime).Milliseconds())
		return -1, ms
	}

	target := serialize(goal)
	goalPosMap := make(map[int][2]int)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			goalPosMap[goal[i][j]] = [2]int{i, j}
		}
	}

	limit := manhattanDistance(initial, n, goalPosMap)
	var solution interface{} = nil

	for solution == nil {
		visited := make(map[string]int)
		result := dfsIterative(initial, []string{}, 0, limit, n, target, goalPosMap, visited)

		if result.path != nil {
			solution = result.path
		} else if math.IsInf(result.f, 1) {
			solution = -1
			break
		} else {
			limit = int(result.f)
		}
	}

	elapsedMs := float64(time.Since(startTime).Milliseconds())
	return solution, elapsedMs
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	i, _ := strconv.Atoi(scanner.Text())

	tilesNum := int(math.Sqrt(float64(n + 1)))

	goalBoard := make([]int, 0, n+1)
	for k := 1; k <= n; k++ {
		goalBoard = append(goalBoard, k)
	}
	goalBoard = append(goalBoard, 0)

	targetZeroIdx := i
	if i == -1 {
		targetZeroIdx = n
	}

	if 0 <= targetZeroIdx && targetZeroIdx < len(goalBoard) {
		goalBoard[len(goalBoard)-1], goalBoard[targetZeroIdx] = goalBoard[targetZeroIdx], goalBoard[len(goalBoard)-1]
	}

	goal := make(Board, tilesNum)
	for k := 0; k < tilesNum; k++ {
		goal[k] = goalBoard[k*tilesNum : (k+1)*tilesNum]
	}

	var initial []int
	for k := 0; k < tilesNum; k++ {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 0 {
			break
		}
		for _, part := range parts {
			val, _ := strconv.Atoi(part)
			initial = append(initial, val)
		}
	}

	initialBoard := make(Board, tilesNum)
	for k := 0; k < tilesNum; k++ {
		initialBoard[k] = make([]int, tilesNum)
		for l := 0; l < tilesNum; l++ {
			initialBoard[k][l] = initial[k*tilesNum+l]
		}
	}

	solution, _ := solve(initialBoard, goal, tilesNum)

	switch v := solution.(type) {
	case int:
		if v == -1 {
			fmt.Println("-1")
		}
	case []string:
		fmt.Println(len(v))
		for _, step := range v {
			switch step {
			case "right":
				fmt.Println("left")
			case "left":
				fmt.Println("right")
			case "up":
				fmt.Println("down")
			case "down":
				fmt.Println("up")
			}
		}
	default:
		fmt.Println("-1")
	}
}
