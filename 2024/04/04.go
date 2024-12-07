package main

import (
	"aoc"
	"fmt"
	"slices"
	"strings"
)

const example = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

type Coord struct {
	x int
	y int
}

func (c *Coord) move(in Coord) {
	c.x += in.x
	c.y += in.y
}

func (c Coord) addXY(x int, y int) Coord {
	c.x += x
	c.y += y

	return c
}

var nextChars = map[rune]rune{'X': 'M', 'M': 'A', 'A': 'S'}

func main() {

	fmt.Println("Example result:")
	fmt.Println(solve(example))

	fmt.Println("Real:")
	fmt.Println(solve(aoc.GetInputFromFile("04")))
}

func solve(input string) (int, int) {
	rows := strings.Split(input, "\n")
	places, coords := getCharPlacements(rows)
	xmasFound := 0
	masCrossed := 0
	for _, start := range places['X'] {
		xmasFound += findFullXmas(places, 'X', start, Coord{1, 0}) + findFullXmas(places, 'X', start, Coord{1, -1}) + findFullXmas(places, 'X', start, Coord{0, -1}) + findFullXmas(places, 'X', start, Coord{-1, -1}) + findFullXmas(places, 'X', start, Coord{-1, 0}) + findFullXmas(places, 'X', start, Coord{-1, 1}) + findFullXmas(places, 'X', start, Coord{0, 1}) + findFullXmas(places, 'X', start, Coord{1, 1})
	}

	for _, center := range places['A'] {
		masCrossed += findCrossedMas(coords, center)
	}

	return xmasFound, masCrossed
}

func getCharPlacements(rows []string) (map[rune][]Coord, map[Coord]rune) {
	charCoords := make(map[rune][]Coord)
	coords := make(map[Coord]rune)
	for y, row := range rows {
		for x, c := range row {
			if c == 'M' || c == 'X' || c == 'A' || c == 'S' {
				coord := Coord{x, y}
				charCoords[c] = append(charCoords[c], coord)
				coords[coord] = c
			}
		}
	}
	return charCoords, coords
}

func findFullXmas(places map[rune][]Coord, searchable rune, curr Coord, dir Coord) int {
	//fmt.Println(string(searchable), curr, dir)
	if searchable != 'M' && searchable != 'X' && searchable != 'A' && searchable != 'S' {
		return 1
	}

	if !slices.Contains(places[searchable], curr) {
		return 0
	}

	curr.move(dir)
	return findFullXmas(places, nextChars[searchable], curr, dir)
}

func findCrossedMas(coords map[Coord]rune, center Coord) int {
	surrounding := string([]rune{coords[center.addXY(1, 1)], coords[center.addXY(1, -1)], coords[center.addXY(-1, -1)], coords[center.addXY(-1, 1)]})
	if surrounding == "MSSM" || surrounding == "MMSS" || surrounding == "SMMS" || surrounding == "SSMM" {
		return 1
	}
	return 0
}
