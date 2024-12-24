package main

import (
	"aoc"
	"fmt"
	"maps"
	"slices"
	"strconv"
	"strings"
)

type operation struct {
	gate1     string
	gate2     string
	operation string
	target    string
}

func (o *operation) doOperationFromBitMap(bits map[string]int) int {
	switch o.operation {
	case "XOR":
		return bits[o.gate1] ^ bits[o.gate2]
	case "AND":
		return bits[o.gate1] & bits[o.gate2]
	case "OR":
		return bits[o.gate1] | bits[o.gate2]
	default:
		err := fmt.Sprintf("Unrecognized operation: %s", o.operation)
		panic(err)
	}
}

func (o *operation) doOperation(input1, input2 int) int {
	switch o.operation {
	case "XOR":
		return input1 ^ input2
	case "AND":
		return input1 & input2
	case "OR":
		return input1 | input2
	default:
		err := fmt.Sprintf("Unrecognized operation: %s", o.operation)
		panic(err)
	}
}

func parse(input string) (map[string]int, map[string]operation) {
	parts := strings.Split(input, "\n\n")
	bits := make(map[string]int)
	for _, l := range strings.Split(parts[0], "\n") {
		if l[5] == '1' {
			bits[l[0:3]] = 1
		} else {
			bits[l[0:3]] = 0
		}
	}

	splitIns := strings.Split(parts[1], "\n")
	output := make(map[string]operation)
	for _, line := range splitIns {
		split := strings.Split(line, " ")
		output[split[4]] = operation{gate1: split[0], gate2: split[2], operation: split[1], target: split[4]}
	}
	return bits, output
}

func findValue(startOp operation, bits map[string]int, graph map[string]operation) int {
	var left int
	var right int
	if v, ok := bits[startOp.gate1]; !ok {
		left = findValue(graph[startOp.gate1], bits, graph)
	} else {
		left = v
	}

	if v, ok := bits[startOp.gate2]; !ok {
		right = findValue(graph[startOp.gate2], bits, graph)
	} else {
		right = v
	}

	return startOp.doOperation(left, right)
}

func findAllInputs(path []string, start string, graph map[string]operation) []string {
	op := graph[start]
	path = append(path, op.target)
	if op.gate1[0] != 'x' && op.gate1[0] != 'y' {
		path = findAllInputs(path, op.gate1, graph)
	}

	if op.gate2[0] != 'x' && op.gate2[0] != 'y' {
		path = findAllInputs(path, op.gate2, graph)
	}

	return path
}

func swapInstructions(graph map[string]operation, key1, key2 string) map[string]operation {
	temp1 := graph[key1]
	temp2 := graph[key2]
	temp2.target = key1
	temp1.target = key2
	graph[key1] = temp2
	graph[key2] = temp1
	return graph
}

func solve(input string) (int, string) {
	res := 0
	bits, instructions := parse(input)

	// for _, op := range instructions {
	// 	res := op.doOperation(bits)
	// 	bits[op.target] = res
	// }

	mustSwap := map[string]struct{}{}

	for output, instruction := range instructions {
		if instruction.operation != "XOR" && output[0] == 'z' && output != "z45" {
			mustSwap[output] = struct{}{}
		}

		if instruction.operation == "XOR" &&
			(instruction.gate1[0] != 'x' && instruction.gate1[0] != 'y' && instruction.gate1[0] != 'z') &&
			(instruction.gate2[0] != 'x' && instruction.gate2[0] != 'y' && instruction.gate2[0] != 'z') &&
			(output[0] != 'x' && output[0] != 'y' && output[0] != 'z') {
			mustSwap[output] = struct{}{}
		}

		if instruction.operation == "XOR" {
			for _, xop := range instructions {
				if (output == xop.gate1 || output == xop.gate2) && xop.operation == "OR" {
					mustSwap[output] = struct{}{}
				}
			}
		}

		if instruction.operation == "AND" && instruction.gate1 != "x00" && instruction.gate2 != "x00" {
			for _, xop := range instructions {
				if (output == xop.gate1 || output == xop.gate2) && xop.operation != "OR" {
					mustSwap[output] = struct{}{}
				}
			}
		}

		if output[0] != 'z' {
			continue
		}
		output := findValue(instruction, bits, instructions)
		i, _ := strconv.Atoi(instruction.target[1:])
		// 	//fmt.Println()
		// }
		//fmt.Println(k, bits[k])
		res = res + (output << i)
	}

	return res, strings.Join(slices.Sorted(maps.Keys(mustSwap)), ",")
}

func main() {
	fmt.Println("Example result:")
	fmt.Println(solve(example))

	fmt.Println("Real:")
	fmt.Println(solve(aoc.GetInputFromFile("24")))
}

const example = `x00: 0
x01: 1
x02: 0
x03: 1
x04: 0
x05: 1
y00: 0
y01: 0
y02: 1
y03: 1
y04: 0
y05: 1

x00 AND y00 -> z05
x01 AND y01 -> z02
x02 AND y02 -> z01
x03 AND y03 -> z03
x04 AND y04 -> z04
x05 AND y05 -> z00`
