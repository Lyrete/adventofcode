package main

import (
	"aoc"
	"testing"
)

func BenchmarkDay8(b *testing.B) {
	input := aoc.GetInputFromFile("08")
	for i := 0; i < b.N; i++ {
		solve(input)
	}
}
