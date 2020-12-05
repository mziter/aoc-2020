package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/mziter/aoc-2020/day/five"
	"github.com/mziter/aoc-2020/day/four"
	"github.com/mziter/aoc-2020/day/one"
	"github.com/mziter/aoc-2020/day/three"
	"github.com/mziter/aoc-2020/day/two"
)

type flags struct {
	day  int
	part int
}

type solver interface {
	Solve() string
}

func main() {
	flags := parseFlags()
	solvers := getSolvers()

	start := time.Now()
	solver := solvers[flags.day-1][flags.part-1]
	answer := solver.Solve()
	end := time.Now()

	fmt.Printf("Answer to day %d, part %d is: %s\n", flags.day, flags.part, answer)
	fmt.Printf("Solved in %s\n", end.Sub(start).String())
}

func getSolvers() [][]solver {
	var solvers [][]solver
	solvers = make([][]solver, 25)
	solvers[0] = []solver{one.PartOneSolver{}, one.PartTwoSolver{}}
	solvers[1] = []solver{two.PartOneSolver{}, two.PartTwoSolver{}}
	solvers[2] = []solver{three.PartOneSolver{}, three.PartTwoSolver{}}
	solvers[3] = []solver{four.PartOneSolver{}, four.PartTwoSolver{}}
	solvers[4] = []solver{five.PartOneSolver{}, five.PartTwoSolver{}}
	return solvers
}

func parseFlags() flags {
	dayFlag := flag.Int("day", 1, "day to run")
	partFlag := flag.Int("part", 1, "part of day to run")
	flag.Parse()
	if *partFlag < 1 || *partFlag > 2 {
		panic("Part must be either one or two")
	}
	return flags{*dayFlag, *partFlag}
}
