package main

import (
	"aoc"
	"fmt"
	"slices"
	"strings"
)

func parse(input string) []aoc.Coord {
	pathUnique := make(map[aoc.Coord]struct{})
	var start aoc.Coord
	var end aoc.Coord
	for y, line := range strings.Split(input, "\n") {
		for x, r := range line {
			switch r {
			case '.':
				pathUnique[aoc.Coord{X: x, Y: y}] = struct{}{}
			case 'S':
				start = aoc.Coord{X: x, Y: y}
				pathUnique[start] = struct{}{}
			case 'E':
				end = aoc.Coord{X: x, Y: y}
				pathUnique[end] = struct{}{}
			}
		}
	}
	return getNormalPath(pathUnique, start, end)
}

func getNormalPath(pathUnique map[aoc.Coord]struct{}, start aoc.Coord, end aoc.Coord) []aoc.Coord {
	path := make([]aoc.Coord, len(pathUnique))
	path[0] = start
	i := 1
	for start != end {
		for _, n := range start.Neighbours() {
			if i < 2 || path[i-2] != n {
				if _, ok := pathUnique[n]; ok {
					path[i] = n
					start = n
					break
				}
			}
		}
		i++
	}
	return path
}

func findCheatablePositions(path []aoc.Coord, minSaved int) int {
	dirs := [][]int{{1, 0}, {0, 1}, {0, -1}, {-1, 0}}
	res := 0

	for i := 0; i < len(path)-minSaved; i++ {
		c := &path[i]
		for _, dir := range dirs {
			n := c.AddXY(dir[0], dir[1])
			// cant cheat in either direction the path already goes
			if (i < len(path)-1 && path[i+1] == n) || (i > 0 && path[i-1] == n) {
				continue
			}

			n.MoveXY(dir[0], dir[1])
			cheatIndex := slices.Index(path[i:], n)
			if cheatIndex-2 >= minSaved {
				res += 1
			}

		}
	}
	return res
}

func findCheatableWithMaxLength(path []aoc.Coord, startIdx, maxLength int, minSaved int) int {
	start := path[startIdx]
	res := 0
	i := startIdx + minSaved
	for i < len(path) {
		end := path[i]
		cheatLength := start.ManhattanDistance(end)
		if cheatLength > maxLength {
			i += cheatLength - maxLength
			continue
		}

		realDist := i - startIdx
		savedDist := realDist - cheatLength
		if cheatLength <= maxLength && savedDist >= minSaved {
			res += 1
		}

		i++
	}
	return res
}

func solve(input string) (int, int) {
	res, res2 := 0, 0
	path := parse(input)
	res += findCheatablePositions(path, 100)
	for i := 0; i < len(path)-100; i++ {
		//res += findCheatableClose(path, i, 100)
		res2 += findCheatableWithMaxLength(path, i, 20, 100)
	}

	return res, res2
}

func main() {
	fmt.Println("Example result:")
	fmt.Println(solve(example))

	fmt.Println("Real:")
	fmt.Println(solve(aoc.GetInputFromFile("20")))
}

const example = `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`
