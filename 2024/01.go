package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	example := `3   4
	4   3
	2   5
	1   3
	3   9
	3   3`

	fmt.Println("Example:")
	fmt.Println(solve(example))

	fmt.Println("Real:")
	fmt.Println(solve(getInputFromFile("01")))

}

func getInputFromFile(day string) string {
	filename := "data/" + day + ".txt"
	cnt, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Failed to read file")
	}
	return strings.Trim(string(cnt), "\n ")
}

func solve(input string) (int, int) {
	rows := strings.Split(input, "\n")
	left, right := make([]int, len(rows)), make([]int, len(rows))

	occurences := make(map[int]int)
	for idx, row := range rows {
		parts := strings.Fields(row)
		l, _ := strconv.Atoi(parts[0])
		r, _ := strconv.Atoi(parts[1])
		left[idx] = l
		right[idx] = r
		occurences[r] += 1
	}

	slices.Sort(left)
	slices.Sort(right)

	res := 0
	res2 := 0

	for i, l := range left {
		r := right[i]
		res += absInt(r - l)
		res2 += l * occurences[l]
	}
	return res, res2
}

func absInt(x int) int {
	return absDiffInt(x, 0)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
