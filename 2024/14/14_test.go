package main

import (
	"aoc"
	"testing"
)

func BenchmarkDay14(b *testing.B) {
	input := aoc.GetInputFromFile("14")
	for i := 0; i < b.N; i++ {
		solve(input, 101, 103)
	}
}
