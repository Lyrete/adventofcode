import importlib

lib = importlib.import_module('1')

def new_value(value):
    if value + 1 > 9:
        return 1
    else:
        return value + 1

def main():
    graph = {}
    with open('data.txt') as f:
        for y, line in enumerate(f):
                for x, cost in enumerate(line.strip()):
                    graph[(x,y)] = int(cost)

    add = max(x for x,y in graph.keys()) +1

    prev_y_index = 0
    for rep in range(1,5):
        new = {(key[0],key[1]+add):new_value(cost) for key,cost in graph.items() if key[1]>=prev_y_index}
        graph.update(new)
        prev_y_index += add

    prev_x_index = 0
    for rep in range(1,5):
        new = {(key[0]+add,key[1]):new_value(cost) for key,cost in graph.items() if key[0]>=prev_x_index}
        graph.update(new)
        prev_x_index += add

    start = (0,0)
    end = (max(x for x,y in graph.keys()),max(y for x,y in graph.keys()))
    
    print('Task 2:')
    print(lib.dijkstra(graph, start, end)[end])

if __name__ == '__main__':
    main()