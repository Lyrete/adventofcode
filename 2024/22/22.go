package main

import (
	"aoc"
	"fmt"
	"strconv"
	"strings"
)

func parse(input string) []int {
	split := strings.Split(input, "\n")
	ret := make([]int, len(split))
	for i, l := range split {
		val, _ := strconv.Atoi(l)
		ret[i] = val
	}
	return ret
}

func solve(input string) (int, int) {
	res, res2 := 0, 0
	initial := parse(input)
	diffs := [4]int{}
	sequences := make(map[[4]int]int)
	for _, secret := range initial {
		prev := secret % 10
		updated := make(map[[4]int]struct{})
		for n := range ITERATIONS {
			secret = getNext(secret)
			last := secret % 10
			if n > 3 {
				if _, ok := updated[diffs]; !ok {
					sequences[diffs] += prev
					updated[diffs] = struct{}{}
				}

				diffs[0] = diffs[1]
				diffs[1] = diffs[2]
				diffs[2] = diffs[3]
				diffs[3] = last - prev

			} else {
				diffs[n] = last - prev
			}
			prev = last
		}

		res += secret
	}
	//fmt.Println(sequences)
	for _, v := range sequences {
		if res2 < v {
			//fmt.Println("best", k)
			res2 = v
		}
	}
	return res, res2
}

func main() {
	fmt.Println("Example result:")
	fmt.Println(solve(example))

	fmt.Println("Real:")
	fmt.Println(solve(aoc.GetInputFromFile("22")))
}

func getNext(input int) int {
	input = ((input * 64) ^ input) % MOD
	input = ((input / 32) ^ input) % MOD
	return ((input * 2048) ^ input) % MOD
}

const example = `1
2
3
2024`
const MOD = 16777216
const ITERATIONS = 2000
