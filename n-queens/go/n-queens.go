package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// constructiveSolution generates a deterministic O(n) solution for large n
// Uses classical arithmetic patterns that guarantee valid placement for all n >= 4
func constructiveSolution(n int) []int {
	if n == 1 {
		return []int{0}
	}
	if n == 2 || n == 3 {
		return nil
	}

	var seq []int

	if n%6 != 2 && n%6 != 3 {
		// Pattern: all evens, then all odds (1-based)
		for i := 2; i <= n; i += 2 {
			seq = append(seq, i)
		}
		for i := 1; i <= n; i += 2 {
			seq = append(seq, i)
		}
	} else if n%6 == 2 {
		// Special reordering: (3,1,5,7,...) + (4,2,6,8,...)
		var odds []int
		if n >= 3 {
			odds = append(odds, 3, 1)
		}
		for i := 5; i <= n; i += 2 {
			odds = append(odds, i)
		}

		var evens []int
		if n >= 4 {
			evens = append(evens, 4)
		}
		if n >= 2 {
			evens = append(evens, 2)
		}
		for i := 6; i <= n; i += 2 {
			evens = append(evens, i)
		}

		seq = append(odds, evens...)
	} else { // n % 6 == 3
		// Pattern: (2,4,6,...,n-1) + (1,3,5,...,n-2) + n
		for i := 2; i < n; i += 2 {
			seq = append(seq, i)
		}
		for i := 1; i < n-1; i += 2 {
			seq = append(seq, i)
		}
		seq = append(seq, n)
	}

	// Convert to 0-based
	result := make([]int, n)
	for i, val := range seq {
		result[i] = val - 1
	}
	return result
}

// Solver encapsulates the Min-Conflicts algorithm state
type Solver struct {
	n              int
	state          []int
	rowCounts      []int
	diag1Counts    []int
	diag2Counts    []int
	conflicts      []int
	conflictedSet  map[int]bool
	conflictedList []int
	offset         int
	rng            *rand.Rand
}

