package common

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
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

func GetBytesFromFile(filename string) ([][]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return GetBytes(file)
}

func GetBytes(r io.Reader) ([][]byte, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return bytes.Split(data, []byte("\n")), nil
}

func SplitLines(lines []string) [][]string {
	splits := make([][]string, 0)
	currentSplit := make([]string, 0)
	for _, line := range lines {
		if line == "" {
			splits = append(splits, currentSplit)
			currentSplit = make([]string, 0)
		} else {
			currentSplit = append(currentSplit, line)
		}
	}
	if len(currentSplit) > 0 {
		splits = append(splits, currentSplit)
	}
	return splits
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
