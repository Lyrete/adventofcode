legalpairs = {'(':')', '[':']', '<':'>', '{':'}'}

openers = {'(', '[', '<', '{'}
closers = {')', ']', '>', '}'}

point_values = {')':3, ']':57, '}':1197, '>':25137}
completion_values = {')':1, ']':2, '}':3, '>':4}

corrupt_points = 0
completion_points = []

with open('data.txt') as f:
    for line in [line.strip() for line in f]:
        opened = []
        completion_points_current = 0
        for x in line:            
            if x in openers:
                opened.append(x)

            if x in closers:                
                expection = legalpairs[opened.pop()] #We're expecting the last char to be added to openers to have a match             
                if expection != x:
                    #print(f'{line} - Expected {expection}, but found {x} instead.')
                    corrupt_points += point_values[x]
                    opened = [] #Clear the opened list so we don't also get completion points
                    break

        if len(opened) > 0:
            opened.reverse() #reverse the list so we go through the symbols in correct order
            for x in opened:
                completion_points_current = completion_points_current * 5 + completion_values[legalpairs[x]]
            completion_points.append(completion_points_current)

print('Task 1, corrupt lines got this many points:')
print(corrupt_points)
completion_points.sort()
print('Task 2, completing lines got this many points:')
print(completion_points[len(completion_points)//2])
        
        