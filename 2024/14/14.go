package main

import (
	"aoc"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/stat"
)

type vector struct {
	x int
	y int
}

type robot struct {
	pos vector
	v   vector
}

func (r robot) getPositionAfter(n int, limit_x, limit_y int) vector {
	ret := vector{r.pos.x + n*r.v.x, r.pos.y + n*r.v.y}

	if ret.x >= 0 && ret.y >= 0 {
		ret.x = ret.x % limit_x
		ret.y = ret.y % limit_y
	} else if ret.x < 0 && ret.y < 0 {
		//fmt.Println((aoc.AbsInt(ret.x)-1)%limit_x, (aoc.AbsInt(ret.y)-1)%limit_y)
		//fmt.Println(aoc.AbsInt((aoc.AbsInt(ret.x)-1)%limit_x+1-limit_x), aoc.AbsInt((aoc.AbsInt(ret.y)-1)%limit_y+1-limit_y))
		ret.x = aoc.AbsInt((aoc.AbsInt(ret.x)-1)%limit_x + 1 - limit_x)
		ret.y = aoc.AbsInt((aoc.AbsInt(ret.y)-1)%limit_y + 1 - limit_y)
	} else if ret.y < 0 {
		ret.y = aoc.AbsInt((aoc.AbsInt(ret.y)-1)%limit_y + 1 - limit_y)
		ret.x = ret.x % limit_x
	} else {
		ret.x = aoc.AbsInt((aoc.AbsInt(ret.x)-1)%limit_x + 1 - limit_x)
		ret.y = ret.y % limit_y
	}

	return ret
}

func solvePart1(robots []robot, limit_x, limit_y int) int {
	quadrants := []int{0, 0, 0, 0}

	for _, r := range robots {
		pos := r.getPositionAfter(100, limit_x, limit_y)

		// topLeft
		if pos.x < limit_x/2 && pos.y < limit_y/2 {
			quadrants[0] += 1
		}

		// topright
		if pos.x > limit_x/2 && pos.y < limit_y/2 {
			quadrants[1] += 1
		}

		// botleft
		if pos.x < limit_x/2 && pos.y > limit_y/2 {
			quadrants[2] += 1
		}

		// botright
		if pos.x > limit_x/2 && pos.y > limit_y/2 {
			quadrants[3] += 1
		}

	}
	res := 1
	for _, q := range quadrants {
		res *= q
	}
	return res
}

func solvePart2(robots []robot, limit_x, limit_y int) int {
	x_res := make([]float64, len(robots))
	y_res := make([]float64, len(robots))
	best_x := math.MaxFloat64
	best_y := math.MaxFloat64
	tx := 0
	ty := 0

	for n := range limit_y {
		for i, r := range robots {
			pos := r.getPositionAfter(n, limit_x, limit_y)
			x_res[i] = float64(pos.x)
			y_res[i] = float64(pos.y)
		}
		var_x := stat.Variance(x_res, nil)
		var_y := stat.Variance(y_res, nil)

		if var_y < best_y {
			ty = n
			best_y = var_y
		}
		if var_x < best_x {
			tx = n
			best_x = var_x
		}
	}
	fmt.Println(tx, ty)
	iW := powUniversal(limit_x, -1, limit_y)
	t := tx + ((iW*(ty-tx))%limit_y)*limit_x
	//printRobotsAfterN(robots, t, limit_x, limit_y)
	return t
}

func parseRobots(input string) (robots []robot) {
	split := strings.Split(input, "\n")
	robots = make([]robot, len(split))
	for i, line := range split {
		re := regexp.MustCompile(`-?\d+`)
		parts := re.FindAll([]byte(line), -1)
		x, _ := strconv.Atoi(string(parts[0]))
		y, _ := strconv.Atoi(string(parts[1]))
		vx, _ := strconv.Atoi(string(parts[2]))
		vy, _ := strconv.Atoi(string(parts[3]))
		robot := robot{vector{x, y}, vector{vx, vy}}
		robots[i] = robot
	}
	return
}

func solve(input string, limit_x, limit_y int) (int, int) {
	robots := parseRobots(input)
	return solvePart1(robots, limit_x, limit_y), solvePart2(robots, limit_x, limit_y)
}

func printRobotsAfterN(robots []robot, n int, limit_x, limit_y int) {
	printCoords := make(map[vector]struct{})
	for _, r := range robots {
		printCoords[r.getPositionAfter(n, limit_x, limit_y)] = struct{}{}
	}
	for y := range limit_y {
		for x := range limit_x {
			if _, ok := printCoords[vector{x, y}]; ok {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println("Example result:")
	fmt.Println(solve(example, 11, 7))

	fmt.Println("Real:")
	fmt.Println(solve(aoc.GetInputFromFile("14"), 101, 103))
}

const example = `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

// Modular exponentiation
func powWithMod(a, exp, mod int) int {
	if exp < 0 {
		panic("Negative power given")
	}

	result := 1
	power := a

	for exp > 0 {
		if exp&1 == 1 {
			result = result * power % mod
		}
		power = power * power % mod
		exp >>= 1
	}

	return result
}

// Run modular exponentiation but also if integer is negative
func powUniversal(a, exp, mod int) int {
	if exp >= 0 {
		return powWithMod(a, exp, mod)
	}
	subResult := powWithMod(a, -exp, mod)
	return powWithMod(subResult, mod-2, mod)
}
