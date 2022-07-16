total_size = 0
total_ribbon = 0

with open('data.txt') as f:
    for line in f:
        sides = [int(x) for x in line.strip().split('x')]
        rects = [
            sides[0] * sides[1],
            sides[1] * sides[2],
            sides[0] * sides[2]
            ]
        sides.sort()
        rects.sort()
        size = sum(rects) * 2 + rects[0]
        ribbon = sides[0] * 2 + sides[1] * 2 + sides[0] * sides[1] * sides[2]
        total_size += size
        total_ribbon += ribbon

print('Total size:')
print(total_size)

print('Total ribbon:')
print(total_ribbon)