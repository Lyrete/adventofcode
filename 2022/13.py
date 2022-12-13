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
    

lists = list(map(eval, open('inputs/data13.txt').read().strip().split()))

correct = 0

for i in range(1, len(lists), 2):
    l = lists[i-1]
    r = lists[i]
    if compare(l, r) < 1:
        correct += 1 + i // 2 

lists = lists + [[[2]]] + [[[6]]]

lists.sort(key=cmp_to_key(compare))

print("Task 1:", correct)
print("Task 2:", (lists.index([[2]]) + 1) * (lists.index([[6]]) + 1))
