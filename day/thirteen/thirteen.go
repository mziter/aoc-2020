package thirteen

import (
	"fmt"
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

type busInfo struct {
	offset int
	id     int
}

// Solve implements solver interface for part one
func (d PartOneSolver) Solve() string {
	lines, err := common.GetLines("day/thirteen/input.txt")
	if err != nil {
		panic("couldn't read input file for day thirteen")
	}

	time, err := strconv.Atoi(lines[0])
	if err != nil {
		panic("couldn't convert first line to integer")
	}

	buses := parseLineNumsIgnore(lines[1])
	bus, waitTime := findFirstDeparture(buses, time)
	return strconv.Itoa(bus * waitTime)
}

// Solve implements solver interface for part one
func (d PartTwoSolver) Solve() string {
	lines, err := common.GetLines("day/thirteen/input.txt")
	if err != nil {
		panic("couldn't read input file for day thirteen")
	}

	buses := parseLineBusInfo(lines[1])
	return fmt.Sprintf("%v", findFirstAligned(buses))
}

func findFirstDeparture(buses []int, arrivalTime int) (int, int) {
	time := arrivalTime
	for {
		for _, b := range buses {
			if time%b == 0 {
				return b, time - arrivalTime
			}
		}
		time++
	}
}

func findFirstAligned(buses []busInfo) int64 {
	max := getMaxBus(buses)
	time := (int64(max.id) * 1000000000000) - int64(max.offset)
	found := false
	for !found {
		found = true
		for _, b := range buses {
			totalTime := time + int64(b.offset)
			mod := totalTime % int64(b.id)
			if mod != 0 {
				found = false
				break
			}
		}
		if found {
			return time
		}
		time = time + int64(max.id)
	}
	panic("something bad happened")
}

func parseLineNumsIgnore(s string) []int {
	trimmed := strings.Trim(s, " \n")
	tokens := strings.Split(trimmed, ",")
	nums := []int{}
	for _, n := range tokens {
		num, err := strconv.Atoi(n)
		if err != nil {
			continue
		}
		nums = append(nums, num)
	}
	return nums
}

func parseLineBusInfo(s string) []busInfo {
	trimmed := strings.Trim(s, " \n")
	tokens := strings.Split(trimmed, ",")
	nums := []busInfo{}
	for i, n := range tokens {
		num, err := strconv.Atoi(n)
		if err == nil {
			nums = append(nums, busInfo{offset: i, id: num})
		}
	}
	return nums
}

func getMaxBus(buses []busInfo) busInfo {
	var max busInfo
	for _, b := range buses {
		if b.id > max.id {
			max = b
		}
	}
	return max
}
