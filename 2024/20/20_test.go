package main

import (
	"aoc"
	"testing"
)

func BenchmarkDay20(b *testing.B) {

	for i := 0; i < b.N; i++ {
		input := aoc.GetInputFromFile("20")
		solve(input)
	}
}
