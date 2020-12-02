package two

import (
	"strconv"
	"strings"

	"github.com/mziter/aoc-2020/common"
)

type (
	// PartOneSolver implements solver interface for part one
	PartOneSolver struct{}
	// PartTwoSolver implements solver interface for part one
	PartTwoSolver struct{}
)

type policy struct {
	num1 int
	num2 int
}

// Solve implements solver interfae for part one
func (d PartOneSolver) Solve() string {
	return solvePartOne()
}

// Solve implements solver interfae for part one
func (d PartTwoSolver) Solve() string {
	return solvePartTwo()
}

func solvePartOne() string {
	lines, err := common.GetLines("day/two/input.txt")
	if err != nil {
		panic("Couldn't get lines from file")
	}
	count := 0
	for _, l := range lines {
		if validFreqLine(l) {
			count++
		}
	}
	return strconv.Itoa(count)
}

func solvePartTwo() string {
	lines, err := common.GetLines("day/two/input.txt")
	if err != nil {
		panic("Couldn't get lines from file")
	}
	count := 0
	for _, l := range lines {
		if validLocLine(l) {
			count++
		}
	}
	return strconv.Itoa(count)
}

func validFreqLine(s string) bool {
	tokens := strings.Split(s, " ")
	policy := parsePolicy(tokens[0])
	char := rune(tokens[1][0])
	pwd := tokens[2]
	return validFreq(pwd, policy.num1, policy.num2, char)
}

func validLocLine(s string) bool {
	tokens := strings.Split(s, " ")
	policy := parsePolicy(tokens[0])
	char := rune(tokens[1][0])
	pwd := tokens[2]
	return validLoc(pwd, policy.num1, policy.num2, char)
}

func parsePolicy(s string) policy {
	tokens := strings.Split(s, "-")
	least, err := strconv.Atoi(tokens[0])
	if err != nil {
		panic("Couldn't convert " + tokens[0] + " to int")
	}
	most, err := strconv.Atoi(tokens[1])
	if err != nil {
		panic("Couldn't convert " + tokens[1] + " to int")
	}
	return policy{least, most}
}

func validFreq(pwd string, least int, max int, char rune) bool {
	count := 0
	for _, r := range pwd {
		if r == char {
			count++
		}
		if count > max {
			return false
		}
	}
	if count < least {
		return false
	}
	return true
}

func validLoc(pwd string, idx1 int, idx2 int, char rune) bool {
	count := 0
	if rune(pwd[idx1-1]) == char {
		count++
	}
	if rune(pwd[idx2-1]) == char {
		count++
	}
	return count == 1
}
