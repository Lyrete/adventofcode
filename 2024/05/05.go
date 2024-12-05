package main

import (
	"aoc_helpers"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func solve(input string) (int, int) {
	chunks := strings.SplitN(input, "\n\n", 2)
	vertices := strings.Fields(chunks[0])
	graph := createGraph(vertices)

	sorting := func(a int, b int) int {
		aPrevious := graph[a]
		bPrevious := graph[b]
		return slices.Index(aPrevious, b) - slices.Index(bPrevious, a)
	}

	res := 0
	res2 := 0
	for _, pagesString := range strings.Split(chunks[1], "\n") {
		slice := aoc_helpers.ParseIntoIntSlice(pagesString)
		valid := slices.IsSortedFunc(slice, sorting)
		if !valid {
			slices.SortFunc(slice, sorting)
			res2 += slice[len(slice)/2]
		} else {
			res += slice[len(slice)/2]
		}
	}

	return res, res2
}

func createGraph(vertices []string) map[int][]int {
	graph := make(map[int][]int)

	for _, vertex := range vertices {
		parts := strings.Split(vertex, "|")
		headVal, _ := strconv.Atoi(parts[0])
		tailVal, _ := strconv.Atoi(parts[1])

		graph[tailVal] = append(graph[tailVal], headVal)
	}

	return graph
}

func main() {
	fmt.Println("Example result:")
	fmt.Println(solve(example))

	fmt.Println("Real:")
	fmt.Println(solve(aoc_helpers.GetInputFromFile("05")))
}

const example = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`
