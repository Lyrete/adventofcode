package aoc_helpers

import (
	"log"
	"os"
	"strings"
)

func GetInputFromFile(day string) string {
	filename := "../data/" + day + ".txt"
	cnt, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Failed to read file")
	}
	return strings.Trim(string(cnt), "\n ")
}

func GetRowsFromFile(day string) []string {
	return strings.Split(GetInputFromFile(day), "\n")
}

func AbsInt(x int) int {
	return AbsDiffInt(x, 0)
}

func AbsDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
