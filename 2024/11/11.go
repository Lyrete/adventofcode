package main

import (
	"aoc"
	"fmt"
	"strconv"
	"strings"
)

func getNewStones(stone string) []string {
	num, _ := strconv.Atoi(stone)

	if num == 0 {
		return []string{"1"}
	}

	if len(stone)%2 == 0 {
		return []string{stone[0 : len(stone)/2], strings.TrimLeft(stone[len(stone)/2:], "0")}
	}

	return []string{strconv.Itoa(num * 2024)}
}

func blink(stones map[string]int) map[string]int {
	ret := make(map[string]int)
	for stone, amt := range stones {
		newStones := getNewStones(stone)
		for _, n := range newStones {
			ret[n] += amt
		}
	}
	return ret
}

func getStoneAmount(stones map[string]int) int {
	res := 0
	for _, amt := range stones {
		res += amt
	}
	return res
}

func solve(input string) (int, int) {
	ans, ans2 := 0, 0

	stones := make(map[string]int)
	for _, s := range strings.Fields(input) {
		stones[s] = 1
	}

	for range 25 {
		stones = blink(stones)
	}

	ans = getStoneAmount(stones)

	for range 75 - 25 {
		stones = blink(stones)
	}

	ans2 = getStoneAmount(stones)

	return ans, ans2
}

func main() {
	fmt.Println("Example result:")
	fmt.Println(solve(example))

	fmt.Println("Real:")
	fmt.Println(solve(aoc.GetInputFromFile("11")))
}

const example = `125 17`
