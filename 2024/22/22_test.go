package main

import (
	"aoc"
	"testing"
)

func BenchmarkDay22(b *testing.B) {

	for i := 0; i < b.N; i++ {
		input := aoc.GetInputFromFile("22")
		solve(input)
	}
}
