example = """467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598.."""


def check(x, y, coords) -> bool:
    for cx, cy in coords:
        dist = chebyshev((x, y), (cx, cy))
        if dist == 1:
            return True
    return False


def chebyshev(pos1: (int, int), pos2: (int, int)) -> bool:
    return max(abs(pos1[0] - pos2[0]), abs(pos1[1] - pos2[1])) == 1


def solve(input: str) -> (int, int):
    lines = input.splitlines()
    symbols = []
    gears = {}

    for y in range(0, len(lines)):
        # loop over each char in line
        for x in range(0, len(lines[y])):
            if lines[y][x] != "." and not lines[y][x].isdigit():
                symbols.append((x, y))
                if lines[y][x] == "*":
                    gears[(x, y)] = (1, 0)

    s1 = 0

    for y in range(0, len(lines)):
        x = 0
        line = lines[y]
        while x < len(line):
            char = line[x]
            current_num = ""
            validNumber = False
            gear = None
            while char.isdigit():
                if check(x, y, symbols):
                    validNumber = True

                for sym in gears.keys():
                    if chebyshev((x, y), sym):
                        gear = sym

                current_num += char
                x += 1
                if x < len(line):
                    char = line[x]
                else:
                    break

            if validNumber:
                s1 += int(current_num)
                if gear is not None:
                    temp = gears.get(gear)
                    temp = (temp[0] * int(current_num), temp[1] + 1)
                    gears[gear] = temp
            x += 1

    return s1, sum([x for x, y in gears.values() if y == 2])


print("Example (4361):")
ex1, ex2 = solve(example)
print(ex1)
print("Part 1:")
s1, s2 = solve(open("data/3.txt").read())
print(s1)
print("Example (467835):")
print(ex2)
print("Part 2:")
print(s2)
