rows = []

with open('data_ex.txt') as f:
    for line in f:
        rows.append(eval(line.strip()))

for idx, val in enumerate(rows):
    if idx == 0:
        continue

    rows[idx] = [rows[idx-1], val]
    print(rows[idx])