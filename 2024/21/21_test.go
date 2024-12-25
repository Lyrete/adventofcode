package main

import (
	"aoc"
	"testing"
)

func BenchmarkDay21(b *testing.B) {

	for i := 0; i < b.N; i++ {
		input := aoc.GetInputFromFile("21")
		solve(input)
	}
}