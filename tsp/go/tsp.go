package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	x, y float64
}

func distance(p1, p2 Point) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	return math.Sqrt(dx*dx + dy*dy)
}

func pathLength(route []int, dist [][]float64) float64 {
	sum := 0.0
	for i := 0; i < len(route)-1; i++ {
		sum += dist[route[i]][route[i+1]]
	}
	return sum
}

func initPopulation(popSize, n int) [][]int {
	population := make([][]int, popSize)
	base := make([]int, n)
	for i := 0; i < n; i++ {
		base[i] = i
	}

	for i := 0; i < popSize; i++ {
		perm := make([]int, n)
		copy(perm, base)
		rand.Shuffle(n, func(i, j int) { perm[i], perm[j] = perm[j], perm[i] })
		population[i] = perm
	}
	return population
}

type Scored struct {
	score float64
	ind   []int
}

func evaluate(pop [][]int, dist [][]float64) []Scored {
	scored := make([]Scored, len(pop))
	for i, ind := range pop {
		scored[i] = Scored{
			score: pathLength(ind, dist),
			ind:   ind,
		}
	}
	sort.Slice(scored, func(i, j int) bool {
		return scored[i].score < scored[j].score
	})
	return scored
}

func tournamentSelection(scored []Scored, k int) [][]int {
	selected := make([][]int, len(scored))
	for i := range selected {
		best := Scored{score: math.Inf(1)}
		for j := 0; j < k; j++ {
			cand := scored[rand.Intn(len(scored))]
			if cand.score < best.score {
				best = cand
			}
		}
		selected[i] = clone(best.ind)
	}
	return selected
}

func orderCrossover(p1, p2 []int) []int {
	n := len(p1)
	a, b := rand.Intn(n), rand.Intn(n)
	if a > b {
		a, b = b, a
	}

	child := make([]int, n)
	for i := range child {
		child[i] = -1
	}

	copy(child[a:b], p1[a:b])

	ptr := b
	for _, gene := range p2 {
		if !contains(child, gene) {
			if ptr >= n {
				ptr = 0
			}
			child[ptr] = gene
			ptr++
		}
	}
	return child
}

func mutate(route []int, rate float64) {
	for i := range route {
		if rand.Float64() < rate {
			j := rand.Intn(len(route))
			route[i], route[j] = route[j], route[i]
		}
	}
}

func nextGeneration(parents [][]int, oldPop [][]int, dist [][]float64, elitism int) [][]int {
	scored := evaluate(oldPop, dist)
	newPop := make([][]int, 0, len(oldPop))

	// elites
	for i := 0; i < elitism; i++ {
		newPop = append(newPop, clone(scored[i].ind))
	}

	// children
	for len(newPop) < len(oldPop) {
		p1 := parents[rand.Intn(len(parents))]
		p2 := parents[rand.Intn(len(parents))]
		child := orderCrossover(p1, p2)
		mutate(child, 0.03)
		newPop = append(newPop, child)
	}

	return newPop
}

func twoOpt(route []int, dist [][]float64) []int {
	best := clone(route)
	bestLen := pathLength(best, dist)
	improved := true

	for improved {
		improved = false
		for i := 1; i < len(best)-2; i++ {
			for j := i + 1; j < len(best); j++ {
				if j-i == 1 {
					continue
				}
				newRoute := clone(best)
				reverse(newRoute[i:j])
				newLen := pathLength(newRoute, dist)
				if newLen < bestLen {
					best = newRoute
					bestLen = newLen
					improved = true
				}
			}
		}
	}
	return best
}

func genetic(points []Point, generations, popSize int, dist [][]float64) ([]int, float64) {
	n := len(points)
	population := initPopulation(popSize, n)
	scored := evaluate(population, dist)

	fmt.Println(scored[0].score)

	for t := 1; t <= generations; t++ {
		parents := tournamentSelection(scored, 5)

		children := make([][]int, popSize)
		for i := 0; i < popSize; i++ {
			p1 := parents[rand.Intn(len(parents))]
			p2 := parents[rand.Intn(len(parents))]
			c := orderCrossover(p1, p2)
			mutate(c, 0.03)
			children[i] = c
		}

		scoredChildren := evaluate(children, dist)

		// apply 2-opt
		if n <= 20 {
			k := popSize / 10
			if k < 1 {
				k = 1
			}
			for i := 0; i < k; i++ {
				r := scoredChildren[i].ind
				improved := twoOpt(r, dist)
				scoredChildren[i].ind = improved
				scoredChildren[i].score = pathLength(improved, dist)
			}
		} else {
			best := scoredChildren[0].ind
			best = twoOpt(best, dist)
			scoredChildren[0].ind = best
			scoredChildren[0].score = pathLength(best, dist)
		}

		population = nextGeneration(extract(scoredChildren), population, dist, 1)
		scored = evaluate(population, dist)

		if t == generations || t%(generations/9) == 0 {
			fmt.Println(scored[0].score)
		}
	}

	return scored[0].ind, scored[0].score
}

func readInput() (string, []string, []Point) {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	first := strings.TrimSpace(in.Text())

	if _, err := strconv.Atoi(first); err == nil {
		// RANDOM mode
		N, _ := strconv.Atoi(first)
		pts := make([]Point, N)
		for i := 0; i < N; i++ {
			pts[i] = Point{
				x: rand.Float64() * 1000,
				y: rand.Float64() * 1000,
			}
		}
		return "RANDOM", nil, pts
	}

	// named dataset
	datasetName := first

	in.Scan()
	cityCount, _ := strconv.Atoi(strings.TrimSpace(in.Text()))

	cities := make([]string, cityCount)
	points := make([]Point, cityCount)

	for i := 0; i < cityCount; i++ {
		in.Scan()
		parts := strings.Fields(in.Text())
		name := parts[0]
		x, _ := strconv.ParseFloat(parts[1], 64)
		y, _ := strconv.ParseFloat(parts[2], 64)

		cities[i] = name
		points[i] = Point{x, y}
	}

	return datasetName, cities, points
}

func main() {
	rand.Seed(time.Now().UnixNano())

	dataset, cityNames, points := readInput()

	n := len(points)
	dist := make([][]float64, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			dist[i][j] = distance(points[i], points[j])
		}
	}

	route, best := genetic(points, 200, 250, dist)

	fmt.Println()

	if dataset == "RANDOM" {
		fmt.Println(best)
	} else {
		out := make([]string, len(route))
		for i, idx := range route {
			out[i] = cityNames[idx]
		}
		fmt.Println(strings.Join(out, " -> "))
		fmt.Println(best)
	}
}

//
// --- Utility helpers ---
//

func clone(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	return b
}

func contains(a []int, x int) bool {
	for _, v := range a {
		if v == x {
			return true
		}
	}
	return false
}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func extract(sc []Scored) [][]int {
	out := make([][]int, len(sc))
	for i := range sc {
		out[i] = sc[i].ind
	}
	return out
}
