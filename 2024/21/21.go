package main

import (
	"aoc"
	"fmt"
	"strconv"
	"strings"
)

func getRobotFromKeypad(code string) []string {
	ret := make([]string, len(code))
	lastChar := 'A'
	for i, r := range code {
		ret[i] = keypadPressMap[lastChar][r] + "A"
		lastChar = r
	}
	return ret
}

func getFollowingPresses(code string) []string {
	ret := make([]string, len(code))
	lastChar := 'A'
	for i, r := range code {
		ret[i] = robotPressMap[lastChar][r] + "A"
		lastChar = r
	}
	return ret
}

func getPresses(robotInstructions []string, robotAmount int) int {
	patterns := make(map[string]int)
	for _, p := range robotInstructions {
		patterns[p] += 1
	}
	for range robotAmount {
		tempPatterns := make(map[string]int)
		for pattern, amt := range patterns {
			next := getFollowingPresses(pattern)
			for _, n := range next {
				tempPatterns[n] += amt
			}
		}
		patterns = tempPatterns
	}
	size := 0
	for k, v := range patterns {
		size += (len(k)) * v
	}
	return size
}

func solve(input string) (int, int) {
	res, res2 := 0, 0
	codes := parse(input)
	for _, code := range codes {
		robot := getRobotFromKeypad(code)

		numeric, _ := strconv.Atoi(code[:len(code)-1])

		res2 += getPresses(robot, 25) * numeric
		res += getPresses(robot, 2) * numeric
	}
	return res, res2
}

func main() {
	fmt.Println("Example result:")
	fmt.Println(solve(example))

	fmt.Println("Real:")
	fmt.Println(solve(aoc.GetInputFromFile("21")))
}

const example = `029A
980A
179A
456A
379A`

func parse(input string) []string {
	return strings.Split(input, "\n")
}

var keypadPressMap = map[rune]map[rune]string{
	// A starts
	'A': {
		'0': "<",
		'1': "^<<",
		'2': "<^",
		'3': "^",
		'4': "^^<<",
		'5': "<^^",
		'6': "^^",
		'7': "^^^<<",
		'8': "<^^^",
		'9': "^^^"},
	'0': {
		'A': ">",
		'1': "^<",
		'2': "^",
		'3': ">^",
		'4': "^^<",
		'5': "^^",
		'6': ">^^",
		'7': "^^^<",
		'8': "^^^",
		'9': ">^^^"},
	'1': {
		'0': ">v",
		'A': ">>v",
		'2': ">",
		'3': ">>",
		'4': "^",
		'5': ">^",
		'6': ">>^",
		'7': "^^",
		'8': "^^>",
		'9': ">>^^"},
	'2': {
		'0': "v",
		'1': "<",
		'A': "v>",
		'3': ">",
		'4': "<^",
		'5': "^",
		'6': ">^",
		'7': "<^^",
		'8': "^^",
		'9': "^^>"},
	'3': {
		'0': "<v",
		'1': "<<",
		'2': "<",
		'A': "v",
		'4': "<<^",
		'5': "<^",
		'6': "^",
		'7': "<<^^",
		'8': "<^^",
		'9': "^^"},
	'4': {
		'0': ">vv",
		'1': "v",
		'2': "v>",
		'3': "v>>",
		'A': ">>vv",
		'5': ">",
		'6': ">>",
		'7': "^",
		'8': ">^",
		'9': ">>^"},
	'5': {
		'0': "vv",
		'1': "<v",
		'2': "v",
		'3': "v>",
		'4': "<",
		'A': "vv>",
		'6': ">",
		'7': "<^",
		'8': "^",
		'9': ">^"},
	'6': {
		'0': "<vv",
		'1': "<<v",
		'2': "<v",
		'3': "v",
		'4': "<<",
		'5': "<",
		'A': "vv",
		'7': "<<^",
		'8': "<^",
		'9': "^"},
	'7': {
		'0': "vvv>",
		'1': "vv",
		'2': "vv>",
		'3': "vv>>",
		'4': "v",
		'5': "v>",
		'6': "v>>",
		'A': ">>vvv",
		'8': ">",
		'9': ">>"},
	'8': {
		'0': "vvv",
		'1': "<vv",
		'2': "vv",
		'3': "vv>",
		'4': "v<",
		'5': "v",
		'6': "v>",
		'7': "<",
		'A': "vvv>",
		'9': ">"},
	'9': {
		'0': "<vvv",
		'1': "<<vv",
		'2': "<vv",
		'3': "vv",
		'4': "<<v",
		'5': "<v",
		'6': "v",
		'7': "<<",
		'8': "<",
		'A': "vvv"},
}

var robotPressMap = map[rune]map[rune]string{
	// A starts
	'A': {
		'<': "v<<",
		'v': "<v",
		'^': "<",
		'>': "v"},
	// > starts
	'>': {
		'^': "<^",
		'A': "^",
		'v': "<",
		'<': "<<"},
	// <
	'<': {
		'^': ">^",
		'A': ">>^",
		'v': ">",
		'>': ">>"},
	// ^
	'^': {
		'<': "v<",
		'A': ">",
		'v': "v",
		'>': "v>"},
	// v
	'v': {
		'<': "<",
		'A': "^>",
		'^': "^",
		'>': ">"},
}
