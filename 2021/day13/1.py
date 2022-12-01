points = set()
ins = []

#Helper for visualizing
def print_grid(values):
    grid = [[' '] * 80] * 8


    for y in range(len(grid)):
        for x in range(len(grid[y])):
            if (x,y) in values:
                print('█', end='')
            else:
                print(grid[y][x], end='')
        print()
                 
#Read file
with open('data.txt') as f:
    coord_read = False
    for line in f:
        if len(line.strip()) == 0:
            coord_read = True
            continue

        if not coord_read:
            x,y = line.strip().split(',')
            points.add((int(x),int(y)))
        else:
            axis, coord = line.strip().split()[2].split('=')
            ins.append((axis, int(coord)))

#Go through the folds
idx = 1
for i in ins:
    newset = set()
    axis, coord = i

    while len(points) > 0:
        x, y = points.pop()
        if axis == 'x':
            dist = x - coord
            if dist > 0:
                x, y = coord - dist, y            
        else:
            dist = y - coord
            if dist > 0:
                x, y = x, coord - dist
        
        newset.add((x,y))

    points = newset
    print(f'Fold {idx}: {len(points)}')

    idx += 1

print_grid(points)




      

        


