package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func generatePathIter(n int) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)

		length := 2*n + 1
		state := make([]rune, length)

		for i := 0; i < n; i++ {
			state[i] = '>'
		}
		state[n] = '_'
		for i := n + 1; i < length; i++ {
			state[i] = '<'
		}

		target := strings.Repeat("<", n) + "_" + strings.Repeat(">", n)
		emptyIdx := n
		moveRight := true

		ch <- string(state)

		for string(state) != target {
			moved := false

			if emptyIdx > 1 && state[emptyIdx-2] == '>' && state[emptyIdx-1] == '<' {
				state[emptyIdx], state[emptyIdx-2] = state[emptyIdx-2], state[emptyIdx]
				emptyIdx -= 2
				moved = true
			} else if emptyIdx < length-2 && state[emptyIdx+2] == '<' && state[emptyIdx+1] == '>' {
				state[emptyIdx], state[emptyIdx+2] = state[emptyIdx+2], state[emptyIdx]
				emptyIdx += 2
				moved = true
			} else {
				if moveRight && emptyIdx > 0 && state[emptyIdx-1] == '>' {
					state[emptyIdx], state[emptyIdx-1] = state[emptyIdx-1], state[emptyIdx]
					emptyIdx--
					moveRight = false
					moved = true
				} else if !moveRight && emptyIdx < length-1 && state[emptyIdx+1] == '<' {
					state[emptyIdx], state[emptyIdx+1] = state[emptyIdx+1], state[emptyIdx]
					emptyIdx++
					moveRight = true
					moved = true
				} else if emptyIdx > 0 && state[emptyIdx-1] == '>' {
					state[emptyIdx], state[emptyIdx-1] = state[emptyIdx-1], state[emptyIdx]
					emptyIdx--
					moveRight = false
					moved = true
				} else if emptyIdx < length-1 && state[emptyIdx+1] == '<' {
					state[emptyIdx], state[emptyIdx+1] = state[emptyIdx+1], state[emptyIdx]
					emptyIdx++
					moveRight = true
					moved = true
				}
			}

			if !moved {
				break
			}

			ch <- string(state)
		}
	}()

	return ch
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil && err.Error() != "EOF" {
		return
	}

	input = strings.TrimSpace(input)
	if input == "" {
		return
	}

	n, err := strconv.Atoi(input)
	if err != nil {
		return
	}

	measureOnly := os.Getenv("FMI_TIME_ONLY") != ""

	start := time.Now()
	sequence := generatePathIter(n)

	if measureOnly {
		for range sequence {
		}
	} else {
		for state := range sequence {
			fmt.Println(state)
		}
	}

	elapsedMs := int(time.Since(start).Milliseconds())
	fmt.Printf("# TIMES_MS: alg=%d\n", elapsedMs)
}
