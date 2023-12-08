# Path: 2023/9.py
from functools import reduce
import math


def parse(s: str):
    lines = s.splitlines()
    path = lines[0]
    rules = {}
    starts = ["AAA"]

    for line in lines[2:]:
        origin = line[0:3]

        if origin.endswith("A"):
            starts.append(origin)

        left = line[7:10]
        right = line[12:15]
        rules[origin] = (left, right)

    return path, rules, starts


def lcm(nums):
    return reduce(lambda a, b: a * b // math.gcd(a, b), nums)


def find_end(start: str, rules: dict[str, tuple[str, str]], path: str) -> int:
    curr = start
    traveled = 0
    while True:

        for direction in path:
            if curr[2] == "Z" or (start == "AAA" and curr == "ZZZ"):
                return traveled

            l, r = rules[curr]

            match direction:
                case "L":
                    curr = l
                case "R":
                    curr = r

            traveled += 1


def solve(s: str) -> int:
    path, rules, starts = parse(s)

    encounter_end = []

    for start in starts:
        encounter_end.append(find_end(start, rules, path))

    return encounter_end[0], lcm(encounter_end)


input_data = open("data/8.in").read()

s1, s2 = solve(input_data)

print("Part 1:")
print(s1)

print("Part 2:")
print(s2)
