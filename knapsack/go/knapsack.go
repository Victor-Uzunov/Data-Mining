package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

const (
	MAXIMUM_GENERATIONS         = 2000
	MUTATION_RATE               = 0.015
	TOURNAMENT_SIZE             = 20
	POPULATION_SIZE             = 300
	NUMBER_OF_GENERATIONS_PRINT = 8
	ELITE_COUNT                 = 10
)

type Item struct {
	weight float64
	value  float64
	ratio  float64
}

type Chromosome struct {
	genes   []bool
	fitness float64
	valid   bool
}

type GeneticAlgorithm struct {
	items      []Item
	capacity   float64
	population []Chromosome
	rng        *rand.Rand
	numItems   int
}

func NewGeneticAlgorithm(capacity float64, items []Item) *GeneticAlgorithm {
	for i := range items {
		if items[i].weight > 0 {
			items[i].ratio = items[i].value / items[i].weight
		}
	}

	ga := &GeneticAlgorithm{
		items:      items,
		capacity:   capacity,
		population: make([]Chromosome, POPULATION_SIZE),
		rng:        rand.New(rand.NewSource(time.Now().UnixNano())),
		numItems:   len(items),
	}
	ga.initializePopulation()
	return ga
}

func (ga *GeneticAlgorithm) initializePopulation() {
	greedyCount := POPULATION_SIZE / 5
	for i := 0; i < greedyCount; i++ {
		ga.population[i] = ga.greedyChromosome(0.8 + ga.rng.Float64()*0.2)
	}

	for i := greedyCount; i < 2*greedyCount; i++ {
		ga.population[i] = ga.valueGreedyChromosome()
	}

	for i := 2 * greedyCount; i < POPULATION_SIZE; i++ {
		chromosome := Chromosome{genes: make([]bool, ga.numItems)}
		totalWeight := 0.0

		for j := 0; j < ga.numItems; j++ {
			if ga.rng.Float64() < 0.3 && totalWeight+ga.items[j].weight <= ga.capacity {
				chromosome.genes[j] = true
				totalWeight += ga.items[j].weight
			}
		}
		ga.population[i] = chromosome
	}
}

func (ga *GeneticAlgorithm) greedyChromosome(threshold float64) Chromosome {
	indices := make([]int, ga.numItems)
	for i := range indices {
		indices[i] = i
	}

	sort.Slice(indices, func(i, j int) bool {
		return ga.items[indices[i]].ratio > ga.items[indices[j]].ratio
	})

	chromosome := Chromosome{genes: make([]bool, ga.numItems)}
	totalWeight := 0.0

	for _, idx := range indices {
		if ga.rng.Float64() < threshold && totalWeight+ga.items[idx].weight <= ga.capacity {
			chromosome.genes[idx] = true
			totalWeight += ga.items[idx].weight
		}
	}

	return chromosome
}

func (ga *GeneticAlgorithm) valueGreedyChromosome() Chromosome {
	indices := make([]int, ga.numItems)
	for i := range indices {
		indices[i] = i
	}

	sort.Slice(indices, func(i, j int) bool {
		return ga.items[indices[i]].value > ga.items[indices[j]].value
	})

	chromosome := Chromosome{genes: make([]bool, ga.numItems)}
	totalWeight := 0.0

	for _, idx := range indices {
		if totalWeight+ga.items[idx].weight <= ga.capacity {
			chromosome.genes[idx] = true
			totalWeight += ga.items[idx].weight
		}
	}

	return chromosome
}

func (ga *GeneticAlgorithm) evaluateFitness(chromosome *Chromosome) float64 {
	if chromosome.valid {
		return chromosome.fitness
	}

	totalWeight := 0.0
	totalValue := 0.0

	for i := 0; i < ga.numItems; i++ {
		if chromosome.genes[i] {
			totalWeight += ga.items[i].weight
			totalValue += ga.items[i].value
		}
	}

	chromosome.valid = true
	if totalWeight > ga.capacity {
		chromosome.fitness = 0
		return 0
	}

	chromosome.fitness = totalValue
	return totalValue
}

func (ga *GeneticAlgorithm) tournamentSelection() *Chromosome {
	bestIdx := ga.rng.Intn(POPULATION_SIZE)
	bestFitness := ga.evaluateFitness(&ga.population[bestIdx])

	for i := 1; i < TOURNAMENT_SIZE; i++ {
		idx := ga.rng.Intn(POPULATION_SIZE)
		if idx == bestIdx {
			continue
		}
		fitness := ga.evaluateFitness(&ga.population[idx])
		if fitness > bestFitness {
			bestIdx = idx
			bestFitness = fitness
		}
	}

	return &ga.population[bestIdx]
}

