example = """
#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#"""

# Lazy reverse, cba using a stack


def transpose(lines: list[str]) -> list[str]:
    t_lines = []
    for i in range(len(lines[0])):
        t_lines.append(''.join([s[i] for s in lines]))

    return t_lines


def reverse(s: str):
    s = list(s)
    s.reverse()
    return "".join(s)


def check_symmetry(left: str, right: str) -> bool:
    rev_left = reverse(left)
    if len(left) <= len(right):
        return right.startswith(rev_left)
    else:
        return rev_left.startswith(right)


def find_symmetrical(block: str) -> list[int]:
    lines = block.split()
    w = len(lines[0])
    h = len(lines)

    valids = []

    x = 1
    while x < w:
        if all([check_symmetry(line[:x], line[x:]) for line in lines]):
            # return x, x+1
            valids.append(x)
        x += 1

    t_lines = transpose(lines)

    y = 1
    while y < h:
        if all([check_symmetry(line[:y], line[y:]) for line in t_lines]):
            # return y, y+1
            valids.append(y * 100)
        y += 1

    return valids


def find_smudge(block: str, prev_result: int) -> int:
    for i in range(len(block)):
        if block[i] == "\n":
            continue

        l, r = block[:i], block[i+1:]
        c = block[i]

        if c == "#":
            potential = find_symmetrical(l + "." + r)
        else:
            potential = find_symmetrical(l + "#" + r)

        if len(potential) > 0 and (prev_result not in potential or len(potential) > 1):
            return [p for p in potential if p != prev_result][0]

    return 0


def parse(s: str):
    blocks = s.strip().split("\n\n")
    for block in blocks:
        p1 = find_symmetrical(block)
        p2 = find_smudge(block, p1[0])
        yield p1[0], p2


def solve(s: str) -> tuple[int, int]:
    res = parse(s)
    r1 = 0
    r2 = 0
    for p1, p2 in res:
        r1 += p1
        r2 += p2

    return r1, r2


input_data = open("data/13.in").read()

s1, s2 = solve(input_data)
print("Part 1:")
print(s1)

print("Part 2:")
print(s2)
