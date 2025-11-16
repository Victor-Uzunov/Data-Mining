package main

import (
	"fmt"
	"math/rand"
	"sort"
)

const (
	MAXIMUM_GENERATIONS             = 1600
	MUTATION_RATE                   = 0.01
	TOURNAMENT_SIZE                 = 15
	POPULATION_SIZE                 = 250
	NUMBER_OF_GENERATIONS_FOR_PRINT = 8
	ELITE_COUNT                     = 7
)

type Knapsack struct {
	numItems   int
	capacity   float64
	weights    []float64
	values     []float64
	population [][]bool
}

func NewKnapsack(numItems int, capacity float64, weights, values []float64) *Knapsack {
	k := &Knapsack{
		numItems: numItems,
		capacity: capacity,
		weights:  weights,
		values:   values,
	}
	k.population = k.initializePopulation()
	return k
}

// true means pick item on position i, false means skip it
func (k *Knapsack) initializePopulation() [][]bool {
	population := make([][]bool, 0, POPULATION_SIZE)

	for i := 0; i < POPULATION_SIZE; i++ {
		chromosome := make([]bool, k.numItems)
		totalWeight := 0.0

		for j := 0; j < k.numItems; j++ {
			getItem := rand.Intn(2) == 1
			if getItem && totalWeight+k.weights[j] <= k.capacity {
				chromosome[j] = getItem
				totalWeight += k.weights[j]
			}
		}
		population = append(population, chromosome)
	}
	return population
}

// Checks if weight of items we pick is below CAPACITY, if it is above return 0, otherwise its Value
func (k *Knapsack) EvaluateChromosome(chromosome []bool) float64 {
	chromosomeWeight := 0.0
	chromosomeValue := 0.0

	for i := 0; i < k.numItems; i++ {
		if chromosome[i] {
			chromosomeWeight += k.weights[i]
			chromosomeValue += k.values[i]
		}
	}

	if chromosomeWeight > k.capacity {
		return 0
	}

	return chromosomeValue
}

// Select a parent using tournament selection
func (k *Knapsack) tournamentSelection() []bool {
	currentBestIndividual := k.population[rand.Intn(POPULATION_SIZE)]

	for i := 0; i < TOURNAMENT_SIZE; i++ {
		individual := k.population[rand.Intn(POPULATION_SIZE)]

		// Check if same individual (comparing pointers)
		same := true
		for j := range individual {
			if individual[j] != currentBestIndividual[j] {
				same = false
				break
			}
		}

		if same {
			i--
			continue
		}

		if k.EvaluateChromosome(individual) > k.EvaluateChromosome(currentBestIndividual) {
			currentBestIndividual = individual
		}
	}

	return currentBestIndividual
}

func (k *Knapsack) twoPointCrossover(parent1, parent2 []bool) [][]bool {
	length := len(parent1)
	crossoverPoint1 := rand.Intn(length)
	crossoverPoint2 := rand.Intn(length-crossoverPoint1) + crossoverPoint1

	child1 := make([]bool, length)
	child2 := make([]bool, length)
	copy(child1, parent1)
	copy(child2, parent2)

	for i := crossoverPoint1; i < crossoverPoint2; i++ {
		child1[i] = parent2[i]
		child2[i] = parent1[i]
	}

	children := [][]bool{child1, child2}
	return children
}

// Mutate with a small probability
func (k *Knapsack) mutate(chromosome []bool) {
	for i := 0; i < k.numItems; i++ {
		if rand.Float64() < MUTATION_RATE {
			chromosome[i] = !chromosome[i]
		}
	}
}

func (k *Knapsack) Solve() []bool {
	bestValue := 0.0
	var bestSolution []bool
	newPopulation := make([][]bool, 0, POPULATION_SIZE)

	iterationsToPrint := make(map[int]bool)

	for len(iterationsToPrint) < NUMBER_OF_GENERATIONS_FOR_PRINT {
		iterationsToPrint[rand.Intn(MAXIMUM_GENERATIONS-3)+2] = true
	}
	iterationsToPrint[0] = true
	iterationsToPrint[MAXIMUM_GENERATIONS] = true

	for generation := 0; generation < MAXIMUM_GENERATIONS; generation++ {

		// We take the best chromosomes from the population into the new one
		sort.Slice(k.population, func(i, j int) bool {
			fitnessI := k.EvaluateChromosome(k.population[i])
			fitnessJ := k.EvaluateChromosome(k.population[j])
			return fitnessI > fitnessJ
		})

		if iterationsToPrint[generation] {
			fmt.Println(k.EvaluateChromosome(k.population[0]))
		}

		for i := 0; i < ELITE_COUNT; i++ {
			newPopulation = append(newPopulation, k.population[i])
		}

		for len(newPopulation) < POPULATION_SIZE {
			// Selection
			firstParent := k.tournamentSelection()
			secondParent := k.tournamentSelection()

			// Crossover
			children := k.twoPointCrossover(firstParent, secondParent)

			// Mutation
			k.mutate(children[0])
			k.mutate(children[1])

			newPopulation = append(newPopulation, children[0])
			if len(newPopulation) < POPULATION_SIZE {
				newPopulation = append(newPopulation, children[1])
			}
		}

		for i := 0; i < POPULATION_SIZE; i++ {
			k.population[i] = newPopulation[i]
		}

		// Update the best solution if we find a better one
		for _, individual := range k.population {
			currValue := k.EvaluateChromosome(individual)
			if currValue > bestValue {
				bestValue = currValue
				bestSolution = make([]bool, len(individual))
				copy(bestSolution, individual)
			}
		}

		newPopulation = newPopulation[:0]
	}

	return bestSolution
}

func main() {
	var maxWeight float64
	var countItems int

	fmt.Scan(&maxWeight)
	fmt.Scan(&countItems)

	weights := make([]float64, countItems)
	values := make([]float64, countItems)

	for i := 0; i < countItems; i++ {
		var w, v int
		fmt.Scan(&w, &v)
		weights[i] = float64(w)
		values[i] = float64(v)
	}

	knapsack := NewKnapsack(countItems, maxWeight, weights, values)
	// Run the genetic algorithm and get the best solution
	bestCombination := knapsack.Solve()
	bestValue := knapsack.EvaluateChromosome(bestCombination)

	fmt.Printf("Best solution is: %v\n", bestValue)
	fmt.Println("Items included in it are: ")

	for i := 0; i < len(bestCombination); i++ {
		if bestCombination[i] {
			fmt.Printf("Item %d\n", i+1)
		}
	}
}
