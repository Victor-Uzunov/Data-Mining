GRAPH = {
    'A': ['B', 'C'],
    'B': ['D'],
    'C': ['G'],
    'D': ['G'],
    'G': []
}
HEURISTICS = {
    'A': 10, 'B': 1, 'C': 5, 'D': 2, 'G': 0
}

def beam_search(graph: dict, heuristics: dict, start: str, goal: str, beam_width: int) -> list or None: # type: ignore
    h_start = heuristics.get(start, float('inf'))
    current_level = [(h_start, start, [start])]

    visited = {start}
    while current_level:
        for h, node, path in current_level:
            if node == goal:
                return path
        
        next_level = []

        for h, parent, path in current_level:
            for child in graph.get(parent, []):
                if child not in visited:
                    h_child = heuristics.get(child, float('inf'))
                    next_level.append((h_child, child, path + [child]))
        if not next_level:
            return None
        
        next_level.sort(key=lambda x: x[0])
        current_level = next_level[:beam_width]

        for _, node, _ in current_level:
            visited.add(node)

    return None

if __name__ == "__main__":
    start_node = 'A'
    goal_node = 'G'
    
    for k in [1, 2, 3]:
        path = beam_search(GRAPH, HEURISTICS, start_node, goal_node, beam_width=k)
        if path:
            print(f"Beam width k={k}: Path found -> {path}")
        else:
            print(f"Beam width k={k}: Goal not found")