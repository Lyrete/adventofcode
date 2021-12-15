from queue import PriorityQueue

def adjacent(node, graph):
    x0, y0 = node
    possibilities = [(x0-1, y0),(x0+1,y0), (x0,y0-1), (x0, y0+1)]
    return [(x,y) for (x,y) in possibilities if (x,y) in graph]

def dist_to_end(start, end):
    x0,y0 = start
    x1,y1 = end

    return abs(x1-x0) + abs(y1-y0)

def dijkstra(graph, start, end):
    D = {}
    D[start] = 0
    pq = PriorityQueue()
    pq.put((0, start))

    while not pq.empty():
        prio, current_node = pq.get()

        if current_node == end:
            break

        for neighbor in adjacent(current_node, graph):
            risk_level = D[current_node] + graph[neighbor]
            if neighbor not in D or risk_level < D[neighbor]:
                D[neighbor] = risk_level
                pq.put((risk_level + dist_to_end(neighbor, end), neighbor))        

    return D

def main():
    graph = {}
    with open('data.txt') as f:    
        for y, line in enumerate(f):
            for x, cost in enumerate(line.strip()):
                graph[(x,y)] = int(cost)

    start = (0,0)
    end = (max(x for x,y in graph.keys()),max(y for x,y in graph.keys()))

    print('Task 1:')
    print(dijkstra(graph, start, end)[end])

if __name__ == '__main__':
    main()

    
