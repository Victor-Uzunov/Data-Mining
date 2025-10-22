package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func generatePathIter(N int) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)

		s := make([]rune, 2*N+1)
		for i := 0; i < N; i++ {
			s[i] = '>'
		}
		s[N] = '_'
		for i := N + 1; i < 2*N+1; i++ {
			s[i] = '<'
		}

		target := strings.Repeat("<", N) + "_" + strings.Repeat(">", N)
		L := 2*N + 1
		blank := N
		moveBlankRight := true

		ch <- string(s)

		for string(s) != target {
			moved := false

			if blank > 1 && s[blank-2] == '>' && s[blank-1] == '<' {
				s[blank], s[blank-2] = s[blank-2], s[blank]
				blank -= 2
				moved = true
			} else if blank < L-2 && s[blank+2] == '<' && s[blank+1] == '>' {
				s[blank], s[blank+2] = s[blank+2], s[blank]
				blank += 2
				moved = true
			} else {
				if moveBlankRight && blank > 0 && s[blank-1] == '>' {
					s[blank], s[blank-1] = s[blank-1], s[blank]
					blank -= 1
					moved = true
					moveBlankRight = false
				} else if !moveBlankRight && blank < L-1 && s[blank+1] == '<' {
					s[blank], s[blank+1] = s[blank+1], s[blank]
					blank += 1
					moved = true
					moveBlankRight = true
				} else if blank > 0 && s[blank-1] == '>' {
					s[blank], s[blank-1] = s[blank-1], s[blank]
					blank -= 1
					moved = true
					moveBlankRight = false
				} else if blank < L-1 && s[blank+1] == '<' {
					s[blank], s[blank+1] = s[blank+1], s[blank]
					blank += 1
					moved = true
					moveBlankRight = true
				}
			}

			if !moved {
				break
			}

			ch <- string(s)
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

	N, err := strconv.Atoi(input)
	if err != nil {
		return
	}

	timeOnly := os.Getenv("FMI_TIME_ONLY") != ""

	t0 := time.Now()

	gen := generatePathIter(N)

	if timeOnly {
		for range gen {

		}
	} else {
		for state := range gen {
			fmt.Println(state)
		}
	}

	tMs := int(time.Since(t0).Milliseconds())
	fmt.Printf("# TIMES_MS: alg=%d\n", tMs)
}
