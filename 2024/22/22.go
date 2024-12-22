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
	lastDigit := [5]int{}
	sequences := make(map[[4]int]int)
	for _, secret := range initial {
		lastDigit[0] = secret % 10
		updated := make(map[[4]int]struct{})
		for n := range ITERATIONS {
			secret = getNext(secret)
			last := secret % 10
			if n > 3 {
				lastDigit[0] = lastDigit[1]
				lastDigit[1] = lastDigit[2]
				lastDigit[2] = lastDigit[3]
				lastDigit[3] = lastDigit[4]
				lastDigit[4] = last

				curr := [4]int{lastDigit[1] - lastDigit[0],
					lastDigit[2] - lastDigit[1],
					lastDigit[3] - lastDigit[2],
					last - lastDigit[3]}

				if _, ok := updated[curr]; !ok {
					sequences[curr] += last
					updated[curr] = struct{}{}
				}

			} else {
				lastDigit[n+1] = last
			}
		}

		res += secret
	}
	for _, v := range sequences {
		if res2 < v {
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
