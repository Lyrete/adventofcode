example = """
R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)
"""


def move(point: tuple[int, int], direction: str, length: int) -> tuple[int, int]:
    x, y = point
    match direction:
        case "R" | "0": return (x+length, y)
        case "L" | "2": return (x-length, y)
        case "U" | "3": return (x, y-length)
        case "D" | "1": return (x, y+length)


def parse(s: str):
    x1 = y1 = x2 = y2 = 0
    for line in s.strip().splitlines():
        d, l1, code = line.split(" ")
        code = code[1:-1]
        x1_1, y1_1 = move((x1, y1), d, int(l1))
        x2_1, y2_1 = move((x2, y2), code[-1], int(code[1:-1], 16))
        yield (x1, y1), (x2, y2)
        x1, y1 = x1_1, y1_1
        x2, y2 = x2_1, y2_1


def manhattan_distance(a, b):
    return abs(b[0] - a[0]) + abs(a[1] - b[1])


def calculate_inside(path: list[tuple[int, int]]) -> int:
    area = 0
    # Add the start point to the end to help with the loop
    path = path + [path[0]]
    # Shoelace formula
    i = 0
    for n in range(len(path) - 1):
        x1, y1 = path[n]
        x2, y2 = path[n + 1]

        i += manhattan_distance((x1, y1), (x2, y2))

        area += x1 * y2 - x2 * y1

    # Area is negative if we traversed the path the wrong way
    area = abs(area // 2)

    # Correct the area with Pick's theorem, but add the perimeter back
    # A = i + b/2 - 1 <=> i = A - b/2 + 1
    return int(area - i / 2 + 1) + i


def solve(s: str) -> tuple[int, int]:
    p1 = []
    p2 = []
    for (x1, y1), (x2, y2) in parse(s):
        p1.append((x1, y1))
        p2.append((x2, y2))
    s1 = calculate_inside(p1)
    s2 = calculate_inside(p2)

    return s1, s2


input_data = open("data/18.in").read()

s1, s2 = solve(input_data)
print("Part 1:")
print(s1)
print("Part 2:")
print(s2)
