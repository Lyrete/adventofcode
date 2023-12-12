def parse(s: str):
    lines = s.strip().splitlines()
    lines = [line.split() for line in lines]
    return lines


def solve(s: str) -> int:
    lines = parse(s)

    result, result2 = 0, 0

    for line in lines:
        line = [int(x) for x in line]

        last = [line[-1]]
        first = [line[0]]

        while any(line):
            for i in range(len(line) - 1):
                line[i] = line[i + 1] - line[i]

            line.pop()  # We never update the last value so pop it (as there's nothing to compare to)
            first.append(line[0])
            last.append(line[-1])

        for i in reversed(range(1, len(last))):
            last[i-1] = last[i-1] + last[i]
            first[i-1] = first[i-1] - first[i]

        result2 += first[0]
        result += last[0]

    return result, result2


input_data = open("data/9.in").read()

s1, s2 = solve(input_data)

print(s1)
print(s2)
