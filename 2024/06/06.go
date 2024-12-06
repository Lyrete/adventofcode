package main

import (
	"aoc_helpers"
	"fmt"
	"slices"
	"strings"
)

type dir struct {
	dx int8
	dy int8
}

func (d *dir) turnRight() {
	tempy := d.dy
	d.dy = d.dx
	d.dx = -tempy
}

type guardPosition struct {
	x        int
	y        int
	next_dir dir
}

func (g *guardPosition) setPosXY(x int, y int) {
	g.x = x
	g.y = y
}

func (g *guardPosition) turnRight() {
	g.next_dir.turnRight()
}

func (g *guardPosition) move() {
	g.x += int(g.next_dir.dx)
	g.y += int(g.next_dir.dy)
}

func (g *guardPosition) moveBack() {
	g.x -= int(g.next_dir.dx)
	g.y -= int(g.next_dir.dy)
}

func (g guardPosition) getNextPositionAsPair() Pair[int, int] {
	return Pair[int, int]{g.x + int(g.next_dir.dx), g.y + int(g.next_dir.dy)}
}

func solve(input string) (int, int) {
	origStart, pillars, limit := getStartAndPillars(input)
	res, path := traverse(origStart, pillars, limit)
	possible_extras := make(map[Pair[int, int]]struct{})
	for pos, dirs := range path {
		if pos.First == origStart.x && pos.Second == origStart.y {
			continue
		}
		if pos.First < 0 || pos.Second < 0 || pos.First > limit || pos.Second > limit {
			continue
		}
		newStart := guardPosition{pos.First, pos.Second, dirs[0]}
		newStart.moveBack()
		pillars[pos] = struct{}{}
		size, _ := traverse(newStart, pillars, limit)
		if size == 0 {
			possible_extras[pos] = struct{}{}
		}
		delete(pillars, pos) // Remove pillar after traversal
	}

	return res, len(possible_extras)
}

func getStartAndPillars(input string) (guardPosition, map[Pair[int, int]]struct{}, int) {
	pillars := make(map[Pair[int, int]]struct{})
	start := guardPosition{0, 0, dir{0, -1}}
	limit := 0
	for y, row := range strings.Split(input, "\n") {
		for x := range row {
			c := row[x]
			if c == '#' {
				pair := Pair[int, int]{x, y}
				pillars[pair] = struct{}{}
			}
			if c == '^' {
				start.setPosXY(x, y)
			}

		}
		limit = y
	}
	return start, pillars, limit
}

func traverse(start guardPosition, pillars map[Pair[int, int]]struct{}, limit int) (int, map[Pair[int, int]][]dir) {
	visited := make(map[Pair[int, int]][]dir)

	for i := 0; start.x <= limit && start.y <= limit && start.y >= 0 && start.x >= 0; i++ {
		curr := Pair[int, int]{start.x, start.y}
		if slices.Contains(visited[curr], start.next_dir) {
			return 0, visited
		}
		visited[curr] = append(visited[curr], start.next_dir)

		for next := start.getNextPositionAsPair(); hasKey(pillars, next); next = start.getNextPositionAsPair() {
			start.turnRight()
		}
		start.move()
	}

	return len(visited), visited
}

func hasKey[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}

type Pair[T, U int] struct {
	First  int
	Second int
}

func printGrid(symbols map[Pair[int, int]]string, size int) {
	for y := range size + 1 {
		for x := range size + 1 {
			if sym, ok := symbols[Pair[int, int]{x, y}]; ok {
				fmt.Print(sym)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println("Example result:")
	fmt.Println(solve(example))

	fmt.Println("Real:")
	fmt.Println(solve(aoc_helpers.GetInputFromFile("06")))
}

const example = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`
