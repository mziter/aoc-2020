package fourteen

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

// Solve implements solver interface for part one
func (d PartOneSolver) Solve() string {
	lines, err := common.GetLines("day/fourteen/input.txt")
	if err != nil {
		panic("couldn't read file for day fourteen")
	}

	memory := map[int]int{}
	mask := [36]int{}

	for _, l := range lines {
		isMask := strings.HasPrefix(l, "mask")
		tokens := strings.Split(l, " = ")
		if isMask {
			mask = parseMask(tokens[1])
		} else {
			loc, err := strconv.Atoi(tokens[0][4 : len(tokens[0])-1])
			if err != nil {
				panic("failed trying to parse mem location to integer")
			}
			val, err := strconv.Atoi(tokens[1])
			if err != nil {
				panic("failed trying to parse mem value to integer")
			}
			memory[loc] = applyMask(mask, val)
		}
	}

	sum := 0
	for _, v := range memory {
		sum = sum + v
	}

	return strconv.Itoa(sum)
}

// Solve implements solver interface for part one
func (d PartTwoSolver) Solve() string {
	return "not implemented"
}

func parseMask(str string) [36]int {
	mask := [36]int{}
	for i, c := range str {
		idx := 35 - i
		switch c {
		case 'X':
			mask[idx] = -1
		case '0':
			mask[idx] = 0
		case '1':
			mask[idx] = 1
		default:
			panic("encountered a bad character when parsing mask")
		}
	}
	return mask
}

func applyMask(mask [36]int, n int) int {
	num := n
	for i, m := range mask {
		switch m {
		case 1:
			num = setBit(num, uint(i))
		case 0:
			num = clearBit(num, uint(i))
		}
	}
	return num
}

// Sets the bit at pos in the integer n.
func setBit(n int, pos uint) int {
	n |= (1 << pos)
	return n
}

// Clears the bit at pos in n.
func clearBit(n int, pos uint) int {
	mask := ^(1 << pos)
	n &= mask
	return n
}
