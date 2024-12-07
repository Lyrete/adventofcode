package main

import (
	"aoc"
	"testing"
)

func BenchmarkDay7Backwards(b *testing.B) {
	input := aoc.GetInputFromFile("07")
	for i := 0; i < b.N; i++ {
		solve(input)
	}
}

func BenchmarkDay7Forwards(b *testing.B) {
	input := aoc.GetInputFromFile("07")
	for i := 0; i < b.N; i++ {
		solveForward(input)
	}
}
