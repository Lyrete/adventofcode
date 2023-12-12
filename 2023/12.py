
import functools


def parse(s: str):
    lines = s.strip().splitlines()
    for line in lines:
        halves = line.split()
        sizes = [int(e) for e in halves[1].split(",")]
        sizes2 = sizes * 5
        string2 = "?".join([halves[0]] * 5)
        string2 = "".join([e for i, e in enumerate(string2)
                          if i == len(string2) - 1 or not (e == "." and string2[i+1] == ".")])

        yield solve_nonogram_row(halves[0], tuple(sizes), 0), solve_nonogram_row(string2, tuple(sizes2), 0)


@functools.cache
def solve_nonogram_row(s: str, sizes: tuple[int, ...], current_size: int) -> int:
    if len(s) == 0:
        if len(sizes) == 0 and current_size == 0:
            return 1
        elif len(sizes) == 1 and current_size == sizes[0]:
            return 1
        else:
            return 0

    if len(sizes) > 0 and current_size > sizes[0]:
        return 0
    elif len(sizes) == 0 and current_size != 0:
        return 0

    possibilities = 0
    if s[0] in ["?", "#"]:
        possibilities += solve_nonogram_row(s[1:], sizes, current_size + 1)

    if s[0] in ["?", "."]:
        if len(sizes) > 0 and sizes[0] == current_size:
            possibilities += solve_nonogram_row(s[1:], sizes[1:], 0)
        elif current_size == 0:
            possibilities += solve_nonogram_row(s[1:], sizes, 0)

    return possibilities


def solve(input_string: str):
    res = parse(input_string)
    s1, s2 = 0, 0
    for a, b in res:
        s1 += a
        s2 += b

    return s1, s2


input_data = open("data/12.in").read()

s1, s2 = solve(input_data)
print("Part 1:")
print(s1)

print("Part 2:")
print(s2)
