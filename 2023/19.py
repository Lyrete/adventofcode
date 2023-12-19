import copy
from typing import Self


example = """
px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}
"""


class Part:
    x: int | tuple[int, int]
    m: int | tuple[int, int]
    a: int | tuple[int, int]
    s: int | tuple[int, int]

    def __init__(
        self,
        x: int | tuple[int, int] = None,
        m: int | tuple[int, int] = None,
        a: int | tuple[int, int] = None,
        s: int | tuple[int, int] = None
    ) -> None:
        self.x = x
        self.m = m
        self.a = a
        self.s = s

    def set_value(self, attr: str, val: int | tuple[int, int]) -> None:
        setattr(self, attr, val)

    def compare_value(self, attr: str, compare_to: int, op: str) -> bool:
        match op:
            case "<": return getattr(self, attr) < compare_to
            case ">": getattr(self, attr) > compare_to

    def split_on_op(self, attr: str, middle: int, op: str):
        c = getattr(self, attr)
        match op:
            case "<":
                if middle <= c[0]:
                    return None
                elif middle < c[1]:
                    new = copy.deepcopy(self)
                    setattr(new, attr, (c[0], middle - 1))
                    setattr(self, attr, (middle, c[1]))
                    return new
            case ">":
                if middle > c[1]:
                    return None
                elif middle > c[0]:
                    new = copy.deepcopy(self)
                    setattr(new, attr, (middle + 1, c[1]))
                    setattr(self, attr, (c[0], middle))
                    return new

    def __str__(self) -> str:
        return f"Part, x: {self.x}, m: {self.m}, a: {self.a}, s: {self.s}"

    def __repr__(self) -> str:
        return f"Part, x: {self.x}, m: {self.m}, a: {self.a}, s: {self.s}"

    def get_sum(self) -> int:
        return self.x + self.m + self.a + self.s

    def get_range_sum(self) -> int:
        return (self.x[1] - self.x[0] + 1) * (self.m[1] - self.m[0] + 1) * (self.a[1] - self.a[0] + 1) * (self.s[1] - self.s[0] + 1)


def parse(s: str) -> tuple[list, list[Part]]:
    s = s.strip()
    workflow, p = s.split("\n\n")
    flows = {}
    for line in workflow.splitlines():
        name, r = line[:-1].split("{")
        r = r.split(",")
        final = r[-1]

        rules = []
        for e in r[:-1]:
            rule, target = e.split(":")
            rules.append((rule[0], rule[1], int(rule[2:]), target))

        flows[name] = (rules, final)

    parts = []
    for line in p.splitlines():
        vals = line[1:-1].split(",")
        part = Part()
        for e in vals:
            attr, val = e.split("=")
            val = int(val)
            part.set_value(attr, val)
        parts.append(part)

    return flows, parts


def solve2(rules: dict[str, tuple[str, str, int, str]]) -> int:
    ranges = [("in", Part(x=(1, 4000), s=(1, 4000), a=(1, 4000), m=(1, 4000)))]

    amount = 0

    while len(ranges) > 0:
        curr, part = ranges.pop()

        if curr == "R":
            continue

        if curr == "A":
            amount += part.get_range_sum()
            continue

        r, end = rules[curr]
        for attr, op, val, target in r:
            new = part.split_on_op(attr, val, op)
            if new is not None:
                ranges.append((target, new))

        ranges.append((end, part))

    return amount


def solve(s: str) -> tuple[int, int]:
    f, p = parse(s)
    s1 = 0
    for part in p:

        curr = "in"
        while curr not in ("A", "R"):
            prev = curr
            rules, final = f[curr]
            for attr, op, val, target in rules:
                if part.compare_value(attr, val, op):
                    curr = target
                    break

            if prev == curr:
                curr = final

        s1 += part.get_sum() if curr == "A" else 0

    return s1, solve2(f)


input_data = open("data/19.in").read()

s1, s2 = solve(input_data)
print("Part 1:")
print(s1)
print("Part 2:")
print(s2)
