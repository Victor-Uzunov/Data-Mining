package main

// Graph is represented as a map where keys are nodes and values are lists of neighbors.
type Graph map[string][]string

func DepthFirstSearch(graph Graph, startNode string) []string {
	visited := map[string]bool{}
	var result []string
	stack := []string{startNode}

	for len(stack) > 0 {
		// Pop a node from the stack
		currentNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !visited[currentNode] {
			visited[currentNode] = true
			result = append(result, currentNode)

			// Push all unvisited neighbors onto the stack
			for i := len(graph[currentNode]) - 1; i >= 0; i-- {
				neighbor := graph[currentNode][i]
				if !visited[neighbor] {
					stack = append(stack, neighbor)
				}
			}
		}
	}
	return result
}

func DepthFirstSearchRecursive(graph Graph, startNode string) []string {
	var result []string
	var dfs func(graph Graph, node string, visited map[string]bool)
	visited := map[string]bool{}

	dfs = func(graph Graph, node string, visited map[string]bool) {
		if visited[node] {
			return
		}
		visited[node] = true
		result = append(result, node)

		for _, neighbor := range graph[node] {
			dfs(graph, neighbor, visited)
		}
	}

	dfs(graph, startNode, visited)
	return result
}

func BreadthFirstSearch(graph Graph, startNode string) []string {
	visited := map[string]bool{}
	queue := []string{startNode}
	result := make([]string, 0)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if visited[node] {
			continue
		}

		visited[node] = true
		result = append(result, node)

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}

	return result
}

func main() {
	// Example usage
	graph := Graph{
		"A": {"B", "C"},
		"B": {"D", "E"},
		"C": {"F"},
		"D": {},
		"E": {"F"},
		"F": {},
	}

	result := DepthFirstSearch(graph, "A")
	println("DFS Traversal Order:")
	for _, node := range result {
		println(node)
	}

	resultRecursive := DepthFirstSearchRecursive(graph, "A")
	println("DFS Recursive Traversal Order:")
	for _, node := range resultRecursive {
		println(node)
	}

	resultBFS := BreadthFirstSearch(graph, "A")
	println("BFS Traversal Order:")
	for _, node := range resultBFS {
		println(node)
	}
}
