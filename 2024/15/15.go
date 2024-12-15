package main

import (
	"aoc"
	"fmt"
	"strings"
)

type robot struct {
	pos aoc.Coord
}

type box struct {
	l aoc.Coord
	r aoc.Coord
}

func (b box) AddXy(x, y int) box {
	return box{l: b.l.AddXY(x, y), r: b.r.AddXY(x, y)}
}

func (b *box) canMove(dx, dy int, grid [][]rune) bool {
	nextPos := b.AddXy(dx, dy)

	if grid[nextPos.l.Y][nextPos.l.X] == '#' || grid[nextPos.r.Y][nextPos.r.X] == '#' {
		return false
	}
	if dy != 0 {
		if grid[nextPos.l.Y][nextPos.l.X] == '.' && grid[nextPos.r.Y][nextPos.r.X] == '.' {
			return true
		}
		if grid[nextPos.l.Y][nextPos.l.X] == '[' && grid[nextPos.r.Y][nextPos.r.X] == ']' {
			pushableBox := nextPos
			return pushableBox.canMove(dx, dy, grid)
		}
		if grid[nextPos.l.Y][nextPos.l.X] == ']' && grid[nextPos.r.Y][nextPos.r.X] == '[' {
			lBox := box{r: nextPos.l, l: nextPos.l.AddXY(-1, 0)}
			rBox := box{l: nextPos.r, r: nextPos.r.AddXY(1, 0)}
			return lBox.canMove(dx, dy, grid) && rBox.canMove(dx, dy, grid)
		}
		if grid[nextPos.l.Y][nextPos.l.X] == ']' && grid[nextPos.r.Y][nextPos.r.X] == '.' {
			lBox := box{r: nextPos.l, l: nextPos.l.AddXY(-1, 0)}
			return lBox.canMove(dx, dy, grid)
		}
		if grid[nextPos.l.Y][nextPos.l.X] == '.' && grid[nextPos.r.Y][nextPos.r.X] == '[' {
			rBox := box{l: nextPos.r, r: nextPos.r.AddXY(1, 0)}
			return rBox.canMove(dx, dy, grid)
		}
	} else {
		if grid[nextPos.l.Y][nextPos.l.X] == '.' || grid[nextPos.r.Y][nextPos.r.X] == '.' {
			return true
		}
		pushableBox := nextPos.AddXy(dx, dy)
		if grid[pushableBox.l.Y][pushableBox.l.X] == '[' && grid[pushableBox.r.Y][pushableBox.r.X] == ']' {
			return pushableBox.canMove(dx, dy, grid)
		}
	}

	return false
}

func (b *box) drawToGrid(grid [][]rune) {
	grid[b.l.Y][b.l.X] = '['
	grid[b.r.Y][b.r.X] = ']'
}

func (b *box) removeFromGrid(grid [][]rune) {
	grid[b.l.Y][b.l.X] = '.'
	grid[b.r.Y][b.r.X] = '.'
}

func (b *box) move(dx, dy int, grid [][]rune) bool {
	if !b.canMove(dx, dy, grid) {
		return false
	}

	nextPos := b.AddXy(dx, dy)

	if dx != 0 {
		pushableBox := nextPos.AddXy(dx, dy)
		if grid[pushableBox.l.Y][pushableBox.l.X] == '[' && grid[pushableBox.r.Y][pushableBox.r.X] == ']' {

			pushableBox.move(dx, dy, grid)
		}
	} else {
		if grid[nextPos.l.Y][nextPos.l.X] == '[' && grid[nextPos.r.Y][nextPos.r.X] == ']' {
			pushableBox := nextPos
			pushableBox.move(dx, dy, grid)
		}

		if grid[nextPos.l.Y][nextPos.l.X] == ']' {
			lBox := box{l: nextPos.l.AddXY(-1, 0), r: nextPos.l}
			lBox.move(dx, dy, grid)
		}

		if grid[nextPos.r.Y][nextPos.r.X] == '[' {
			rBox := box{l: nextPos.r, r: nextPos.r.AddXY(1, 0)}
			rBox.move(dx, dy, grid)
		}
	}

	b.removeFromGrid(grid)
	nextPos.drawToGrid(grid)
	return true
}

