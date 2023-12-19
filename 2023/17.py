from collections import defaultdict
import heapq
from math import inf
example = """
2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533
"""


def parse(s: str):
    lines = s.strip().splitlines()
    board = []

    for line in lines:
        new = []
        for c in line:
            new.append(int(c))
        board.append(new)

    return board


def manhattan_distance(a, b):
    return abs(a[0] - a[1]) + abs(b[0] - b[1])

# Convoluted dijkstra


def find_path(board, min_step=1, max_step=3):
    start = (0, 0)
    end = (len(board[0]) - 1, len(board) - 1)
    pq = [(0, (start, (0, 1))), (0, (start, (1, 0)))]
    heat_values = defaultdict(lambda: inf)

    while pq:
        prio, (curr, direction) = heapq.heappop(pq)

        if curr == end:
            return prio

        if prio > heat_values[curr, direction]:
            continue

        dx0, dy0 = direction
        for dx, dy in ((-dy0, dx0), (dy0, -dx0)):
            new_cost = prio
            for d in range(1, max_step + 1):
                x, y = curr[0] + dx * d, curr[1] + dy * d

                if 0 <= x < len(board[0]) and 0 <= y < len(board):
                    new_cost += board[y][x]
                    if d < min_step:
                        continue
                    key = ((x, y), (dx, dy))
                    if new_cost < heat_values[key]:
                        heat_values[key] = new_cost
                        heapq.heappush(pq, (new_cost, key))

    return -1


def print_grid(board, marks):
    for (x, y) in marks:
        board[y][x] = ">"

    for line in board:
        for c in line:
            print(c, end="")
        print()


def solve(s: str) -> tuple[int, int]:
    board = parse(s)

    heat = find_path(board)
    heat2 = find_path(board, 4, 10)

    return heat, heat2


ex1, ex2 = solve(example)

print(ex1)
print(ex2)


input_data = open("data/17.in").read()

s1, s2 = solve(input_data)
print("Part 1:")
print(s1)
print("Part 2:")
print(s2)
