example = """
rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7
"""


def get_hash(s: str):
    value = 0
    for c in s:
        value += ord(c)
        value = 17 * value
        value = value % 256

    return value


def get_hash_codes(s: str):
    vals = s.strip().split(",")
    score = [get_hash(s) for s in vals]
    return sum(score)


def get_hashmap(s: str):
    boxes: dict[int, dict[str, int]] = {}
    vals = s.strip().split(",")
    for val in vals:
        label = ""
        while val[0] not in ["-", "="]:
            label += val[0]
            val = val[1:]
        box = get_hash(label)
        sym = val[0]

        if box not in boxes:
            boxes[box] = {}

        if sym == "=":
            foc = int(val[1:])
            boxes[box][label] = foc
        elif label in boxes[box]:
            boxes[box].pop(label)

    value = 0
    for b, vals in boxes.items():
        for i, foc in enumerate(vals.values()):
            curr = (b + 1) * (i + 1) * foc
            value += curr

    return value


def solve(s: str) -> tuple[int, int]:
    s1 = get_hash_codes(s)
    s2 = get_hashmap(s)

    return s1, s2


input_data = open("data/15.in").read()

s1, s2 = solve(input_data)
print("Part 1:")
print(s1)
print("Part 2:")
print(s2)
