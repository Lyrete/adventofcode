package main

import (
	"aoc"
	"testing"
)

func BenchmarkDay19(b *testing.B) {

	for i := 0; i < b.N; i++ {
		input := aoc.GetInputFromFile("19")
		solve(input)
	}
}
