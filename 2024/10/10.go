package main

import (
	"aoc"
	"fmt"
	"strings"
)

func solve(input string) (int, int) {
	grid, startLocs, limits := parseGrid(input)
	res := 0
	res2 := 0
	for loc := range startLocs {
		visitedFromStart := make(map[aoc.Coord]struct{})
		res += countPossibilities(grid, loc, -1, limits, visitedFromStart, true)
		res2 += countPossibilities(grid, loc, -1, limits, make(map[aoc.Coord]struct{}), false)
	}
	return res, res2
}

func countPossibilities(grid [][]int, curr aoc.Coord, prev int, limits []int, visited map[aoc.Coord]struct{}, countVisited bool) int {
	if _, ok := visited[curr]; ok {
		return 0
	}
	if curr.Y < 0 || curr.X < 0 || curr.Y > limits[1] || curr.X > limits[0] {
		return 0
	}
	handleable := grid[curr.Y][curr.X]
	if (handleable - prev) != 1 {
		return 0
	}
	if countVisited {
		visited[curr] = struct{}{}
	}
	if handleable == 9 {
		return 1
	}

	return countPossibilities(grid, curr.AddXY(1, 0), handleable, limits, visited, countVisited) + countPossibilities(grid, curr.AddXY(0, 1), handleable, limits, visited, countVisited) + countPossibilities(grid, curr.AddXY(0, -1), handleable, limits, visited, countVisited) + countPossibilities(grid, curr.AddXY(-1, 0), handleable, limits, visited, countVisited)
}

func parseGrid(input string) ([][]int, map[aoc.Coord]struct{}, []int) {
	rows := [][]int{}
	starts := make(map[aoc.Coord]struct{})
	limits := make([]int, 2)
	for y, line := range strings.Split(input, "\n") {
		col := make([]int, len(line))
		for x, r := range line {
			col[x] = int(r - '0')
			if r == '0' {
				starts[aoc.Coord{X: x, Y: y}] = struct{}{}
			}
			limits[0] = x
		}
		rows = append(rows, col)
		limits[1] = y
	}
	return rows, starts, limits
}

func main() {
	fmt.Println("Example result:")
	fmt.Println(solve(example))

	fmt.Println("Real:")
	fmt.Println(solve(aoc.GetInputFromFile("10")))
}

const example = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`
