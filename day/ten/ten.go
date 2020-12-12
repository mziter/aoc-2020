package ten

import (
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
	nums, err := common.GetIntLines("day/ten/input.txt")
	if err != nil {
		panic("couldn't read input file for day ten")
	}
	answer := oneAndThreeJoltProduct(nums)
	return strconv.Itoa(answer)
}

// Solve implements solver interface for part one
func (d PartTwoSolver) Solve() string {
	nums, err := common.GetIntLines("day/ten/input.txt")
	if err != nil {
		panic("couldn't read input file for day ten")
	}
	answer := possibilities(nums)
	return strconv.Itoa(answer)
}

func possibilities(nums []int) int {
	sort.Ints(nums)
	adapters := append([]int{0}, nums...)
	dp := make([]int, len(adapters))

	dp[0] = 1
	dp[1] = 1

	for i := 2; i < len(dp); i++ {
		for j := i - 1; j >= 0; j-- {
			if adapters[i]-adapters[j] <= 3 {
				dp[i] += dp[j]
			} else {
				break
			}
		}
	}

	return dp[len(dp)-1]
}

func oneAndThreeJoltProduct(nums []int) int {
	sort.Ints(nums)

	oneCount := 0
	threeCount := 1 // built in adapter is always +3

	last := 0 // outlet
	for _, n := range nums {
		if last+3 == n {
			threeCount++
		}
		if last+1 == n {
			oneCount++
		}
		last = n
	}

	return oneCount * threeCount
}
