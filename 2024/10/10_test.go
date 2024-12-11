package main

import (
	"aoc"
	"testing"
)

func BenchmarkDay10(b *testing.B) {
	input := aoc.GetInputFromFile("10")
	for i := 0; i < b.N; i++ {
		solve(input)
	}
}
