package seven

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

type (
	bagInfo struct {
		bag    string
		amount int
	}
	rule struct {
		bagInfo  bagInfo
		contains []bagInfo
	}
)

// Solve implements solver interface for part one
func (d PartOneSolver) Solve() string {
	lines, err := common.GetLines("day/seven/input.txt")
	if err != nil {
		panic("couldn't open input file for day seven")
	}
	rules := mustParseRules(lines)
	return strconv.Itoa(solvePartOne(rules))
}

// Solve implements solver interface for part one
func (d PartTwoSolver) Solve() string {
	lines, err := common.GetLines("day/seven/input.txt")
	if err != nil {
		panic("couldn't open input file for day seven")
	}
	rules := mustParseRules(lines)
	return strconv.Itoa(solvePartTwo(rules))
}

func solvePartOne(rules []rule) int {
	ruleGraph := map[string][]string{}
	visited := map[string]bool{}
	for _, r := range rules {
		for _, b := range r.contains {
			v, ok := ruleGraph[b.bag]
			if !ok {
				ruleGraph[b.bag] = []string{}
			}
			ruleGraph[b.bag] = append(v, r.bagInfo.bag)
		}
	}
	return countBagContainers(ruleGraph, visited, "shiny gold")
}

func solvePartTwo(rules []rule) int {
	bagRules := map[string][]bagInfo{}
	for _, r := range rules {
		bagRules[r.bagInfo.bag] = r.contains
	}
	return countTotalBags(bagRules, "shiny gold", 1) - 1
}

func countBagContainers(bagRules map[string][]string, visited map[string]bool, name string) int {
	visited[name] = true
	v, _ := bagRules[name]
	if len(v) == 0 {
		return 0
	}
	count := 0
	for _, b := range v {
		if visited[b] != true {
			count += 1 + countBagContainers(bagRules, visited, b)
		}
	}
	return count
}

func countTotalBags(bagRules map[string][]bagInfo, name string, amount int) int {
	bags := bagRules[name]
	if len(bags) == 0 {
		return amount
	}
	sum := 0
	for _, b := range bags {
		sum += countTotalBags(bagRules, b.bag, b.amount)
	}
	return amount + (amount * sum)
}

func mustParseRules(lines []string) []rule {
	rules := []rule{}
	for _, line := range lines {
		rules = append(rules, mustParseLine(line))
	}
	return rules
}

func mustParseLine(line string) rule {
	sections := strings.Split(line, "bags contain ")
	subjectBag := mustParseBagInfo(sections[0])
	containedBags := []bagInfo{}
	splitBags := strings.Split(sections[1], ", ")
	if len(splitBags) == 1 {
		if strings.Contains(sections[1], "no other bags") {
			return rule{
				bagInfo:  subjectBag,
				contains: []bagInfo{},
			}
		}
	}
	for _, b := range splitBags {
		bag := mustParseBagInfo(b)
		containedBags = append(containedBags, bag)
	}
	return rule{
		bagInfo:  subjectBag,
		contains: containedBags,
	}
}

func mustParseBagInfo(s string) bagInfo {
	tokens := strings.Split(s, " ")
	if len(tokens) == 4 {
		amount, err := strconv.Atoi(tokens[0])
		if err != nil {
			panic("couldn't parse amount portion of bag rule")
		}
		return bagInfo{
			bag:    strings.Join(tokens[1:3], " "),
			amount: amount,
		}
	} else {
		return bagInfo{
			bag:    strings.Join(tokens[0:2], " "),
			amount: 1,
		}
	}
}
