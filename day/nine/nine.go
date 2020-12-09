package nine

import (
	"math"
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
	nums, err := common.GetIntLines("day/nine/input.txt")
	if err != nil {
		panic("could not open day nine input file")
	}
	answer := firstNonSum(nums, 25)
	return strconv.Itoa(answer)

}

// Solve implements solver interface for part one
func (d PartTwoSolver) Solve() string {
	nums, err := common.GetIntLines("day/nine/input.txt")
	if err != nil {
		panic("could not open day nine input file")
	}
	answer := contiguousSumMinMax(nums, 18272118)
	return strconv.Itoa(answer)
}

func firstNonSum(nums []int, lookback int) int {
	prevNums := map[int]bool{}
	leftIdx := 0
	rightIdx := 0
	for _, n := range nums {
		prevNums[n] = true
		reachedLookback := (rightIdx - leftIdx) == lookback
		if reachedLookback {
			foundSum := false
			for i := leftIdx; i < rightIdx; i++ {
				diff := n - nums[i]
				_, ok := prevNums[diff]
				if ok {
					foundSum = true
				}
			}
			if !foundSum {
				return n
			}
			delete(prevNums, nums[leftIdx])
			leftIdx++
		}
		rightIdx++
	}
	return -1
}

func contiguousSumMinMax(nums []int, target int) int {
	runningSum := []int{}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		runningSum = append(runningSum, sum)
	}

	for lIdx := range nums {
		for rIdx := range nums {
			if rIdx > lIdx {
				if runningSum[rIdx]-runningSum[lIdx] == target {
					min := math.MaxInt64
					max := math.MinInt64
					for i := lIdx + 1; i <= rIdx; i++ {
						if nums[i] < min {
							min = nums[i]
						}
						if nums[i] > max {
							max = nums[i]
						}
					}
					return min + max
				}
			}
		}
	}

	return -1
}
