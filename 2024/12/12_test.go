package main

import (
	"aoc"
	"testing"
)

func BenchmarkDay12(b *testing.B) {
	input := aoc.GetInputFromFile("12")
	for i := 0; i < b.N; i++ {
		solve(input)
	}
}
