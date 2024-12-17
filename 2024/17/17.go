package main

import (
	"aoc"
	"fmt"
	"strconv"
	"strings"
)

type registers struct {
	a                  int
	b                  int
	c                  int
	instructionPointer int
}

func (r *registers) doOp(op string, operand string) string {
	switch op {
	case "0":
		r.adv(r.getOperandValue(operand))
		return ""
	case "1":
		r.bxl(getLiteralValue(operand))
		return ""
	case "2":
		r.bst(r.getOperandValue(operand))
		return ""
	case "3":
		r.jnz(getLiteralValue(operand))
		return ""
	case "4":
		r.bxc(getLiteralValue(operand))
		return ""
	case "5":
		return strconv.Itoa(r.out(r.getOperandValue(operand)))
	case "6":
		r.bdv(r.getOperandValue(operand))
		return ""
	case "7":
		r.cdv(r.getOperandValue(operand))
		return ""
	default:
		panic("Unexpected operation")
	}

}

func getLiteralValue(operand string) int {
	val, _ := strconv.Atoi(operand)
	return val
}

func (r *registers) getOperandValue(operand string) int {
	switch operand {
	case "0":
		return 0
	case "1":
		return 1
	case "2":
		return 2
	case "3":
		return 3
	case "4":
		return r.a
	case "5":
		return r.b
	case "6":
		return r.c
	default:
		panic("Invalid program")
	}
}

func (r *registers) adv(combo_operand int) {
	value := r.a / intPow(2, combo_operand)
	r.a = value
	r.instructionPointer += 2
}

func (r *registers) bxl(lit_operand int) {
	value := r.b ^ lit_operand
	r.b = value
	r.instructionPointer += 2
}

func (r *registers) bst(combo_operand int) {
	r.b = combo_operand % 8
	r.instructionPointer += 2
}

func (r *registers) jnz(lit_operand int) {
	if r.a == 0 {
		r.instructionPointer += 2
		return
	}
	r.instructionPointer = lit_operand
}

func (r *registers) bxc(lit_operand int) {
	value := r.b ^ r.c
	r.b = value
	r.instructionPointer += 2
}

func (r *registers) out(combo_operand int) int {
	r.instructionPointer += 2
	return combo_operand % 8
}

func (r *registers) bdv(combo_operand int) {
	value := r.a / intPow(2, combo_operand)
	r.b = value
	r.instructionPointer += 2
}

func (r *registers) cdv(combo_operand int) {
	value := r.a / intPow(2, combo_operand)
	r.c = value
	r.instructionPointer += 2
}

func parse(input string) (registers, []string) {
	a := strings.Split(input, "\n")
	reg_a, _ := strconv.Atoi(a[0][12:])
	reg_b, _ := strconv.Atoi(a[1][12:])
	reg_c, _ := strconv.Atoi(a[2][12:])

	return registers{reg_a, reg_b, reg_c, 0}, strings.Split(a[4][9:], ",")
}

// Walk instructions backwards, run ops until we get an output and then check if the instruction had the same at that depth
func findA(instructions []string, a_val, depth int) int {
	// If we end up at under 0, all instructions have been output
	if depth < 0 {
		return a_val
	}
	for i := range 8 {
		checkablea := a_val*8 + i
		reg := registers{checkablea, 0, 0, 0}
		for reg.instructionPointer < len(instructions) {
			res := reg.doOp(instructions[reg.instructionPointer], instructions[reg.instructionPointer+1])
			if res == "" {
				continue
			}

			if res != instructions[depth] {
				break
			}

			if r := findA(instructions, checkablea, depth-1); r > 0 {
				return r
			}
		}
	}

	return 0
}

func solve(input string) (string, int) {
	reg, ins := parse(input)
	var sb strings.Builder
	//p1
	for reg.instructionPointer < len(ins) {
		ret := reg.doOp(ins[reg.instructionPointer], ins[reg.instructionPointer+1])
		if ret != "" {
			fmt.Fprintf(&sb, "%s,", ret)
		}
	}
	res := sb.String()

	return res[:len(res)-1], findA(ins, 0, len(ins)-1)
}

func main() {
	fmt.Println("Example result:")
	fmt.Println(solve(example))

	fmt.Println("Real:")
	fmt.Println(solve(aoc.GetInputFromFile("17")))
}

const example = `Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`

func intPow(n, m int) int {
	if m == 0 {
		return 1
	}

	if m == 1 {
		return n
	}

	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}
