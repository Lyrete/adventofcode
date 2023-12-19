import sys


example = """
.|...\\....
|.-.\\.....
.....|-...
........|.
..........
.........\\
..../.\\\\..
.-.-/..|..
.|....-|.\\
..//.|....
"""

sys.setrecursionlimit(4000)


def new_directions(symbol: str, direction: tuple[int, int]) -> list[tuple[int, int]]:
    dx, dy = direction
    match symbol:
        case ".":
            return [(dx, dy)]
        case "/":
            return [(-dy, -dx)]
        case "\\":
            return [(dy, dx)]
        case "-":
            if dy == 0:
                return [(dx, dy)]
            return [(1, 0), (-1, 0)]
        case "|":
            if dx == 0:
                return [(dx, dy)]
            return [(0, 1), (0, -1)]


def traverse(start: tuple[int, int], direction: tuple[int, int], board: list[list[str]], visited: dict = {}):

    # Exit out of traverse when we go outside the grid
    if start[0] < 0 or start[1] < 0 or start[1] > len(board) - 1 or start[0] > len(board[0]) - 1:
        return

    if start not in visited:
        visited[start] = [direction]
    elif direction in visited[start]:
        return

    visited[start].append(direction)
    x, y = start

    sym = board[y][x]

    for dir in new_directions(sym, direction):
        traverse((x + dir[0], y + dir[1]), dir, board, visited)

    return len(visited)


def parse(s: str) -> list[list[str]]:
    lines = s.strip().splitlines()
    board = []
    for y in range(len(lines)):
        board.append([])
        for c in lines[y]:
            board[y].append(c)

    return board


def print_grid(board: list[list[str]]):
    for line in board:
        for c in line:
            print(c, end="")
        print()


def solve(s: str) -> tuple[int, int]:
    board = parse(s)

    s1 = traverse((0, 0), (1, 0), board, {})

    s2 = s1

    for x in range(len(board[0])):
        s2 = max(s2, traverse((x, 0), (0, 1), board, {}))

    for x in range(len(board[0])):
        s2 = max(s2, traverse((x, len(board) - 1), (0, -1), board, {}))

    for y in range(len(board)):
        s2 = max(s2, traverse((0, y), (1, 0), board, {}))

    for y in range(len(board)):
        s2 = max(s2, traverse((len(board[0]) - 1, y), (-1, 0), board, {}))

    return s1, s2


ex1, ex2 = solve(example)

print(ex1)
print(ex2)

input_data = open("data/16.in").read()

s1, s2 = solve(input_data)
print("Part 1:")
print(s1)
print("Part 2:")
print(s2)
