caves = {}

with open('data.txt') as f:
    for line in f:
        cave1, cave2 = line.strip().split('-')
        
        if cave1 not in caves.keys():
            caves[cave1] = set()

        if cave2 not in caves.keys():
            caves[cave2] = set()
        
        caves[cave1].add(cave2)
        caves[cave2].add(cave1)

def find_all_paths(caves, start, end, path=[]):
    path = path + [start]
    if start == end:
        return [path]
    
    paths = []

    for cave in caves[start]:
        lower = [x for x in path if x.islower()]
        lower_set = set(lower)

        check_lower = len(lower) < len(lower_set) + 1 and cave != 'start' #Check whether there's only one double small cave and that we don't try to visit start twice

        if cave not in path or cave.isupper() or check_lower:
            newpaths = find_all_paths(caves, cave, end, path)
            for newpath in newpaths:
                
                paths.append(newpath)

    return paths

res = find_all_paths(caves, 'start', 'end')

print(len(res))


