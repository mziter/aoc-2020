package five

import (
	"math"
	"sort"
	"strconv"

	"github.com/mziter/aoc-2020/common"
)

type (
	// PartOneSolver implements solver interface for part one
	PartOneSolver struct{}
	// PartTwoSolver implements solver interface for part one
	PartTwoSolver struct{}
)

// Solve implements solver interface for part one
func (d PartOneSolver) Solve() string {
	lines, err := common.GetLines("day/five/input.txt")
	if err != nil {
		panic("couldn't open file for day 5")
	}
	max := 0
	for _, line := range lines {
		res := getBoardingPassID(line)
		if res > max {
			max = res
		}
	}
	return strconv.Itoa(max)
}

// Solve implements solver interface for part one
func (d PartTwoSolver) Solve() string {
	lines, err := common.GetLines("day/five/input.txt")
	if err != nil {
		panic("couldn't open file for day 5")
	}

	ids := []int{}
	for _, line := range lines {
		res := getBoardingPassID(line)
		ids = append(ids, res)
	}

	sort.Ints(ids)
	for i, id := range ids {
		if ids[i+1] != id+1 {
			return strconv.Itoa(id + 1)
		}
	}
	return "not found"
}

func getBoardingPassID(chars string) int {
	if len(chars) != 10 {
		panic("boarding pass has to be 10 chars long")
	}
	row := getBoardingNumber(chars[0:7])
	col := getBoardingNumber(chars[7:])
	return int((row * 8) + col)
}

func getBoardingNumber(chars string) float64 {
	if len(chars) == 0 {
		return 0
	}
	c := chars[0]
	if c == 'B' || c == 'R' {
		val := math.Pow(2, float64(len(chars)-1))
		return float64(val) + getBoardingNumber(chars[1:])
	}
	return getBoardingNumber(chars[1:])
}
