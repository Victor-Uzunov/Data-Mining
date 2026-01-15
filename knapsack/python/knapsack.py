#!/usr/bin/env python3
import sys
import random
import time
import os
import math

# Constants
MAXIMUM_GENERATIONS = 2000
MUTATION_RATE = 0.015
TOURNAMENT_SIZE = 20
POPULATION_SIZE = 300
NUMBER_OF_GENERATIONS_PRINT = 8
ELITE_COUNT = 10

class Item:
    def __init__(self, weight, value):
        self.weight = float(weight)
        self.value = float(value)
        # Handle division by zero if weight is 0
        self.ratio = (value / weight) if weight > 0 else 0.0

class Chromosome:
    def __init__(self, size, genes=None):
        if genes is None:
            self.genes = [False] * size
        else:
            self.genes = genes
        self.fitness = 0.0
        self.valid = False

class GeneticAlgorithm:
    def __init__(self, capacity, items):
        self.capacity = float(capacity)
        self.items = items
        self.num_items = len(items)
        self.population = []
        # Python's random is seeded by system time by default
        self.initialize_population()

    def initialize_population(self):
        greedy_count = POPULATION_SIZE // 5

        # 1. Ratio Greedy with random threshold
        for _ in range(greedy_count):
            threshold = 0.8 + random.random() * 0.2
            self.population.append(self.greedy_chromosome(threshold))

        # 2. Value Greedy
        for _ in range(greedy_count, 2 * greedy_count):
            self.population.append(self.value_greedy_chromosome())

        # 3. Random initialization
        for _ in range(2 * greedy_count, POPULATION_SIZE):
            genes = [False] * self.num_items
            total_weight = 0.0

            for j in range(self.num_items):
                # 30% chance to include item if it fits
                if random.random() < 0.3 and total_weight + self.items[j].weight <= self.capacity:
                    genes[j] = True
                    total_weight += self.items[j].weight

            self.population.append(Chromosome(self.num_items, genes))

    def greedy_chromosome(self, threshold):
        # Create indices sorted by ratio (descending)
        indices = list(range(self.num_items))
        indices.sort(key=lambda i: self.items[i].ratio, reverse=True)

        genes = [False] * self.num_items
        total_weight = 0.0

        for idx in indices:
            if random.random() < threshold and total_weight + self.items[idx].weight <= self.capacity:
                genes[idx] = True
                total_weight += self.items[idx].weight

        return Chromosome(self.num_items, genes)

    def value_greedy_chromosome(self):
        # Create indices sorted by value (descending)
        indices = list(range(self.num_items))
        indices.sort(key=lambda i: self.items[i].value, reverse=True)

        genes = [False] * self.num_items
        total_weight = 0.0

        for idx in indices:
            if total_weight + self.items[idx].weight <= self.capacity:
                genes[idx] = True
                total_weight += self.items[idx].weight

        return Chromosome(self.num_items, genes)

    def evaluate_fitness(self, chromosome):
        if chromosome.valid:
            return chromosome.fitness

        total_weight = 0.0
        total_value = 0.0

        for i in range(self.num_items):
            if chromosome.genes[i]:
                total_weight += self.items[i].weight
                total_value += self.items[i].value

        chromosome.valid = True
        if total_weight > self.capacity:
            chromosome.fitness = 0.0
            return 0.0

        chromosome.fitness = total_value
        return total_value

    def tournament_selection(self):
        best_idx = random.randrange(POPULATION_SIZE)
        best_fitness = self.evaluate_fitness(self.population[best_idx])

        for _ in range(1, TOURNAMENT_SIZE):
            idx = random.randrange(POPULATION_SIZE)
            if idx == best_idx:
                continue

            fitness = self.evaluate_fitness(self.population[idx])
            if fitness > best_fitness:
                best_idx = idx
                best_fitness = fitness

        return self.population[best_idx]

    def uniform_crossover(self, parent1, parent2):
        genes1 = [False] * self.num_items
        genes2 = [False] * self.num_items

        for i in range(self.num_items):
            if random.random() < 0.5:
                genes1[i] = parent1.genes[i]
                genes2[i] = parent2.genes[i]
            else:
                genes1[i] = parent2.genes[i]
                genes2[i] = parent1.genes[i]

        return Chromosome(self.num_items, genes1), Chromosome(self.num_items, genes2)

    def mutate(self, chromosome):
        for i in range(self.num_items):
            if random.random() < MUTATION_RATE:
                chromosome.genes[i] = not chromosome.genes[i]
                chromosome.valid = False

    def repair(self, chromosome):
        total_weight = sum(self.items[i].weight for i in range(self.num_items) if chromosome.genes[i])

        if total_weight <= self.capacity:
            return

        # Collect indices of selected items
        selected_indices = [i for i in range(self.num_items) if chromosome.genes[i]]

        # Sort selected items by ratio ascending (remove least valuable per weight first)
        selected_indices.sort(key=lambda i: self.items[i].ratio)

        for idx in selected_indices:
            if total_weight <= self.capacity:
                break
            chromosome.genes[idx] = False
            total_weight -= self.items[idx].weight

        chromosome.valid = False

    def solve(self):
        best_solution = [False] * self.num_items
        best_value = 0.0

        # Calculate print generations to match Go logic
        print_gens = [0]
        step = MAXIMUM_GENERATIONS // (NUMBER_OF_GENERATIONS_PRINT + 1)
        for i in range(1, NUMBER_OF_GENERATIONS_PRINT + 1):
            print_gens.append(i * step)
        print_gens.append(MAXIMUM_GENERATIONS - 1)

        print_set = set(print_gens)
        output_values = []

        for generation in range(MAXIMUM_GENERATIONS):
            # Evaluate all
            for chrome in self.population:
                self.evaluate_fitness(chrome)

            # Sort by fitness descending
            self.population.sort(key=lambda x: x.fitness, reverse=True)

            if generation in print_set:
                output_values.append(self.population[0].fitness)

            # Update best global
            if self.population[0].fitness > best_value:
                best_value = self.population[0].fitness
                best_solution = list(self.population[0].genes)

            new_population = []

            # Elitism
            for i in range(ELITE_COUNT):
                elite = Chromosome(self.num_items, list(self.population[i].genes))
                elite.fitness = self.population[i].fitness
                elite.valid = self.population[i].valid
                new_population.append(elite)

            # Generate rest of population
            while len(new_population) < POPULATION_SIZE:
                parent1 = self.tournament_selection()
                parent2 = self.tournament_selection()

                child1, child2 = self.uniform_crossover(parent1, parent2)

                self.mutate(child1)
                self.mutate(child2)

                self.repair(child1)
                self.repair(child2)

                new_population.append(child1)
                if len(new_population) < POPULATION_SIZE:
                    new_population.append(child2)

            self.population = new_population

        return best_solution, best_value, output_values

def main():
    # Use buffered input reading for speed similar to Go's Scanner/Bufio
    input_data = sys.stdin.read().split()
    if not input_data:
        return

    iterator = iter(input_data)

    try:
        capacity = float(next(iterator))
        num_items = int(next(iterator))

        items = []
        for _ in range(num_items):
            w = float(next(iterator))
            v = float(next(iterator))
            items.append(Item(w, v))

    except StopIteration:
        return

    time_only = os.environ.get("FMI_TIME_ONLY") == "1"

    start_time = time.time()

    ga = GeneticAlgorithm(capacity, items)
    _, best_value, output_values = ga.solve()

    end_time = time.time()
    elapsed_ms = (end_time - start_time) * 1000

    if time_only:
        print(f"# TIMES_MS: alg={elapsed_ms:.3f}")
    else:
        for val in output_values:
            print(f"{val:.0f}")
        print()
        print(f"{best_value:.0f}")

if __name__ == "__main__":
    main()