package main

import (
	"aoc"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func getLen(i int) int {
	if i == 0 {
		return 1
	}
	length := 0
	for i != 0 {
		i /= 10
		length++
	}
	return length
}

func getNewStones(stone int) (int, int) {
	if stone == 0 {
		return 1, -1
	}

	length := getLen(stone)
	if length%2 == 0 {
		left := stone / int(math.Pow10(length/2))
		right := stone % int(math.Pow10(length/2))
		return left, right
	}

	return stone * 2024, -1
}

func blink(stones map[int]int) map[int]int {
	ret := make(map[int]int)
	for stone, amt := range stones {
		l, r := getNewStones(stone)
		if r != -1 {
			ret[r] += amt
		}
		ret[l] += amt
	}
	return ret
}

func getStoneAmount(stones map[int]int) int {
	res := 0
	for _, amt := range stones {
		res += amt
	}
	return res
}

func solve(input string) (int, int) {
	ans, ans2 := 0, 0

	stones := make(map[int]int)
	for _, s := range strings.Fields(input) {
		num, _ := strconv.Atoi(s)
		stones[num] = 1
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
