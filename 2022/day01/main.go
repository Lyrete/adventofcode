package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
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

	sort.Slice(carried, func(i, j int) bool {
		return carried[i] > carried[j]
	})

	//printSlice(carried)

	maxTop3 := carried[0] + carried[1] + carried[2]

	fmt.Printf("First task, max carried: %d\n", carried[0])
	fmt.Printf("Second task, top 3 carried: %d\n", maxTop3)
}
