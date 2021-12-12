def print_grid(grid):
    for y in range(len(grid)):
        row = grid[y]
        for x in range(len(row)):
            print(row[x], end='')
        print()

def flash(x, y, grid):
    if x < 0 or y < 0 or y >= len(grid) or x >= len(grid[0]): #Check for invalid x and y
        return False

    grid[y][x] += 1
    
    if grid[y][x] != 10:
        return False     

    #If the value is now 9 we recursively check every adjacent octopi
    flash(x-1, y-1, grid)
    flash(x, y-1, grid)
    flash(x+1, y-1, grid)
    flash(x-1, y, grid)   
    flash(x+1, y, grid)
    flash(x-1, y+1, grid)
    flash(x, y+1, grid)
    flash(x+1, y+1, grid)

def increment(grid):
    for y in range(len(grid)):
        row = grid[y]
        for x in range(len(row)):
            flash(x, y, grid)

def remove_overflow(grid):
    count = 0
    for y in range(len(grid)):
        row = grid[y]
        for x in range(len(row)):
            if grid[y][x] > 9:
                grid[y][x] = 0
                count += 1

    return count

octopi = []

with open('data.txt') as f:
    for line in f:
        octopi.append([int(n) for n in list(line.strip())]) #split the line into a list of ints

flashes = 0
for i in range(1, 100):
    increment(octopi)
    newflashes = remove_overflow(octopi)    
    flashes += newflashes

print('Task 1, octopi flashed this many times:')
print(flashes)

print('Task 2, octopi all flashed simultaneously at step number:')
for i in range(100,10000):
    increment(octopi)
    newflashes = remove_overflow(octopi)
    if len(octopi) * len(octopi[0]) == newflashes:
        print(i)
        break








