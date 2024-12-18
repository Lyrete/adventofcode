package main

import (
	"aoc"
	"testing"
)

func BenchmarkDay18(b *testing.B) {

	for i := 0; i < b.N; i++ {
		input := aoc.GetInputFromFile("18")
		solve(input, false)
	}
}
