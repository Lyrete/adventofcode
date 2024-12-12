package aoc

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strconv"
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

func ParseIntoIntSlice(input string) []int {
	separators := []byte{byte(',')}
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		if i := bytes.Index(data, separators); i >= 0 {
			return i + 1, data[0:i], nil
		}

		if atEOF {
			return len(data), data, nil
		}

		return 0, nil, nil
	})
	ret := []int{}
	for scanner.Scan() {
		v, _ := strconv.Atoi(scanner.Text())
		ret = append(ret, v)
	}
	return ret
}

type Coord struct {
	X int
	Y int
}

func (c *Coord) Move(in Coord) {
	c.X += in.X
	c.Y += in.Y
}

func (c *Coord) MoveXY(x int, y int) {
	c.X += x
	c.Y += y
}

func (c Coord) AddXY(x int, y int) Coord {
	c.X += x
	c.Y += y

	return c
}

func (c *Coord) Set(x int, y int) {
	c.X = x
	c.Y = y
}

func HasKey[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}
