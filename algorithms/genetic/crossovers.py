def single_point_crossover(parent1: list, parent2: list, crossover_point: int) -> tuple[list, list]:
    prefix1 = parent1[:crossover_point]
    suffix1 = parent1[crossover_point:]
    prefix2 = parent2[:crossover_point]
    suffix2 = parent2[crossover_point:]

    return prefix1 + suffix2, prefix2 + suffix1

def two_point_crossover(parent1: list, parent2: list, point1: int, point2: int) -> tuple[list, list]:
    if point1 > point2:
        point1, point2 = point2, point1

    prefi1 = parent1[:point1]
    mid1 = parent1[point1:point2]
    suf1 = parent1[point2:]

    prefi2 = parent2[:point1]
    mid2 = parent2[point1:point2]
    suf2 = parent2[point2:]

    return prefi1 + mid2 + suf1, prefi2 + mid1 + suf2

def uniform_crossover(parent1: list, parent2: list, mask: list) -> tuple[list, list]:
    lenght = len(parent1)
    offspring1 = [0] * lenght
    offspring2 = [0] * lenght

    for i in range(lenght):
        if mask[i] == 1:
            offspring1[i] = parent2[i]
            offspring2[i] = parent1[i]
        else:
            offspring1[i] = parent1[i]
            offspring2[i] = parent2[i]
    return offspring1, offspring2


if __name__ == "__main__":
    parent1 = [1, 2, 3, 4, 5]
    parent2 = [6, 7, 8, 9, 10]
    crossover_point = 2

    offspring1, offspring2 = single_point_crossover(parent1, parent2, crossover_point)

    print(f"Parent 1: {parent1}")
    print(f"Parent 2: {parent2}")
    print(f"Offspring 1: {offspring1}")
    print(f"Offspring 2: {offspring2}")

    point1, point2 = 1, 4
    offspring1, offspring2 = two_point_crossover(parent1, parent2, point1, point2)
    print(f"\nTwo-Point Crossover between points {point1} and {point2}:")
    print(f"Offspring 1: {offspring1}")
    print(f"Offspring 2: {offspring2}")

    mask = [0, 1, 0, 1, 0]
    offspring1, offspring2 = uniform_crossover(parent1, parent2, mask)
    print(f"\nUniform Crossover with mask {mask}:")
    print(f"Offspring 1: {offspring1}")
    print(f"Offspring 2: {offspring2}")