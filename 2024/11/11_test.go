package main

import (
	"aoc"
	"testing"
)

func BenchmarkDay11(b *testing.B) {
	input := aoc.GetInputFromFile("11")
	for i := 0; i < b.N; i++ {
		solve(input)
	}
}
