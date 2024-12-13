package main

import (
	"aoc"
	"testing"
)

func BenchmarkDay13(b *testing.B) {
	input := aoc.GetInputFromFile("13")
	for i := 0; i < b.N; i++ {
		solve(input)
	}
}
