from typing import DefaultDict

x_coord = 0
y_coord = 0

x_coord2 = 0
y_coord2 = 0

tiles = DefaultDict(int)

tiles[(x_coord, y_coord)] += 1

with open('data.txt') as f:
    while True:
        char = f.read(1)
        if not char:
            break

        if char == '^':
            y_coord += 1
        
        if char == '>':
            x_coord += 1

        if char == '<':
            x_coord -= 1

        if char == 'v':
            y_coord -= 1

        tiles[(x_coord, y_coord)] += 1

        char = f.read(1)
        if not char:
            break

        if char == '^':
            y_coord2 += 1
        
        if char == '>':
            x_coord2 += 1

        if char == '<':
            x_coord2 -= 1

        if char == 'v':
            y_coord2 -= 1

        tiles[(x_coord2, y_coord2)] += 1

print(len(tiles))