package main

import (
	"aoc"
	"fmt"
	"strconv"
	"strings"
)

func solve(input string) (int, int) {
	res := 0
	res2 := 0
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ":")
		expectedVal, _ := strconv.Atoi(parts[0])
		instructions := strings.Fields(parts[1])
		convInstructions := make([]int, len(instructions))
		for i, in := range instructions {
			conv, _ := strconv.Atoi(in)
			convInstructions[i] = conv
		}

		if checkIfValidRow(convInstructions[1:], convInstructions[0], expectedVal, false) {
			res += expectedVal
		} else if checkIfValidRow(convInstructions[1:], convInstructions[0], expectedVal, true) {
			res2 += expectedVal
		}
	}

	return res, res + res2
}

func checkIfValidRow(instructions []int, prev int, expectedValue int, withConcat bool) bool {
	if prev > expectedValue {
		return false
	}

	multiply := prev * instructions[0]
	addition := prev + instructions[0]
	concat, _ := strconv.Atoi(strconv.Itoa(prev) + strconv.Itoa(instructions[0]))
	if len(instructions) == 1 {
		return multiply == expectedValue || addition == expectedValue || (withConcat && concat == expectedValue)
	}

	return checkIfValidRow(instructions[1:], multiply, expectedValue, withConcat) || checkIfValidRow(instructions[1:], addition, expectedValue, withConcat) || (withConcat && checkIfValidRow(instructions[1:], concat, expectedValue, withConcat))
}

func main() {
	fmt.Println("Example result:")
	fmt.Println(solve(example))

	fmt.Println("Real:")
	fmt.Println(solve(aoc.GetInputFromFile("07")))
}

const example = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`
