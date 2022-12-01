package main

import (
	"bufio"
	"fmt"
	"os"
)

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func findMax(s []int) (int, int) {
	var max int = s[0]
	var maxIdx int = 0
	for i, value := range s {
		if max < value {
			max = value
			maxIdx = i
		}
	}
	return maxIdx, max
}

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		fmt.Printf("Error")
		return
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	i := 0
	load := 0
	var carried []int

	for scanner.Scan() {
		line := scanner.Text()
		intLine := 0

		if len(line) == 0 {
			carried = append(carried, load)
			load = 0
			i++
			continue
		}

		_, err := fmt.Sscan(line, &intLine)

		if err != nil {
			fmt.Println(err)
			continue
		}

		load += intLine
	}

	carried = append(carried, load)

	//printSlice(carried)

	idx, max := findMax(carried)
	carried = RemoveIndex(carried, idx)
	idx, max2 := findMax(carried)
	carried = RemoveIndex(carried, idx)
	idx, max3 := findMax(carried)

	maxTop3 := max + max2 + max3

	fmt.Printf("First task, max carried: %d\n", max)
	fmt.Printf("Second task, top 3 carried: %d\n", maxTop3)
}
