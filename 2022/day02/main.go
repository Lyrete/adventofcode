package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getPoints(char1 rune, char2 rune) int {
	resultPoints := 0
	if char1 == char2 {
		resultPoints = 3
	} else if char2-char1 == -1 || char2-char1 == 2 {
		resultPoints = 0
	} else if char2-char1 == -2 || char2-char1 == 1 {
		resultPoints = 6
	}
	return resultPoints + int(char2) - int('A') + 1
}

func main() {
	modify := func(c rune) rune {
		if c == 'X' {
			return 'A'
		} else if c == 'Y' {
			return 'B'
		} else {
			return 'C'
		}
	}

	f, _ := os.Open("data.txt")

	defer f.Close()
	scanner := bufio.NewScanner(f)

	points := 0
	points2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, " ")
		first, second := splitLine[0], splitLine[1]
		secondModified := strings.Map(modify, second)

		char1 := []rune(first)[0]
		char2 := []rune(secondModified)[0]

		choices := [3]rune{'A', 'B', 'C'}
		if second == "Y" {
			chosen := char1
			points2 += getPoints(char1, chosen)
		} else {
			min := 100
			max := 0
			for _, char := range choices {
				given_points := getPoints(char1, char)
				if given_points < min {
					min = given_points
				}
				if given_points > max {
					max = given_points
				}
			}

			if second == "X" {
				points2 += min
			}

			if second == "Z" {
				points2 += max
			}
		}

		points += getPoints(char1, char2)
	}

	fmt.Println("First task:", points)
	fmt.Println("Second task:", points2)
}

// 1 should be win on A and 2 loss
// -1 should be loss on B and 1 win
// -2 should be win and -1 should be loss