func (r *robot) move(dx, dy int, grid [][]rune) {
	target := r.pos.AddXY(dx, dy)
	origTarget := target
	for grid[target.Y][target.X] != '#' {
		if grid[target.Y][target.X] == '.' {
			temp := grid[origTarget.Y][origTarget.X]
			grid[origTarget.Y][origTarget.X] = '@'
			grid[r.pos.Y][r.pos.X] = '.'
			r.pos.Set(origTarget.X, origTarget.Y)
			if temp == 'O' {
				grid[target.Y][target.X] = temp
			}
			break
		}
		target.MoveXY(dx, dy)
	}
}

func (r *robot) move2(dx, dy int, grid [][]rune) {
	target := r.pos.AddXY(dx, dy)
	if grid[target.Y][target.X] == '.' {
		grid[r.pos.Y][r.pos.X] = '.'
		r.pos.MoveXY(dx, dy)
		grid[r.pos.Y][r.pos.X] = '@'
		return
	}
	//origTarget := target
	if grid[target.Y][target.X] == '[' {
		pushedBox := box{l: target, r: target.AddXY(1, 0)}
		if pushedBox.move(dx, dy, grid) {
			grid[r.pos.Y][r.pos.X] = '.'
			r.pos.MoveXY(dx, dy)
			grid[r.pos.Y][r.pos.X] = '@'
			return
		}
	} else if grid[target.Y][target.X] == ']' {
		pushedBox := box{l: target.AddXY(-1, 0), r: target}
		if pushedBox.move(dx, dy, grid) {
			grid[r.pos.Y][r.pos.X] = '.'
			r.pos.MoveXY(dx, dy)
			grid[r.pos.Y][r.pos.X] = '@'
			return
		}
	}
}

func parseMove(r rune) (dx int, dy int) {
	switch r {
	case '<':
		return -1, 0
	case '>':
		return 1, 0
	case '^':
		return 0, -1
	case 'v':
		return 0, 1
	default:
		fmt.Println(r)
		panic("Unrecognized move:")
	}
}

func parseMap(input string) (locs [][]rune, start robot) {
	lines := strings.Split(input, "\n")
	ret := make([][]rune, len(lines))
	for y, line := range lines {
		row := make([]rune, len(line))
		for x, r := range line {
			if r == '@' {
				start.pos.Set(x, y)
			}
			row[x] = r
		}
		ret[y] = row
	}

	return ret, start

}

func parseMap2(input string) (locs [][]rune, start robot) {
	lines := strings.Split(input, "\n")
	ret := make([][]rune, len(lines))
	for y, line := range lines {
		row := make([]rune, len(line)*2)
		for x, r := range line {
			switch r {
			case '@':
				row[x*2] = r
				row[x*2+1] = '.'
				start.pos.Set(x*2, y)
			case '.':
				row[x*2] = r
				row[x*2+1] = r
			case '#':
				row[x*2] = r
				row[x*2+1] = r
			case 'O':
				row[x*2] = '['
				row[x*2+1] = ']'
			}
		}
		ret[y] = row
	}

	return ret, start

}

func solve(input string) (int, int) {
	split := strings.Split(input, "\n\n")
	locations, start := parseMap(split[0])
	instructions := strings.ReplaceAll(split[1], "\n", "")
	locations2, start2 := parseMap2(split[0])
	for _, move := range instructions {
		dx, dy := parseMove(move)
		start.move(dx, dy, locations)
		start2.move2(dx, dy, locations2)
	}
	return getCoordinateSum(locations, 1, 100), getCoordinateSum(locations2, 1, 100)
}

func getCoordinateSum(locs [][]rune, weight_x, weight_y int) int {
	sum := 0
	for y, line := range locs {
		for x, r := range line {
			if r == 'O' || r == '[' {
				sum += weight_x*x + weight_y*y
			}
		}
	}
	return sum
}

func printMap(locs [][]rune) {
	for _, line := range locs {
		for _, r := range line {
			fmt.Print(string(r))
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println("Example result:")
	fmt.Println(solve(example))

	fmt.Println("Real:")
	fmt.Println(solve(aoc.GetInputFromFile("15")))
}

const example = `##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`
