package one

import (
	"fmt"
	"strconv"

	"github.com/mziter/aoc-2020/common"
)

type twoNums struct {
	num1 int
	num2 int
}

type threeNums struct {
	num1 int
	num2 int
	num3 int
}

type PartOneSolver struct{}
type PartTwoSolver struct{}

func (d PartOneSolver) Solve() string {
	return solvePartOne()
}

func (d PartTwoSolver) Solve() string {
	return solvePartTwo()
}

func solvePartOne() string {
	nums, err := common.GetIntLines("day/one/input.txt")
	if err != nil {
		panic(fmt.Errorf("Error reading file! %w", err))
	}
	twoNum := getTwoNumSum(nums, 2020)
	return strconv.Itoa(twoNum.num1 * twoNum.num2)
}

func solvePartTwo() string {
	nums, err := common.GetIntLines("day/one/input.txt")
	if err != nil {
		panic(fmt.Errorf("Error reading file! %w", err))
	}
	threeNum := getThreeNumSum(nums, 2020)
	return strconv.Itoa(threeNum.num1 * threeNum.num2 * threeNum.num3)
}

func getTwoNumSum(nums []int, target int) twoNums {
	numMap := make(map[int]bool, len(nums))
	for _, num := range nums {
		numMap[num] = true
	}
	for k := range numMap {
		other := target - k
		_, ok := numMap[other]
		if ok {
			return twoNums{k, other}
		}
	}
	panic(fmt.Sprintf("Could not find any two values that sum to %d", target))
}

func getThreeNumSum(nums []int, target int) threeNums {
	twoSums := make(map[int]twoNums, len(nums))
	for i, num1 := range nums {
		for j, num2 := range nums {
			if i != j {
				twoSums[num1+num2] = twoNums{num1, num2}
			}
		}
	}
	for _, n := range nums {
		other := target - n
		v, ok := twoSums[other]
		if ok {
			return threeNums{v.num1, v.num2, n}
		}
	}
	panic(fmt.Sprintf("Could not find any three values that sum to %d", target))
}
