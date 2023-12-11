example = """
...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#....."""


def parse(s: str):
    lines = s.strip().splitlines()
    size = (0, 0)
    for y, line in enumerate(lines):
        for x, c in enumerate(line):
            if c == "#":
                yield (x, y)

    if size == (0, 0):
        size = (x, y)
        yield size


def manhattan_distance(a, b):
    return abs(a[0] - b[0]) + abs(a[1] - b[1])


def calculate_distances_between(asteroids):
    distance = 0
    while len(asteroids) > 0:
        origin = asteroids.pop()
        distance += sum([manhattan_distance(origin, a) for a in asteroids])

    return distance


def solve(input_string: str):
    asteroids = list(parse(input_string))
    size = asteroids.pop()

    x_offsets = []
    for x in range(size[0]):
        if len([a for a in asteroids if a[0] == x]) == 0:
            x_offsets.append(x)

    y_offsets = []
    for y in range(size[1]):
        if len([a for a in asteroids if a[1] == y]) == 0:
            y_offsets.append(y)

    moved_asteroids = []
    for x, y in asteroids:
        moved_asteroids.append(
            (x + len([e for e in x_offsets if e < x]), y + len([e for e in y_offsets if e < y])))

    moved_part2 = []
    for x, y in asteroids:
        moved_part2.append(
            (x + (1e6 - 1) * len([e for e in x_offsets if e < x]),
             y + (1e6 - 1) * len([e for e in y_offsets if e < y])))

    return calculate_distances_between(moved_asteroids), int(calculate_distances_between(moved_part2))


input_data = open("data/11.in").read()

s1, s2 = solve(input_data)

print("Part 1:")
print(s1)
print("Part 2:")
print(s2)
