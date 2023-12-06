import time


example = """seeds: 82 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4"""


def solve(s: str) -> int:
    categories = s.split("\n\n")

    values = [int(x) for x in categories.pop(0).replace("seeds: ", "").split()]

    while len(categories) > 0:
        category = categories.pop(0)
        updated = [False] * len(values)
        for line in category.split("\n")[1:]:
            dest_start, source_start, length = [int(x) for x in line.split()]

            for i in range(len(values)):
                if updated[i]:
                    continue

                curr = values[i]
                if not (curr < source_start or curr >= source_start + length):
                    values[i] = dest_start + (curr - source_start)
                    updated[i] = True

    return min(values)


def solve2(s: str) -> int:
    categories = s.split("\n\n")
    seeds = [int(x) for x in categories.pop(0).replace("seeds: ", "").split()]
    seed_ranges = [(seeds[i], seeds[i + 1]) for i in range(0, len(seeds), 2)]

    conversions = []

    while len(categories) > 0:
        category = categories.pop(0)
        new = []
        for line in category.split("\n")[1:]:
            dest_start, source_start, length = [int(x) for x in line.split()]
            new.append((dest_start, source_start, length))
        conversions.append(new)

    seed_ranges = seed_ranges

    for conv in conversions:
        seed_range_amount = len(seed_ranges)
        for i in range(seed_range_amount):
            for dest_start, source_start, conv_length in conv:
                orig_start, orig_length = seed_ranges[i]

                # Test if range fully out
                if orig_start + orig_length < source_start or orig_start >= source_start + conv_length:
                    continue

                # Range starts before conversion (so we keep the starting part)
                if orig_start < source_start:
                    new_start = orig_start
                    new_length = source_start - orig_start
                    seed_ranges.append((new_start, new_length))

                    orig_start = source_start
                    orig_length = orig_length - new_length

                # Range ends after conversion (so we keep the ending part)
                if orig_start + orig_length > source_start + conv_length:
                    rem_start = source_start + conv_length
                    rem_length = orig_length - (rem_start - orig_start)
                    seed_ranges.append((rem_start, rem_length))
                    orig_length -= rem_length

                new_start = dest_start + (orig_start - source_start)
                new_length = orig_length
                seed_ranges[i] = (new_start, new_length)
                break

    return min([x for x, y in seed_ranges])


def solveRealInput():
    in_str = open("data/5.txt").read().strip()
    s1, s2 = solve(in_str), solve2(in_str)

    return s1, s2


total_time = 0

for i in range(1000):
    time_start = time.perf_counter()
    s1, s2 = solveRealInput()
    time_end = time.perf_counter()
    total_time += time_end - time_start

print(f"Total time (1000 runs): {total_time:.4f} seconds")
print(f"Average time: {total_time / 1000:.4f} seconds")
print(f"Part 1: {s1}")
print(f"Part 2: {s2}")
