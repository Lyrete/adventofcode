with open('data.txt') as f:
    for line in f:
        coord = line[13:]
        x_str, y_str = coord.split(', ')
        x_ends, y_ends = x_str[2:].split('..'), y_str[2:].split('..')
        x_range = list(range(int(x_ends[0]), int(x_ends[1]) + 1))
        y_range = list(range(int(y_ends[0]), int(y_ends[1]) + 1))

def check_hit(v_0):
    x, y = 0, 0
    v_x, v_y = v_0
    max_y = 0    

    while x not in x_range or y not in y_range:
        x += v_x
        y += v_y

        #Check if we've passed the target area
        if all(i < x for i in x_range) or all(i > y for i in y_range):
            return False, max_y, v_0
        
        if y > max_y:
            max_y = y
        
        # Apply drag
        if v_x > 0:
            v_x -= 1
        elif v_x < 0:
            v_x += 1

        # Apply gravity
        v_y -= 1

    return True, max_y, v_0

max_found = 0
found_hits = []
for x in range(0,x_range[-1] + 1):
    for y in range(y_range[0],200):
        hit, max_y, v_0 = check_hit((x,y))
        if hit:
            if max_found < max_y:
                max_found = max_y
            found_hits.append(v_0)

print('Task 1:')
print(max_found)
print('Task 2:')
print(len(found_hits))







