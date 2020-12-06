package four

import (
	"regexp"
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
	lines, err := common.GetLines("day/four/input.txt")
	if err != nil {
		panic("Could not open input file for day four")
	}
	passports := common.SplitLines(lines)
	count := 0
	for _, p := range passports {
		if validPartOnePassport(p) {
			count++
		}
	}
	return strconv.Itoa(count)
}

// Solve implements solver interface for part one
func (d PartTwoSolver) Solve() string {
	lines, err := common.GetLines("day/four/input.txt")
	if err != nil {
		panic("Could not open input file for day four")
	}
	passports := common.SplitLines(lines)
	count := 0
	for _, p := range passports {
		if validPartTwoPassport(p) {
			count++
		}
	}
	return strconv.Itoa(count)
}

func validPartOnePassport(lines []string) bool {
	fieldCount := 0
	foundCID := false
	for _, line := range lines {
		lineTokens := strings.Split(line, " ")
		fieldCount = fieldCount + len(lineTokens)
		for _, lineToken := range lineTokens {
			fieldTokens := strings.Split(lineToken, ":")
			k := fieldTokens[0]
			if k == "cid" {
				foundCID = true
			}
		}
	}
	switch {
	case fieldCount == 8:
		return true
	case fieldCount == 7 && !foundCID:
		return true
	default:
		return false
	}
}

func validPartTwoPassport(lines []string) bool {
	fieldCount := 0
	foundCid := false
	for _, line := range lines {
		lineTokens := strings.Split(line, " ")
		fieldCount = fieldCount + len(lineTokens)
		for _, lineToken := range lineTokens {
			fieldTokens := strings.Split(lineToken, ":")
			k := fieldTokens[0]
			if k == "cid" {
				foundCid = true
			}
			v := fieldTokens[1]
			if !isValidField(k, v) {
				return false
			}
		}
	}
	switch {
	case fieldCount == 8:
		return true
	case fieldCount == 7 && !foundCid:
		return true
	default:
		return false
	}
}

func isValidField(key string, value string) bool {
	switch key {
	case "byr":
		yr := mustConvertToInt(value)
		return inRange(yr, 1920, 2002)
	case "iyr":
		yr := mustConvertToInt(value)
		return inRange(yr, 2010, 2020)
	case "eyr":
		yr := mustConvertToInt(value)
		return inRange(yr, 2020, 2030)
	case "hgt":
		return isValidHeight(value)
	case "hcl":
		return isValidHairColor(value)
	case "ecl":
		return isValidEyeColor(value)
	case "pid":
		return isValidPID(value)
	case "cid":
		return true
	default:
		return false
	}
}

func isValidHairColor(value string) bool {
	re := regexp.MustCompile(`^#[a-f\d]{6}$`)
	return re.Match([]byte(value))
}

func isValidPID(value string) bool {
	re := regexp.MustCompile(`^[\d]{9}$`)
	return re.Match([]byte(value))
}

func isValidEyeColor(value string) bool {
	validEyes := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, validEye := range validEyes {
		if value == validEye {
			return true
		}
	}
	return false
}

func isValidHeight(value string) bool {
	i := strings.Index(value, "in")
	isCm := false
	if i == -1 {
		i = strings.Index(value, "cm")
		if i == -1 {
			return false
		}
		isCm = true
	}

	h := mustConvertToInt(value[0:i])
	if isCm {
		return inRange(h, 150, 293)
	}
	return inRange(h, 59, 76)
}

func mustConvertToInt(value string) int {
	n, err := strconv.Atoi(value)
	if err != nil {
		panic("Error converting value of " + value + " to integer")
	}
	return n
}

func inRange(value, lower, upper int) bool {
	return value >= lower && value <= upper
}
