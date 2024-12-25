package main

import (
	"aoc"
	"fmt"
	"maps"
	"strings"
)

func parseHeights(rows []string) [5]int {
	ret := [5]int{}
	for _, row := range rows {
		for i, r := range row {
			if r == '#' {
				ret[i] += 1
			}
		}
	}

	return ret
}

func parse(input string) ([][5]int, [][5]int) {
	parseable := strings.Split(input, "\n\n")
	locks := make([][5]int, 0)
	keys := make([][5]int, 0)
	for _, p := range parseable {
		rows := strings.Split(p, "\n")
		if rows[0] == "#####" {
			locks = append(locks, parseHeights(rows[1:]))
		} else {
			keys = append(keys, parseHeights(rows[:len(rows)-1]))
		}
	}
	return locks, keys
}

type lockKeyPair struct {
	lock [5]int
	key  [5]int
}

func (lp *lockKeyPair) hasNoOverlap() bool {
	for i := range len(lp.lock) {
		if lp.lock[i]+lp.key[i] > 5 {
			return false
		}
	}
	return true
}

func solve(input string) (int, int) {
	res := 0
	locks, keys := parse(input)
	fitting := map[lockKeyPair]struct{}{}
	for _, lock := range locks {
		for _, key := range keys {
			pair := lockKeyPair{lock, key}
			if pair.hasNoOverlap() {
				fitting[pair] = struct{}{}
			}
		}
	}

	for range maps.Keys(fitting) {
		res++
	}

	return res, 0
}

func main() {
	fmt.Println("Example result:")
	fmt.Println(solve(example))

	fmt.Println("Real:")
	fmt.Println(solve(aoc.GetInputFromFile("25")))
}

const example = `#####
.####
.####
.####
.#.#.
.#...
.....

#####
##.##
.#.##
...##
...#.
...#.
.....

.....
#....
#....
#...#
#.#.#
#.###
#####

.....
.....
#.#..
###..
###.#
###.#
#####

.....
.....
.....
#....
#.#..
#.#.#
#####`
