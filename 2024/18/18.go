package main

import (
	"aoc"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func parse(input string) (walls []aoc.Coord) {
	re := regexp.MustCompile(`\d+`)
	split := strings.Split(input, "\n")
	walls = make([]aoc.Coord, len(split))
	for i, line := range split {
		nums := re.FindAllString(line, 2)
		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])
		walls[i] = aoc.Coord{X: x, Y: y}
	}
	return
}

type elem struct {
	c   aoc.Coord
	len int
}

func (e *elem) getNeigbours() []aoc.Coord {
	return e.c.Neighbours()
}

func findPath(walls []aoc.Coord, visited map[aoc.Coord]struct{}, start, end aoc.Coord) int {
	q := []elem{{c: start, len: 0}}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if slices.Contains(walls, curr.c) {
			continue
		}

		if curr.c == end {
			return curr.len
		}

		// is it already visited
		if _, ok := visited[curr.c]; ok {
			continue
		}
		visited[curr.c] = struct{}{}
		// Can't cheat by going out of the grid
		if curr.c.X < 0 || curr.c.Y < 0 || curr.c.X > end.X || curr.c.Y > end.Y {
			continue
		}

		for _, n := range curr.getNeigbours() {
			e := elem{c: n, len: curr.len + 1}
			q = append(q, e)
		}
	}

	return -1
}

func solve(input string, example bool) (int, string) {
	var end aoc.Coord
	var byteLim int
	start := aoc.Coord{X: 0, Y: 0}
	if example {
		end = aoc.Coord{X: 6, Y: 6}
		byteLim = 12
	} else {
		end = aoc.Coord{X: 70, Y: 70}
		byteLim = 1024
	}
	walls := parse(input)
	res := findPath(walls[:byteLim], map[aoc.Coord]struct{}{}, start, end)

	var res2 string
	if !example {
		l := byteLim
		r := len(walls)
		for l < r {
			gap := (r - l) / 2

			if gap == 0 {
				res2 = walls[l].AsString()
				break
			}

			mid_res := findPath(walls[:l+gap], map[aoc.Coord]struct{}{}, start, end)
			if mid_res != -1 {
				l += gap
			} else {
				r -= gap
			}

		}
	}

	return res, res2
}

func main() {
	fmt.Println("Example result:")
	fmt.Println(solve(example, true))

	fmt.Println("Real:")
	fmt.Println(solve(aoc.GetInputFromFile("18"), false))
}

const example = `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`
