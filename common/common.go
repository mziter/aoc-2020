package common

import (
	"bufio"
	"os"
	"strconv"
)

func GetLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		text := scanner.Text()
		lines = append(lines, text)
	}
	return lines, nil
}

func GetIntLines(filename string) ([]int, error) {
	lines, err := GetLines(filename)
	if err != nil {
		return nil, err
	}
	var nums []int
	for _, l := range lines {
		num, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}
	return nums, nil
}
