import os

floor = 0
char_pos = 0

below_zero = 0

with open('data.txt') as f:
    while True:
        char = f.read(1)
        if not char:
            break
        char_pos += 1
        if char == ')':
            floor -= 1
        if char == '(':
            floor += 1

        if floor < 0 and below_zero == 0:
            below_zero = char_pos

print('Floor ended on:')
print(floor)

print('Basement entered on:')
print(below_zero)