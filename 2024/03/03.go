package main

import (
	"aoc_helpers"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	example := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

	fmt.Println("Example result:")
	fmt.Println(solve(example))

	fmt.Println("Real:")
	fmt.Println(solve(aoc_helpers.GetInputFromFile("03")))
}

func solve(input string) (int, int) {
	res := 0
	res2 := 0
	do := true
	for i := 0; i < len(input); i++ {
		c := input[i]
		if c == 'd' {
			do = findDoDont(input[i:], do)
		}

		if c != 'm' || input[i:i+3] != "mul" {
			continue
		}
		i += 3
		multiplicationResult, skip := findMultiply(input[i:])
		i += skip - 1
		res += multiplicationResult
		if do {
			res2 += multiplicationResult
		}
	}
	return res, res2
}

func findMultiply(input string) (int, int) {
	if input[0] != '(' {
		return 0, 1
	}

	valid := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', ',', ')'}
	buffer := ""
	endIndex := 0

	for i, c := range input[1:] {
		if !slices.Contains(valid, c) {
			return 0, i
		}
		if c == ')' {
			break
		}
		buffer += string(c)
		endIndex = i
	}
	parts := strings.Split(buffer, ",")
	l, _ := strconv.Atoi(parts[0])
	r, _ := strconv.Atoi(parts[1])
	return l * r, endIndex
}

func findDoDont(input string, current bool) bool {
	doString := "do()"
	dontString := "don't()"

	if len(input) >= 4 && input[:4] == doString {
		return true
	}

	if len(input) >= 7 && input[:7] == dontString {
		return false
	}

	return current
}
