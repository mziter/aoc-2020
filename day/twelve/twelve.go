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

type waypoint struct {
	deltaX int
	deltaY int
}

type coord struct {
	x int
	y int
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
	instructions, err := common.GetLines("day/twelve/input.txt")
	if err != nil {
		panic("couldn't read input file for day twelve")
	}

	c := coord{x: 0, y: 0}
	wp := waypoint{deltaX: 10, deltaY: 1}
	for _, i := range instructions {
		c, wp = executeWaypointInstruction(i, c, wp)
	}

	deltaX := int(math.Abs(float64(c.x)))
	deltaY := int(math.Abs(float64(c.y)))

	return strconv.Itoa(deltaX + deltaY)
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

func executeWaypointInstruction(instruction string, c coord, wp waypoint) (coord, waypoint) {
	action := instruction[0]
	strVal := instruction[1:]
	value, err := strconv.Atoi(strVal)
	if err != nil {
		panic("couldn't convert string to number")
	}

	switch action {
	// adjust waypoint by amount
	case 'N':
		return c, waypoint{deltaX: wp.deltaX, deltaY: wp.deltaY + value}
	case 'S':
		return c, waypoint{deltaX: wp.deltaX, deltaY: wp.deltaY - value}
	case 'E':
		return c, waypoint{deltaX: wp.deltaX + value, deltaY: wp.deltaY}
	case 'W':
		return c, waypoint{deltaX: wp.deltaX - value, deltaY: wp.deltaY}
	// rotations
	case 'R':
		return c, rotateRight(wp, value)
	case 'L':
		return c, rotateLeft(wp, value)
	// execute waypoint N (value) times
	case 'F':
		return coord{x: c.x + (value * wp.deltaX), y: c.y + (value * wp.deltaY)}, wp
	default:
		panic("encountered a bad instruction")
	}
}

func rotateLeft(wp waypoint, degrees int) waypoint {
	return rotateRight(wp, 360-degrees)
}

func rotateRight(wp waypoint, degrees int) waypoint {
	d := degrees % 360
	for d > 0 {
		wp = waypoint{deltaX: wp.deltaY, deltaY: -(wp.deltaX)}
		d -= 90
	}
	return wp
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
