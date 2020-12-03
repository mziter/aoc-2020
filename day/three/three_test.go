package three

import (
	"testing"
)

func TestGetTreesEncountered(t *testing.T) {
	lines := exampleLines()
	treeLoc := parseTrees(lines)
	g := grid{width: len(lines[0]), height: len(lines)}
	s := slope{deltaY: 1, deltaX: 3}
	treeCount := getTreeEncountered(treeLoc, g, s)
	if treeCount != 7 {
		t.Errorf("Expected number of tree to be 7, but was %d", treeCount)
	}
}

func TestGetTreesEncounteredSlopes(t *testing.T) {
	lines := exampleLines()
	treeLoc := parseTrees(lines)
	g := grid{width: len(lines[0]), height: len(lines)}
	cases := []struct {
		s      slope
		output int
	}{
		{slope{deltaX: 1, deltaY: 1}, 2},
		{slope{deltaX: 3, deltaY: 1}, 7},
		{slope{deltaX: 5, deltaY: 1}, 3},
		{slope{deltaX: 7, deltaY: 1}, 4},
		{slope{deltaX: 1, deltaY: 2}, 2},
	}
	for _, c := range cases {
		treeCount := getTreeEncountered(treeLoc, g, c.s)
		if treeCount != c.output {
			t.Errorf("Expected number of tree to be %d, but was %d", c.output, treeCount)
		}
	}
}

func exampleLines() []string {
	return []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}
}
