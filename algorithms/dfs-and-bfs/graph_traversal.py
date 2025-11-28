from collections import deque

# Graph is represented as a dictionary where keys are nodes and values are lists of neighbors.
Graph = dict

def depth_first_search(graph: Graph, start_node: str):
    visited = set()
    result = []
    stack = [start_node]

    while stack:
        current = stack.pop()

        if current not in visited:
            visited.add(current)
            result.append(current)

            # Push neighbors in reverse order to mimic Go implementation
            for neighbor in reversed(graph.get(current, [])):
                if neighbor not in visited:
                    stack.append(neighbor)

    return result


def depth_first_search_recursive(graph: Graph, start_node: str):
    visited = set()
    result = []

    def dfs(node):
        if node in visited:
            return
        visited.add(node)
        result.append(node)

        for neighbor in graph.get(node, []):
            dfs(neighbor)

    dfs(start_node)
    return result


def breadth_first_search(graph: Graph, start_node: str):
    visited = set()
    queue = deque([start_node])
    result = []

    while queue:
        node = queue.popleft()

        if node in visited:
            continue

        visited.add(node)
        result.append(node)

        for neighbor in graph.get(node, []):
            if neighbor not in visited:
                queue.append(neighbor)

    return result


if __name__ == "__main__":
    graph = {
        "A": ["B", "C"],
        "B": ["D", "E"],
        "C": ["F"],
        "D": [],
        "E": ["F"],
        "F": []
    }

    print("DFS Traversal Order:")
    print(depth_first_search(graph, "A"))

    print("DFS Recursive Traversal Order:")
    print(depth_first_search_recursive(graph, "A"))

    print("BFS Traversal Order:")
    print(breadth_first_search(graph, "A"))