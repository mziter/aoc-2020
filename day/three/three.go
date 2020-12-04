package three

import (
	"strconv"

	"github.com/mziter/aoc-2020/common"
)

type coordinate struct {
	x int
	y int
}

type slope struct {
	deltaX int
	deltaY int
}

type grid struct {
	width  int
	height int
}

type (
	// PartOneSolver implements solver interface for part one
	PartOneSolver struct{}
	// PartTwoSolver implements solver interface for part one
	PartTwoSolver struct{}
)

// Solve implements solver interface for part one
func (d PartOneSolver) Solve() string {
	lines, err := common.GetLines("day/three/input.txt")
	if err != nil {
		panic("Could not open input file for day three")
	}
	treeLoc := parseTrees(lines)
	g := grid{width: len(lines[0]), height: len(lines)}
	s := slope{deltaY: 1, deltaX: 3}
	treeCount := getTreeEncountered(treeLoc, g, s)
	return strconv.Itoa(treeCount)
}

// Solve implements solver interface for part one
func (d PartTwoSolver) Solve() string {
	lines, err := common.GetLines("day/three/input.txt")
	if err != nil {
		panic("Could not open input file for day three")
	}
	treeLoc := parseTrees(lines)
	g := grid{width: len(lines[0]), height: len(lines)}
	slopes := []slope{
		{deltaX: 1, deltaY: 1},
		{deltaX: 3, deltaY: 1},
		{deltaX: 5, deltaY: 1},
		{deltaX: 7, deltaY: 1},
		{deltaX: 1, deltaY: 2},
	}
	treeCount := 1
	for _, s := range slopes {
		trees := getTreeEncountered(treeLoc, g, s)
		treeCount = treeCount * trees
	}
	return strconv.Itoa(treeCount)
}

func getTreeEncountered(treeLoc map[coordinate]bool, g grid, s slope) int {
	x := 0
	y := 0
	treeCount := 0
	for y < g.height {
		x = (x + s.deltaX) % g.width
		y = y + s.deltaY
		_, ok := treeLoc[coordinate{x: x, y: y}]
		if ok {
			treeCount++
		}
	}
	return treeCount
}

func parseTrees(lines []string) map[coordinate]bool {
	treeLoc := make(map[coordinate]bool)
	for y, line := range lines {
		for x, char := range line {
			if rune(char) == '#' {
				treeLoc[coordinate{x: x, y: y}] = true
			}
		}
	}
	return treeLoc
}
