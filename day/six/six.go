package six

import (
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
	lines, err := common.GetLines("day/six/input.txt")
	if err != nil {
		return "Error access input file for day six"
	}
	resp := common.SplitLines(lines)
	sum := 0
	for _, groupResp := range resp {
		sum += countDistinctChars(groupResp)
	}
	return strconv.Itoa(sum)
}

// Solve implements solver interface for part one
func (d PartTwoSolver) Solve() string {
	lines, err := common.GetLines("day/six/input.txt")
	if err != nil {
		return "Error access input file for day six"
	}
	resp := common.SplitLines(lines)
	sum := 0
	for _, groupResp := range resp {
		sum += countAllYes(groupResp)
	}
	return strconv.Itoa(sum)
}

func countDistinctChars(groupResp []string) int {
	yesTracker := map[rune]bool{}
	for _, pResp := range groupResp {
		for _, resp := range pResp {
			yesTracker[resp] = true
		}
	}
	return len(yesTracker)
}

func countAllYes(groupResp []string) int {
	yesTracker := map[rune]int{}
	for _, pResp := range groupResp {
		for _, resp := range pResp {
			count, ok := yesTracker[resp]
			if !ok {
				yesTracker[resp] = 1
			}
			yesTracker[resp] = count + 1
		}
	}
	count := 0
	for _, v := range yesTracker {
		if v == len(groupResp) {
			count++
		}
	}
	return count
}
