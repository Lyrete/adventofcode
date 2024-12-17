package main

import (
	"aoc"
	"fmt"
	"strings"
	"testing"
)

func BenchmarkDay17(b *testing.B) {

	for i := 0; i < b.N; i++ {
		input := aoc.GetInputFromFile("17")
		solve(input)
	}
}

func BenchmarkDay17P1(b *testing.B) {

	for i := 0; i < b.N; i++ {
		input := aoc.GetInputFromFile("17")
		reg, ins := parse(input)
		var sb strings.Builder
		//p1
		for reg.instructionPointer < len(ins) {
			ret := reg.doOp(ins[reg.instructionPointer], ins[reg.instructionPointer+1])
			if ret != "" {
				fmt.Fprintf(&sb, "%s,", ret)
			}
		}
		sb.String()
	}
}
func BenchmarkDay17P2(b *testing.B) {

	for i := 0; i < b.N; i++ {
		input := aoc.GetInputFromFile("17")
		_, ins := parse(input)
		findA(ins, 0, len(ins)-1)
	}
}
