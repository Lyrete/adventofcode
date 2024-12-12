package main

import (
	"aoc"
	"fmt"
	"strings"
)

func parseInputIntoGrid(input string) map[point]string {
	grid := make(map[point]string)
	for y, line := range strings.Split(input, "\n") {
		for x, s := range line {
			grid[point{x, y}] = string(s)
		}
	}
	return grid
}

type point struct {
	x int
	y int
}

func (p point) right() point {
	return point{p.x + 1, p.y}
}

func (p point) left() point {
	return point{p.x - 1, p.y}
}

func (p point) top() point {
	return point{p.x, p.y - 1}
}

func (p point) bot() point {
	return point{p.x, p.y + 1}
}

func (p point) topright() point {
	return point{p.x + 1, p.y - 1}
}

func (p point) topleft() point {
	return point{p.x - 1, p.y - 1}
}

func (p point) botright() point {
	return point{p.x + 1, p.y + 1}
}

func (p point) botleft() point {
	return point{p.x - 1, p.y + 1}
}

func countCornersAndPerimeter(plot map[point]struct{}) (int, int) {
	cornerCount := 0
	p := 0
	for b := range plot {
		top, right, left, bot := b.top(), b.right(), b.left(), b.bot()
		if !aoc.HasKey(plot, top) {
			p++
		}
		if !aoc.HasKey(plot, left) {
			p++
		}
		if !aoc.HasKey(plot, right) {
			p++
		}
		if !aoc.HasKey(plot, bot) {
			p++
		}

		if !aoc.HasKey(plot, top) && !aoc.HasKey(plot, left) {
			cornerCount++
		}
		if !aoc.HasKey(plot, top) && !aoc.HasKey(plot, right) {
			cornerCount++
		}
		if !aoc.HasKey(plot, bot) && !aoc.HasKey(plot, right) {
			cornerCount++
		}
		if !aoc.HasKey(plot, bot) && !aoc.HasKey(plot, left) {
			cornerCount++
		}

		// Inside corners
		if aoc.HasKey(plot, top) && aoc.HasKey(plot, left) && !aoc.HasKey(plot, b.topleft()) {
			cornerCount++
		}
		if aoc.HasKey(plot, top) && aoc.HasKey(plot, right) && !aoc.HasKey(plot, b.topright()) {
			cornerCount++
		}
		if aoc.HasKey(plot, bot) && aoc.HasKey(plot, right) && !aoc.HasKey(plot, b.botright()) {
			cornerCount++
		}
		if aoc.HasKey(plot, bot) && aoc.HasKey(plot, left) && !aoc.HasKey(plot, b.botleft()) {
			cornerCount++
		}
	}

	return cornerCount, p
}

func calculatePerimeterAndArea(plot map[point]struct{}) (int, int, int) {
	A := len(plot)
	corners, p := countCornersAndPerimeter(plot)
	return p, A, corners
}

func findEnclosedArea(grid map[point]string, visited map[point]struct{}, start point) map[point]struct{} {
	currentArea := make(map[point]struct{})
	queue := []point{start}
	for len(queue) > 0 {
		checkablePoint := queue[0]
		queue = queue[1:]
		currentArea[checkablePoint] = struct{}{}
		visited[checkablePoint] = struct{}{}

		for _, neighbour := range []point{checkablePoint.bot(), checkablePoint.right(), checkablePoint.left(), checkablePoint.top()} {
			if s, ok := grid[neighbour]; ok && s == grid[checkablePoint] {
				if _, ok := visited[neighbour]; !ok {
					queue = append(queue, neighbour)
				}
			}
		}
	}

	return currentArea
}

func solve(input string) (int, int) {

	grid := parseInputIntoGrid(input)
	visited := make(map[point]struct{})
	res, res2 := 0, 0

	for point := range grid {
		if _, ok := visited[point]; ok {
			continue
		}
		p, A, sides := calculatePerimeterAndArea(findEnclosedArea(grid, visited, point))
		//fmt.Println(grid[point], ":", A, "*", sides, "=", A*sides)
		// fmt.Println()
		res += p * A
		res2 += A * sides

	}

	return res, res2
}

func main() {
	fmt.Println("Example result:")
	fmt.Println(solve(example))

	fmt.Println("Real:")
	fmt.Println(solve(aoc.GetInputFromFile("12")))
}

const example = `
RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`
