package main

import (
	"aoc"
	"container/heap"
	"fmt"
	"math"
	"strings"
)

type direction struct {
	dx int
	dy int
}

func (d *direction) getLeft() direction {
	return direction{dx: -d.dy, dy: d.dx}
}

func (d *direction) getRight() direction {
	return direction{dx: d.dy, dy: -d.dx}
}

type finder struct {
	pos      aoc.Coord
	dir      direction
	pathCost int
	path     []aoc.Coord
	priority int
	index    int
}

func (f *finder) getNext() finder {
	nextPos := f.pos.AddXY(f.dir.dx, f.dir.dy)
	newPath := make([]aoc.Coord, len(f.path))
	copy(newPath, f.path)
	newPath = append(newPath, nextPos)
	new := finder{pos: nextPos, dir: f.dir, pathCost: f.pathCost + 1, path: newPath}
	return new
}

func (f *finder) getLeftTurn() finder {
	new := finder{pos: f.pos, dir: f.dir.getLeft(), pathCost: f.pathCost + 1000, path: f.path}
	return new
}

func (f *finder) getRightTurn() finder {
	new := finder{pos: f.pos, dir: f.dir.getRight(), pathCost: f.pathCost + 1000, path: f.path}
	return new
}

func parse(input string) (walls map[aoc.Coord]struct{}, start aoc.Coord, end aoc.Coord) {
	split := strings.Split(input, "\n")
	walls = make(map[aoc.Coord]struct{})
	for y, line := range split {
		for x, r := range line {
			if r == '#' {
				walls[aoc.Coord{X: x, Y: len(split) - y - 1}] = struct{}{}
			}
			if r == 'S' {
				start = aoc.Coord{X: x, Y: len(split) - y - 1}
			}
			if r == 'E' {
				end = aoc.Coord{X: x, Y: len(split) - y - 1}
			}
		}
	}
	return
}

func addToMap(tiles map[int]map[aoc.Coord]struct{}, key int, path []aoc.Coord) map[int]map[aoc.Coord]struct{} {
	for _, e := range path {
		if _, ok := tiles[key]; !ok {
			tiles[key] = make(map[aoc.Coord]struct{})
		}
		tiles[key][e] = struct{}{}
	}
	return tiles
}

type node struct {
	pos aoc.Coord
	dir direction
}

func findPath(walls map[aoc.Coord]struct{}, start aoc.Coord, end aoc.Coord) (int, int) {
	pq := make(PriorityQueue, 1)

	pq[0] = &finder{pos: start, dir: direction{1, 0}, priority: start.ManhattanDistance(end), index: 0, pathCost: 0, path: []aoc.Coord{start}}

	visited := make(map[node]int)
	pathUniques := make(map[int]map[aoc.Coord]struct{})
	best := math.MaxInt

	heap.Init(&pq)

	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*finder)
		if _, ok := walls[curr.pos]; ok {
			continue
		}

		if v, ok := visited[node{pos: curr.pos, dir: curr.dir}]; ok && v < curr.pathCost {
			continue
		}

		visited[node{pos: curr.pos, dir: curr.dir}] = curr.pathCost

		if curr.pos == end {
			//fmt.Println(len(curr.path))
			//fmt.Println(curr.pathCost)
			if best > curr.pathCost {
				best = curr.pathCost
			}
			pathUniques = addToMap(pathUniques, curr.pathCost, curr.path)
			//fmt.Println(len(pathUniques[curr.pathCost]))
			continue
		}

		next := curr.getNext()
		if _, ok := walls[next.pos]; ok {
			visited[node{pos: curr.pos, dir: curr.dir}] += 1000
		} else {
			next.priority = len(next.path)
			heap.Push(&pq, &next)
		}

		right := curr.getRightTurn()
		right = right.getNext()
		right.priority = len(right.path)
		heap.Push(&pq, &right)

		left := curr.getLeftTurn()
		left = left.getNext()
		left.priority = len(left.path)
		heap.Push(&pq, &left)
	}

	marks := pathUniques[best]

	//printMap(walls, marks, 141, 141)

	return best, len(marks)
}

func printMap(walls map[aoc.Coord]struct{}, visited map[aoc.Coord]struct{}, x_size, y_size int) {
	for y := range y_size {
		for x := range x_size {
			if _, ok := walls[aoc.Coord{X: x, Y: y_size - y - 1}]; ok {
				fmt.Print("#")
			} else if _, ok := visited[aoc.Coord{X: x, Y: y_size - y - 1}]; ok {
				fmt.Print("O")
			} else if x == 15 && y_size-8-1 == y {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func solve(input string) (int, int) {
	walls, start, end := parse(input)
	res, res2 := findPath(walls, start, end)
	return res, res2
}

func main() {
	fmt.Println("Example result:")
	fmt.Println(solve(example))

	fmt.Println("Real:")
	fmt.Println(solve(aoc.GetInputFromFile("16")))
}

const example = `###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`

type PriorityQueue []*finder

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*finder)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *finder, value aoc.Coord, priority int) {
	item.pos = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
