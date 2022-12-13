from pathlib import Path
from functools import cmp_to_key

def compare(a, b):
    if type(a) == int and type(b) == int:
        return a - b
    elif type(a) == list and type(b) == list:
        for x, y in zip(a,b):
            res = compare(x, y)
            if res != 0:
                return res
        return len(a) - len(b)
    elif type(a) == int:
        return compare([a], b)
    else:
        return compare(a,[b])
    

txt = Path('inputs/data13.txt').read_text().split("\n\n")

pairs = []

for chunk in txt:
    litarray = chunk.rstrip().split("\n")
    left = eval(litarray[0])
    right = eval(litarray[1])
    pairs.append((left, right))

correct = 0

for (i, pair) in enumerate(pairs):
    if compare(pair[0], pair[1]) < 1:
        correct += i + 1

flattened = [elem for tuple in pairs for elem in tuple]
flattened.append([[2]])
flattened.append([[6]])

flattened.sort(key=cmp_to_key(compare))

print("Task 1:", correct)
print("Task 2:", (flattened.index([[2]]) + 1) * (flattened.index([[6]]) + 1))
