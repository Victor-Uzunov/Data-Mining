import math
MAX_VALUE = float('inf')
MIN_VALUE = float('-inf')

def minimax_alpha_beta(scores: list, depth: int, node_index: int, is_max: bool, alpha: float, beta: float) -> int:
    if depth == 0:
        return scores[node_index]
    if is_max:
        max_eval = MIN_VALUE
        for i in range(2):
            child_value = minimax_alpha_beta(scores, depth - 1, node_index * 2 + i, False, alpha, beta)
            max_eval = max(max_eval, child_value)
            alpha = max(alpha, max_eval)
            if beta <= alpha:
                break
        return max_eval
    else:
        min_eval = MAX_VALUE
        for i in range(2):
            child_value = minimax_alpha_beta(scores, depth - 1, node_index * 2 + i, True, alpha, beta)
            min_eval = min(min_eval, child_value)
            beta = min(beta, min_eval)
            if beta <= alpha:
                break
        return min_eval
    
def calculate_depth(scores: list) -> int:
    num_leaves = len(scores)
    if num_leaves == 0:
        return 0
    
    is_power_of_two = (num_leaves & (num_leaves - 1)) == 0
    if not is_power_of_two:
        raise ValueError("Number of scores must be a power of two.")
    depth = int(math.log2(num_leaves))
    return depth

    
if __name__ == "__main__":
    scores = [3, 5, 6, 9, 1, 2, 0, -1]
    depth = 3
    optimal_value = minimax_alpha_beta(scores, depth, 0, True, MIN_VALUE, MAX_VALUE)
    print(f"The optimal value is: {optimal_value}")