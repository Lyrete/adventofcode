package main

import (
	"aoc"
	"fmt"
	"strconv"
	"strings"
)

/** Find unique solution to the problem x(ai​+bj​)+y(ci​+dj​)=ei​+fj using Cramer's rule
 */
func findUniqueSolution(a, b, c, d, e, f int) (int, int) {
	// 	The equation can be presented as a simple system of
	// ax + cy = e
	// bx + dy = f
	// And such in matrix form
	// {a c} {x} = {e}
	// {b d} {y}   {f}
	D := a*d - b*c
	// If the determinant of the first part is != 0 there is only one solution to the equation
	// If it is 0 there is either infinite or no solutions but we don't really have to consider those cases here
	if D != 0 {
		// Get the adjugate matrix of x and y separaetely (which is also their determinant)
		Dx, Dy := (e*d - c*f), (a*f - e*b)
		// By multiplying it with the inverse of the original determinant we can find the solution
		x, y := Dx/D, Dy/D
		// Our results are likely to not actually be natural numbers (but as we're doing int division they get floored to such)
		// So make sure that the proposed solution is genuinely valid by placing the calculated values in the equation
		if x*a+y*c == e && x*b+y*d == f {
			return x, y
		}
	}
	// The infinite cases dont ever occur as there are no negative x or y values so if the solution wasn't found there is no solution
	return 0, 0
}

func parseMove(input string) (int, int) {
	a := strings.Split(input, "+")
	x, _ := strconv.Atoi(a[1][:len(a[1])-3])
	y, _ := strconv.Atoi(a[2])
	return x, y
}

func parseGoal(input string) (int, int) {
	a := strings.Split(input, "=")
	x, _ := strconv.Atoi(a[1][:len(a[1])-3])
	y, _ := strconv.Atoi(a[2])
	return x, y
}

func solve(input string) (int, int) {
	res, res2 := 0, 0
	for _, gameStr := range strings.Split(input, "\n\n") {
		rows := strings.Split(gameStr, "\n")
		a_x, a_y := parseMove(rows[0])
		b_x, b_y := parseMove(rows[1])
		goal_x, goal_y := parseGoal(rows[2])
		a_press, b_press := findUniqueSolution(a_x, a_y, b_x, b_y, goal_x, goal_y)
		res += 3*a_press + b_press
		a_press, b_press = findUniqueSolution(a_x, a_y, b_x, b_y, 10000000000000+goal_x, 10000000000000+goal_y)
		res2 += 3*a_press + b_press

	}
	return res, res2
}

func main() {
	fmt.Println("Example result:")
	fmt.Println(solve(example))

	fmt.Println("Real:")
	fmt.Println(solve(aoc.GetInputFromFile("13")))
}

const example = `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`
