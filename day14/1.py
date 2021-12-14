import re
from collections import defaultdict

def def_value():
    return 0

pairs = defaultdict(def_value)
letters = defaultdict(def_value)
instructions = {}

with open('data.txt') as f:    
    for line in f:
        if len(line.strip()) <= 0: #Skip the empty line
            continue
        
        ins = line.strip().split(' -> ')

        if len(ins) > 1:
            instructions[ins[0]] = (ins[0][0] + ins[1], ins[1] + ins[0][1])      
        else:
            for letter in ins[0]:
                letters[letter] += 1
            split = re.findall('..', ins[0] + ins[0][1:])
            for pair in split:
                pairs[pair] += 1
            
for i in range(0, 10):
    newpairs = defaultdict(def_value)
    for key, value in pairs.items():
        pair1, pair2 = instructions[key] #Fetch the new pairs
        newpairs[pair1] += 1 * value
        newpairs[pair2] += 1 * value
        letters[pair1[1]] += 1 * value #Only letter to increase is the middle one
    pairs = newpairs

print('Task 1:')
print(max(letters.values()) - min(letters.values()))

for i in range(10, 40):
    newpairs = defaultdict(def_value)
    for key, value in pairs.items():
        pair1, pair2 = instructions[key] #Fetch the new pairs
        newpairs[pair1] += 1 * value
        newpairs[pair2] += 1 * value
        letters[pair1[1]] += 1 * value #Only letter to increase is the middle one
    pairs = newpairs

print('Task 2:')
print(max(letters.values()) - min(letters.values()))






