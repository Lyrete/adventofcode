package main

import (
	"aoc"
	"fmt"
	"strings"
)

func parse(input string) ([]string, []string) {
	parts := strings.Split(input, "\n\n")
	insSplit := strings.Split(parts[0], ", ")
	designSplit := strings.Split(parts[1], "\n")

	return insSplit, designSplit
}

func checkPossible(design string, instructions []string, possible []int) int {

	for i := range len(design) + 1 {
		possible[i] = 0
	}

	possible[0] = 1

	for i := 0; i < len(design); i++ {
		if possible[i] == 0 && i > 0 {
			continue
		}

		for _, next := range instructions {
			if len(next)+i <= len(design) && next == design[i:len(next)+i] {
				possible[len(next)+i] += possible[i]
			}
		}
	}

	return possible[len(design)]
}

func solve(input string) (int, int) {
	ins, designs := parse(input)
	res, res2 := 0, 0
	possible := make([]int, 100)
	for _, d := range designs {
		add := checkPossible(d, ins, possible)
		if add > 0 {
			res += 1
		}
		res2 += add

	}
	return res, res2
}

func main() {
	fmt.Println("Example result:")
	fmt.Println(solve(example))

	fmt.Println("Real:")
	fmt.Println(solve(aoc.GetInputFromFile("19")))
}

const example = `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`
