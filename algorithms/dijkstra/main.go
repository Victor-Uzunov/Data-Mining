package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	Neighbor string
	Weight   int
}

type Graph map[string][]Edge

type Item struct {
	name  string
	cost  int
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func Dijkstra(graph Graph, startNode string) (map[string]int, map[string]string) {
	distances := make(map[string]int)
	predecessors := make(map[string]string)

	for node := range graph {
		distances[node] = math.MaxInt32
		predecessors[node] = ""
	}
	distances[startNode] = 0

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	startItem := &Item{name: startNode, cost: 0}
	heap.Push(&pq, startItem)

	for pq.Len() > 0 {
		currentMinItem := heap.Pop(&pq).(*Item)
		currentNode := currentMinItem.name
		currentCost := currentMinItem.cost

		if currentCost > distances[currentNode] {
			continue
		}

		for _, edge := range graph[currentNode] {
			neighbor := edge.Neighbor
			weight := edge.Weight

			newDistance := currentCost + weight

			if newDistance < distances[neighbor] {
				distances[neighbor] = newDistance
				predecessors[neighbor] = currentNode

				heap.Push(&pq, &Item{name: neighbor, cost: newDistance})
			}
		}
	}

	return distances, predecessors
}

func main() {
	testGraph := Graph{
		"A": {{Neighbor: "B", Weight: 6}, {Neighbor: "D", Weight: 1}},
		"B": {{Neighbor: "E", Weight: 2}, {Neighbor: "C", Weight: 5}},
		"C": {},
		"D": {{Neighbor: "B", Weight: 2}, {Neighbor: "E", Weight: 1}},
		"E": {{Neighbor: "C", Weight: 5}},
	}
	startNode := "A"

	shortestDistances, predecessors := Dijkstra(testGraph, startNode)

	fmt.Printf("--- Dijkstra's Algorithm Results from Node %s ---\n", startNode)
	fmt.Println("Shortest Distances:")
	for node, dist := range shortestDistances {
		if dist == math.MaxInt32 {
			fmt.Printf("  %s: Unreachable\n", node)
		} else {
			fmt.Printf("  %s: %d\n", node, dist)
		}
	}

	fmt.Println("\nPredecessors for Path Reconstruction:")
	fmt.Println(predecessors)
}
