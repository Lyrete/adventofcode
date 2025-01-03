package main

import (
	"aoc"
	"testing"
)

func BenchmarkDay15(b *testing.B) {
	input := aoc.GetInputFromFile("15")
	for i := 0; i < b.N; i++ {
		solve(input)
	}
}