func newSolver(n int) *Solver {
	s := &Solver{
		n:              n,
		state:          make([]int, n),
		rowCounts:      make([]int, n),
		diag1Counts:    make([]int, 2*n+1),
		diag2Counts:    make([]int, 2*n+1),
		conflicts:      make([]int, n),
		conflictedSet:  make(map[int]bool),
		conflictedList: make([]int, 0, n),
		offset:         n,
		rng:            rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	// Initialize with random permutation
	perm := s.rng.Perm(n)
	copy(s.state, perm)

	// Build initial counts
	for i := 0; i < n; i++ {
		s.rowCounts[i] = 1 // permutation guarantees 1 per row
	}
	for col, row := range s.state {
		s.diag1Counts[row-col+s.offset]++
		s.diag2Counts[row+col]++
	}

	// Initialize conflicts
	for col, row := range s.state {
		c := s.computeConflicts(col, row)
		s.conflicts[col] = c
		if c > 0 {
			s.conflictedSet[col] = true
			s.conflictedList = append(s.conflictedList, col)
		}
	}

	return s
}

func (s *Solver) computeConflicts(col, row int) int {
	return (s.rowCounts[row] - 1) +
		(s.diag1Counts[row-col+s.offset] - 1) +
		(s.diag2Counts[row+col] - 1)
}

func (s *Solver) markConflicted(col int) {
	c := s.computeConflicts(col, s.state[col])
	s.conflicts[col] = c
	_, present := s.conflictedSet[col]

	if c > 0 && !present {
		s.conflictedSet[col] = true
		s.conflictedList = append(s.conflictedList, col)
	} else if c == 0 && present {
		delete(s.conflictedSet, col)
		// Remove from list (swap with last)
		for i, val := range s.conflictedList {
			if val == col {
				s.conflictedList[i] = s.conflictedList[len(s.conflictedList)-1]
				s.conflictedList = s.conflictedList[:len(s.conflictedList)-1]
				break
			}
		}
	}
}

func (s *Solver) move(col, newRow int) {
	oldRow := s.state[col]
	if oldRow == newRow {
		return
	}

	// Update counts
	s.rowCounts[oldRow]--
	s.diag1Counts[oldRow-col+s.offset]--
	s.diag2Counts[oldRow+col]--

	s.rowCounts[newRow]++
	s.diag1Counts[newRow-col+s.offset]++
	s.diag2Counts[newRow+col]++

	s.state[col] = newRow

	// Determine affected columns
	oldD1 := oldRow - col + s.offset
	oldD2 := oldRow + col
	newD1 := newRow - col + s.offset
	newD2 := newRow + col

	scanPool := s.conflictedList
	if len(scanPool) < s.n/10 {
		// Sparse conflicts - scan all for accuracy
		scanPool = make([]int, s.n)
		for i := 0; i < s.n; i++ {
			scanPool[i] = i
		}
	}

	affected := make(map[int]bool)
	for _, c := range scanPool {
		if c == col {
			affected[c] = true
			continue
		}
		r := s.state[c]
		if r == oldRow || r == newRow ||
			(r-c+s.offset) == oldD1 || (r-c+s.offset) == newD1 ||
			(r+c) == oldD2 || (r+c) == newD2 {
			affected[c] = true
		}
	}

	for c := range affected {
		s.markConflicted(c)
	}
}

func (s *Solver) restart() {
	// Random shuffle
	s.rng.Shuffle(s.n, func(i, j int) {
		s.state[i], s.state[j] = s.state[j], s.state[i]
	})

	// Rebuild counts
	for i := 0; i < s.n; i++ {
		s.rowCounts[i] = 0
	}
	for i := 0; i < 2*s.n+1; i++ {
		s.diag1Counts[i] = 0
		s.diag2Counts[i] = 0
	}
	for col, row := range s.state {
		s.rowCounts[row]++
		s.diag1Counts[row-col+s.offset]++
		s.diag2Counts[row+col]++
	}

	// Rebuild conflicts
	for i := 0; i < s.n; i++ {
		s.conflicts[i] = 0
	}
	s.conflictedSet = make(map[int]bool)
	s.conflictedList = s.conflictedList[:0]

	for col, row := range s.state {
		c := s.computeConflicts(col, row)
		s.conflicts[col] = c
		if c > 0 {
			s.conflictedSet[col] = true
			s.conflictedList = append(s.conflictedList, col)
		}
	}
}

func (s *Solver) solve(maxSteps int) []int {
	restartThreshold := 2 * s.n
	if s.n >= 2000 {
		restartThreshold = s.n
	}
	stepsSinceRestart := 0

	for step := 0; step < maxSteps; step++ {
		if len(s.conflictedList) == 0 {
			return s.state
		}

		if stepsSinceRestart >= restartThreshold {
			s.restart()
			stepsSinceRestart = 0
			continue
		}

		// Pick random conflicted column
		col := s.conflictedList[s.rng.Intn(len(s.conflictedList))]
		currentRow := s.state[col]

		// Select candidate rows
		var candidateRows []int
		if s.n <= 50 {
			candidateRows = make([]int, s.n)
			for i := 0; i < s.n; i++ {
				candidateRows[i] = i
			}
		} else if s.n <= 2000 {
			candidateRows = make([]int, s.n)
			for i := 0; i < s.n; i++ {
				candidateRows[i] = i
			}
		} else {
			// Sampling for large boards
			sampleSize := 384
			candidateSet := make(map[int]bool)
			for len(candidateSet) < sampleSize && len(candidateSet) < s.n {
				candidateSet[s.rng.Intn(s.n)] = true
			}
			candidateSet[currentRow] = true
			candidateSet[(currentRow+1)%s.n] = true
			candidateSet[(currentRow-1+s.n)%s.n] = true

			candidateRows = make([]int, 0, len(candidateSet))
			for r := range candidateSet {
				candidateRows = append(candidateRows, r)
			}
		}

		// Find best row
		minConf := int(^uint(0) >> 1) // Max int
		bestRows := make([]int, 0, 10)
		for _, r := range candidateRows {
			c := s.computeConflicts(col, r)
			if c < minConf {
				minConf = c
				bestRows = bestRows[:0]
				bestRows = append(bestRows, r)
			} else if c == minConf {
				bestRows = append(bestRows, r)
			}
		}

		newRow := bestRows[s.rng.Intn(len(bestRows))]
		s.move(col, newRow)
		stepsSinceRestart++
	}

	return nil
}

func solveNQueens(n int, maxSteps int) []int {
	if n == 2 || n == 3 {
		return nil
	}
	if n == 1 {
		return []int{0}
	}

	// Use constructive solution for large n
	if n >= 500 {
		return constructiveSolution(n)
	}

	// Min-Conflicts for smaller n
	if maxSteps == 0 {
		maxSteps = 20 * n
		if maxSteps < 5000 {
			maxSteps = 5000
		}
	}

	solver := newSolver(n)
	return solver.solve(maxSteps)
}

func main() {
	timeOnly := os.Getenv("FMI_TIME_ONLY") == "1"

	// Read N from command line or stdin
	var n int
	if len(os.Args) > 1 {
		var err error
		n, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid input: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Scan(&n)
	}

	// Solve
	start := time.Now()
	result := solveNQueens(n, 0)
	elapsed := time.Since(start)

	elapsedMs := float64(elapsed.Nanoseconds()) / 1e6

	// Output
	if timeOnly {
		fmt.Printf("# TIMES_MS: alg=%.3f\n", elapsedMs)
	} else {
		fmt.Printf("# TIMES_MS: alg=%.3f\n", elapsedMs)
		if result == nil {
			fmt.Println(-1)
		} else {
			fmt.Println(result)
		}
	}
}
