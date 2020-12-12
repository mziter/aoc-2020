package twelve

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/mziter/aoc-2020/common"
)

type position struct {
	x      int
	y      int
	facing rune
}

type (
	// PartOneSolver implements solver interface for part one
	PartOneSolver struct{}
	// PartTwoSolver implements solver interface for part one
	PartTwoSolver struct{}
)

// Solve implements solver interface for part one
func (d PartOneSolver) Solve() string {
	instructions, err := common.GetLines("day/twelve/input.txt")
	if err != nil {
		panic("couldn't read input file for day twelve")
	}

	position := position{x: 0, y: 0, facing: 'E'}
	for _, i := range instructions {
		position = executeInstruction(i, position)
	}

	deltaX := int(math.Abs(float64(position.x)))
	deltaY := int(math.Abs(float64(position.y)))

	return strconv.Itoa(deltaX + deltaY)
}

// Solve implements solver interface for part one
func (d PartTwoSolver) Solve() string {
	return "not implemented"
}

func executeInstruction(instruction string, p position) position {
	action := instruction[0]
	strVal := instruction[1:]
	value, err := strconv.Atoi(strVal)
	if err != nil {
		panic("couldn't convert string to number")
	}

	switch action {
	case 'N':
		return position{p.x, p.y + value, p.facing}
	case 'S':
		return position{p.x, p.y - value, p.facing}
	case 'E':
		return position{p.x + value, p.y, p.facing}
	case 'W':
		return position{p.x - value, p.y, p.facing}
	case 'R':
		return position{p.x, p.y, turnRight(p.facing, value)}
	case 'L':
		return position{p.x, p.y, turnLeft(p.facing, value)}
	case 'F':
		return executeInstruction(fmt.Sprintf("%c%s", p.facing, strVal), p)
	default:
		panic("encountered a bad instruction")
	}
}

func turnLeft(initial rune, degree int) rune {
	return turnRight(initial, 360-degree)
}

func turnRight(initial rune, degree int) rune {
	d := degree % 360
	facings := "NESW"

	initialIdx := strings.Index(facings, string(initial))
	changeIdx := d / 90
	newIdx := (initialIdx + changeIdx) % 4

	return rune(facings[newIdx])
}
