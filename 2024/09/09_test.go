package main

import (
	"aoc"
	"testing"
)

func BenchmarkDay9(b *testing.B) {
	input := aoc.GetInputFromFile("09")
	for i := 0; i < b.N; i++ {
		solve(input)
	}
}
