import heapq
import math

Graph = dict

def dijkstra(graph: Graph, start_node: str):
    distances = {node: math.inf for node in graph}
    predecessors = {node: None for node in graph}

    distances[start_node] = 0
    pq = [(0, start_node)]

    while pq:
        current_cost, current_node = heapq.heappop(pq)

        if current_cost > distances[current_node]:
            continue

        for neighbor, weight in graph[current_node]:
            new_distance = current_cost + weight
            if new_distance < distances[neighbor]:
                distances[neighbor] = new_distance
                predecessors[neighbor] = current_node
                heapq.heappush(pq, (new_distance, neighbor))

    return distances, predecessors

if __name__ == "__main__":
    test_graph = {
        "A": [("B", 6), ("D", 1)],
        "B": [("E", 2), ("C", 5)],
        "C": [],
        "D": [("B", 2), ("E", 1)],
        "E": [("C", 5)]
    }

    start_node = "A"
    shortest_distances, predecessors = dijkstra(test_graph, start_node)

    print(f"--- Dijkstra's Algorithm Results from Node {start_node} ---")
    print("Shortest Distances:")
    for node, dist in shortest_distances.items():
        if dist == math.inf:
            print(f"  {node}: Unreachable")
        else:
            print(f"  {node}: {dist}")

    print("\nPredecessors for Path Reconstruction:")
    print(predecessors)