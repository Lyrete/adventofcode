import numpy
from collections import deque

tubes = []

with open('data.txt') as f:
    for line in f:
        tubes.append([int(n) for n in list(line.strip())]) #split the line into a list of ints

#Helper to visualize grid's in same way as the site
def print_grid(grid):
    for y in range(len(grid)):
        row = grid[y]
        for x in range(len(row)):
            print(row[x], end='')
        print()

#Function to help check if a point is a low point (all 4 directions are bigger)
def is_low_point(x,y):
    for x1, y1 in adjacent(x,y):
        if tubes[y][x] == 9 or tubes[y1][x1] < tubes[y][x]:
            return False
    return True

def adjacent(x,y):
    points = [(x-1, y), (x+1,y), (x, y+1), (x,y-1)]
    return [(x1, y1) for (x1, y1) in points if 0 <= x1 < len(tubes[y]) and 0 <= y1 < len(tubes)]

def find_basin(x, y):
    basin = []
    visited = set()
    queue = deque([(x,y)])

    while queue:
        x1, y1 = queue.pop()

        if (x1, y1) in visited:
            continue
        else:
            visited.add((x1, y1))
            if tubes[y1][x1] != 9:
                #Append the coordinates to the basin
                basin.append((x1, y1))
                #Add all adjacent point that have not been checked yet to queue
                queue.extend([(x2, y2) for (x2, y2) in adjacent(x1,y1) if (x2,y2) not in visited])

    return basin

basins = []
for y in range(len(tubes)):
    for x in range(len(tubes[y])):
        if is_low_point(x, y):
            basins.append(find_basin(x,y))

basins.sort(key=len)

print(numpy.prod([len(basin) for basin in basins[-3:]]))