func (ga *GeneticAlgorithm) uniformCrossover(parent1, parent2 *Chromosome) (Chromosome, Chromosome) {
	child1 := Chromosome{genes: make([]bool, ga.numItems)}
	child2 := Chromosome{genes: make([]bool, ga.numItems)}

	for i := 0; i < ga.numItems; i++ {
		if ga.rng.Float64() < 0.5 {
			child1.genes[i] = parent1.genes[i]
			child2.genes[i] = parent2.genes[i]
		} else {
			child1.genes[i] = parent2.genes[i]
			child2.genes[i] = parent1.genes[i]
		}
	}

	return child1, child2
}

func (ga *GeneticAlgorithm) mutate(chromosome *Chromosome) {
	for i := 0; i < ga.numItems; i++ {
		if ga.rng.Float64() < MUTATION_RATE {
			chromosome.genes[i] = !chromosome.genes[i]
			chromosome.valid = false
		}
	}
}

func (ga *GeneticAlgorithm) repair(chromosome *Chromosome) {
	totalWeight := 0.0
	for i := 0; i < ga.numItems; i++ {
		if chromosome.genes[i] {
			totalWeight += ga.items[i].weight
		}
	}

	if totalWeight <= ga.capacity {
		return
	}

	type itemIndex struct {
		idx   int
		ratio float64
	}

	selected := make([]itemIndex, 0)
	for i := 0; i < ga.numItems; i++ {
		if chromosome.genes[i] {
			selected = append(selected, itemIndex{i, ga.items[i].ratio})
		}
	}

	sort.Slice(selected, func(i, j int) bool {
		return selected[i].ratio < selected[j].ratio
	})

	for _, item := range selected {
		if totalWeight <= ga.capacity {
			break
		}
		chromosome.genes[item.idx] = false
		totalWeight -= ga.items[item.idx].weight
	}

	chromosome.valid = false
}

func (ga *GeneticAlgorithm) Solve() ([]bool, float64, []float64) {
	bestSolution := make([]bool, ga.numItems)
	bestValue := 0.0

	printGens := []int{0}
	step := MAXIMUM_GENERATIONS / (NUMBER_OF_GENERATIONS_PRINT + 1)
	for i := 1; i <= NUMBER_OF_GENERATIONS_PRINT; i++ {
		printGens = append(printGens, i*step)
	}
	printGens = append(printGens, MAXIMUM_GENERATIONS-1)

	printMap := make(map[int]bool)
	for _, g := range printGens {
		printMap[g] = true
	}

	outputValues := make([]float64, 0, len(printGens))

	for generation := 0; generation < MAXIMUM_GENERATIONS; generation++ {
		for i := range ga.population {
			ga.evaluateFitness(&ga.population[i])
		}

		sort.Slice(ga.population, func(i, j int) bool {
			return ga.population[i].fitness > ga.population[j].fitness
		})

		if printMap[generation] {
			outputValues = append(outputValues, ga.population[0].fitness)
		}

		if ga.population[0].fitness > bestValue {
			bestValue = ga.population[0].fitness
			copy(bestSolution, ga.population[0].genes)
		}

		newPopulation := make([]Chromosome, 0, POPULATION_SIZE)

		for i := 0; i < ELITE_COUNT; i++ {
			elite := Chromosome{
				genes:   make([]bool, ga.numItems),
				fitness: ga.population[i].fitness,
				valid:   ga.population[i].valid,
			}
			copy(elite.genes, ga.population[i].genes)
			newPopulation = append(newPopulation, elite)
		}

		for len(newPopulation) < POPULATION_SIZE {
			parent1 := ga.tournamentSelection()
			parent2 := ga.tournamentSelection()

			child1, child2 := ga.uniformCrossover(parent1, parent2)

			ga.mutate(&child1)
			ga.mutate(&child2)

			ga.repair(&child1)
			ga.repair(&child2)

			newPopulation = append(newPopulation, child1)
			if len(newPopulation) < POPULATION_SIZE {
				newPopulation = append(newPopulation, child2)
			}
		}

		ga.population = newPopulation
	}

	return bestSolution, bestValue, outputValues
}

func main() {
	timeOnly := os.Getenv("FMI_TIME_ONLY") == "1"

	reader := bufio.NewReader(os.Stdin)
	var capacity float64
	var numItems int
	fmt.Fscan(reader, &capacity, &numItems)

	items := make([]Item, numItems)
	for i := 0; i < numItems; i++ {
		fmt.Fscan(reader, &items[i].weight, &items[i].value)
	}

	start := time.Now()
	ga := NewGeneticAlgorithm(capacity, items)
	_, bestValue, outputValues := ga.Solve()
	elapsed := time.Since(start)

	elapsedMs := float64(elapsed.Nanoseconds()) / 1e6

	if timeOnly {
		fmt.Printf("# TIMES_MS: alg=%.3f\n", elapsedMs)
	} else {
		for _, val := range outputValues {
			fmt.Printf("%.0f\n", val)
		}
		fmt.Println()
		fmt.Printf("%.0f\n", bestValue)
	}
}
