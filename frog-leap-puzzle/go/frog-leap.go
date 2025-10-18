package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func solveDFS(N int) []string {
	L := 2*N + 1
	start := make([]byte, L)
	goal := make([]byte, L)

	for i := 0; i < N; i++ {
		start[i] = '>'
		goal[i] = '<'
	}
	start[N] = '_'
	goal[N] = '_'
	for i := N + 1; i < L; i++ {
		start[i] = '<'
		goal[i] = '>'
	}

	startKey := string(start)
	goalKey := string(goal)

	visited := make(map[string]bool)
	var path []string
	found := false

	var dfs func(state []byte) bool
	dfs = func(state []byte) bool {
		if found {
			return true
		}
		key := string(state)
		if key == goalKey {
			path = append(path, key)
			found = true
			return true
		}
		visited[key] = true

		for i := 0; i < L; i++ {
			if state[i] == '>' {
				if i+2 < L && state[i+1] != '_' && state[i+2] == '_' {
					newState := make([]byte, L)
					copy(newState, state)
					newState[i], newState[i+2] = '_', '>'
					k := string(newState)
					if !visited[k] {
						if dfs(newState) {
							path = append(path, key)
							return true
						}
					}
				}
				if i+1 < L && state[i+1] == '_' {
					newState := make([]byte, L)
					copy(newState, state)
					newState[i], newState[i+1] = '_', '>'
					k := string(newState)
					if !visited[k] {
						if dfs(newState) {
							path = append(path, key)
							return true
						}
					}
				}
			}
		}

		for i := L - 1; i >= 0; i-- {
			if state[i] == '<' {
				if i-2 >= 0 && state[i-1] != '_' && state[i-2] == '_' {
					newState := make([]byte, L)
					copy(newState, state)
					newState[i], newState[i-2] = '_', '<'
					k := string(newState)
					if !visited[k] {
						if dfs(newState) {
							path = append(path, key)
							return true
						}
					}
				}
				if i-1 >= 0 && state[i-1] == '_' {
					newState := make([]byte, L)
					copy(newState, state)
					newState[i], newState[i-1] = '_', '<'
					k := string(newState)
					if !visited[k] {
						if dfs(newState) {
							path = append(path, key)
							return true
						}
					}
				}
			}
		}

		return false
	}

	dfs([]byte(startKey))
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}

func main() {
	timeOnly := os.Getenv("FMI_TIME_ONLY") == "1"

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		_, err := fmt.Fprintln(os.Stderr, "Error reading input")
		if err != nil {
			panic(err)
		}
		os.Exit(1)
	}

	N, err := strconv.Atoi(scanner.Text())
	if err != nil {
		_, err = fmt.Fprintln(os.Stderr, "Invalid input: not an integer")
		if err != nil {
			panic(err)
		}
		os.Exit(1)
	}

	start := time.Now()
	solution := solveDFS(N)
	duration := time.Since(start)

	if timeOnly {
		fmt.Printf("# TIMES_MS: alg=%d\n", duration.Nanoseconds()/1000000)
	} else {
		for _, state := range solution {
			fmt.Println(state)
		}
	}
}
