example = """

"""


def parse(s: str):
    pass


def solve(s: str) -> tuple[int, int]:
    parse(s)

    return 0, 0


ex1, ex2 = solve(example)

print(ex1)


input_data = open("data/14.in").read()

s1, s2 = solve(input_data)
print("Part 1:")
print(s1)
print("Part 2:")
print(s2)
