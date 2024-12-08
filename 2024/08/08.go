package main

import (
	"aoc"
	"fmt"
	"strings"
)

type Coord struct {
	x int
	y int
}

func (c1 Coord) getPossibleAntinodes(c2 Coord) (Coord, Coord) {
	dx, dy := c1.x-c2.x, c1.y-c2.y
	return Coord{c1.x + dx, c1.y + dy}, Coord{c2.x - dx, c2.y - dy}
}

func (c1 Coord) getDist(c2 Coord) (int, int) {
	return c1.x - c2.x, c1.y - c2.y
}

func (c *Coord) addXY(x int, y int) {
	c.x += x
	c.y += y
}

func solve(input string) (int, int) {
	antennas, limit := getAntennas(input)
	foundAntinodes := make(map[Coord]struct{})
	foundAntinodesP2 := make(map[Coord]struct{})
	for _, coords := range antennas {
		for len(coords) > 1 {
			curr := coords[0]
			foundAntinodesP2[curr] = struct{}{}
			coords = coords[1:]
			for _, c := range coords {
				p1, p2 := curr.getPossibleAntinodes(c)
				foundAntinodesP2[c] = struct{}{}
				dx, dy := curr.getDist(c)
				if 0 <= p1.x && p1.x <= limit && 0 <= p1.y && p1.y <= limit {
					foundAntinodes[p1] = struct{}{}
					for 0 <= p1.x && p1.x <= limit && 0 <= p1.y && p1.y <= limit {
						foundAntinodesP2[p1] = struct{}{}
						p1.addXY(dx, dy)
					}
				}
				if 0 <= p2.x && p2.x <= limit && 0 <= p2.y && p2.y <= limit {
					foundAntinodes[p2] = struct{}{}
					for 0 <= p2.x && p2.x <= limit && 0 <= p2.y && p2.y <= limit {
						foundAntinodesP2[p2] = struct{}{}
						p2.addXY(-dx, -dy)
					}
				}
			}
		}
	}
	return len(foundAntinodes), len(foundAntinodesP2)
}

func getAntennas(input string) (map[rune][]Coord, int) {
	antennas := make(map[rune][]Coord)
	limit := 0
	for y, line := range strings.Split(input, "\n") {
		for x, r := range line {
			if r != '.' {
				antennas[r] = append(antennas[r], Coord{x, y})
			}
		}
		limit = y
	}
	return antennas, limit
}

func main() {
	fmt.Println("Example result:")
	fmt.Println(solve(example))

	fmt.Println("Real:")
	fmt.Println(solve(aoc.GetInputFromFile("08")))
}

const example = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`
