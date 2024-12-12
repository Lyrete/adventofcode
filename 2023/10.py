PIPES = {
    "|": {(0, 1): (0, 1), (0, -1): (0, -1)},
    "-": {(1, 0): (1, 0), (-1, 0): (-1, 0)},
    "L": {(0, 1): (1, 0), (-1, 0): (0, -1)},
    "J": {(0, 1): (-1, 0), (1, 0): (0, -1)},
    "7": {(0, -1): (-1, 0), (1, 0): (0, 1)},
    "F": {(0, -1): (1, 0), (-1, 0): (0, 1)},
}


def parse(s: str) -> tuple[dict[tuple[int, int], str], tuple[int, int]]:
    points = {}
    start = (-1, -1)
    for y, line in enumerate(s.strip().splitlines()):
        for x, c in enumerate(line):
            if c != ".":
                points[x, y] = c

            if c == "S":
                start = (x, y)

    return points, start


def find_path(start: tuple[int, int], points: dict[tuple[int, int], str], path: list = []):
    x, y = start
    c = points[(x, y)]

    while c != "S":
        x0, y0 = path[-1]
        path.append((x, y))
        dx, dy = PIPES[c][x - x0, y - y0]

        x, y = x + dx, y + dy
        c = points[(x, y)]

    return path


def calculate_inside(path: list[tuple[int, int]]) -> int:
    area = 0
    perimeter = len(path)
    # Add the start point to the end to help with the loop
    path = path + [path[0]]

    # Shoelace formula
    for i in range(perimeter):
        x1, y1 = path[i]
        x2, y2 = path[i + 1]

        area += x1 * y2 - x2 * y1

    # Area is negative if we traversed the path the wrong way
    area = abs(area // 2)

    # Correct the area with Pick's theorem
    # A = i + b/2 - 1 <=> i = A - b/2 + 1
    return int(area - perimeter / 2 + 1)


def solve(s: str) -> tuple[int, int]:
    points, start = parse(s)
    x, y = start
    deltas = [(1, 0), (-1, 0), (0, 1), (0, -1)]
    filtered = [(x + dx, y + dy)
                for dx, dy in deltas if (dx, dy) in PIPES[points[x + dx, y + dy]].keys()]

    path = find_path(filtered[0], points, [start])
    inside = calculate_inside(path)

    return len(path) // 2, inside


input_data = open("data/10.in").read()

s1, s2 = solve(input_data)

print("Part 1:")
print(s1)

print("Part 2:")
print(s2)
