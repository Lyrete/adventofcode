package main

import (
	"aoc_helpers"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	example := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	fmt.Println("Example result:")
	fmt.Println(solve(example))

	fmt.Println("Real:")
	fmt.Println(solve(aoc_helpers.GetInputFromFile("02")))

}

func solve(input string) (int, int) {
	rows := strings.Split(input, "\n")

	safe := 0
	safe2 := 0

	for _, row := range rows {
		fields := strings.Fields(row)
		safe_row := isRowSafe(fields)
		safe_modified := false

		if !safe_row {
			for i := range fields {
				modifiable := slices.Clone(fields)
				modifiable = slices.Delete(modifiable, i, i+1)
				if isRowSafe(modifiable) {
					safe_modified = true
					break
				}
			}
		}

		if safe_row {
			safe++
		}

		if safe_modified {
			safe2++
		}
	}

	return safe, safe2 + safe
}

func isRowSafe(row []string) bool {
	safe_row := true
	for i := len(row) - 1; i > 1; i-- {
		l, _ := strconv.Atoi(row[i-2])
		c, _ := strconv.Atoi(row[i-1])
		r, _ := strconv.Atoi(row[i])
		if aoc_helpers.AbsDiffInt(c, r) > 3 || aoc_helpers.AbsDiffInt(c, r) < 1 || (r-c)*(c-l) <= 0 || aoc_helpers.AbsDiffInt(c, l) > 3 || aoc_helpers.AbsDiffInt(c, l) < 1 {
			safe_row = false
			break
		}
	}

	return safe_row
}
