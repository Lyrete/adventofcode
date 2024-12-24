package main

import (
	"aoc"
	"testing"
)

func BenchmarkDay23(b *testing.B) {

	for i := 0; i < b.N; i++ {
		input := aoc.GetInputFromFile("23")
		solve(input)
	}
}
