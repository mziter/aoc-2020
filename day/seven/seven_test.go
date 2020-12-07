package seven

import (
	"reflect"
	"testing"
)

func TestParseLine(t *testing.T) {
	cases := []struct {
		input  string
		output rule
	}{
		{"light red bags contain 1 bright white bag, 2 muted yellow bags.",
			rule{
				bagInfo: bagInfo{
					bag:    "light red",
					amount: 1,
				},
				contains: []bagInfo{
					{
						bag:    "bright white",
						amount: 1,
					},
					{
						bag:    "muted yellow",
						amount: 2,
					},
				},
			},
		},
		{"light red bags contain 1 bright white bag.",
			rule{
				bagInfo: bagInfo{
					bag:    "light red",
					amount: 1,
				},
				contains: []bagInfo{
					{
						bag:    "bright white",
						amount: 1,
					},
				},
			},
		},
		{"light red bags contain no other bags",
			rule{
				bagInfo: bagInfo{
					bag:    "light red",
					amount: 1,
				},
				contains: []bagInfo{},
			},
		},
	}
	for _, c := range cases {
		res := mustParseLine(c.input)
		if !reflect.DeepEqual(res, c.output) {
			t.Errorf("expected result of input: %s, to be %v, but was instead %v", c.input, c.output, res)
		}
	}
}

func TestParseBag(t *testing.T) {
	cases := []struct {
		input  string
		output bagInfo
	}{
		{"1 bright yellow bags",
			bagInfo{
				bag:    "bright yellow",
				amount: 1,
			},
		},
		{"dull red bag",
			bagInfo{
				bag:    "dull red",
				amount: 1,
			},
		},
		{"5 yuck brown bags",
			bagInfo{
				bag:    "yuck brown",
				amount: 5,
			},
		},
	}
	for _, c := range cases {
		res := mustParseBagInfo(c.input)
		if res != c.output {
			t.Errorf("expected result of input: %s, to be %v, but was instead %v", c.input, c.output, res)
		}
	}
}

func TestSolvePartOne(t *testing.T) {
	rules := getExampleOneRules()
	res := solvePartOne(rules)
	if res != 4 {
		t.Errorf("Expected answer to be 4 but was %d", res)
	}
}

func TestSolvePartTwoExampleOne(t *testing.T) {
	rules := getExampleOneRules()
	res := solvePartTwo(rules)
	if res != 32 {
		t.Errorf("Expected answer to be 32 but was %d", res)
	}
}

func TestSolvePartTwoExampleTwo(t *testing.T) {
	rules := getExampleTwoRules()
	res := solvePartTwo(rules)
	if res != 126 {
		t.Errorf("Expected answer to be 32 but was %d", res)
	}
}

func getExampleOneRules() []rule {
	return []rule{
		mustParseLine("light red bags contain 1 bright white bag, 2 muted yellow bags."),
		mustParseLine("dark orange bags contain 3 bright white bags, 4 muted yellow bags."),
		mustParseLine("bright white bags contain 1 shiny gold bag."),
		mustParseLine("muted yellow bags contain 2 shiny gold bags, 9 faded blue bags."),
		mustParseLine("shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags."),
		mustParseLine("dark olive bags contain 3 faded blue bags, 4 dotted black bags."),
		mustParseLine("vibrant plum bags contain 5 faded blue bags, 6 dotted black bags."),
		mustParseLine("faded blue bags contain no other bags."),
		mustParseLine("dotted black bags contain no other bags."),
	}

}

func getExampleTwoRules() []rule {
	return []rule{
		mustParseLine("shiny gold bags contain 2 dark red bags."),
		mustParseLine("dark red bags contain 2 dark orange bags."),
		mustParseLine("dark orange bags contain 2 dark yellow bags."),
		mustParseLine("dark yellow bags contain 2 dark green bags."),
		mustParseLine("dark green bags contain 2 dark blue bags."),
		mustParseLine("dark blue bags contain 2 dark violet bags."),
		mustParseLine("dark violet bags contain no other bags."),
	}
}
