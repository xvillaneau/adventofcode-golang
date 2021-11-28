package lib

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"strconv"
)

func ReadBytes (year int, day int) ([]byte, error) {
	ydir := fmt.Sprintf("aoc%v", year)
	fname := fmt.Sprintf("day%02d.txt", day)
	return os.ReadFile(path.Join("data", ydir, fname))
}

func ReadLines (year int, day int) ([]string, error) {
	var lines []string

	b, rderr := ReadBytes(year, day)
	if rderr != nil {
		return lines, rderr
	}

	for _, line := range(bytes.Split(b, []byte{'\n'})) {
		if len(line) != 0 {
			lines = append(lines, string(line))
		}
	}
	return lines, nil
}

func ReadNumbers (year int, day int) ([]int, error) {
	var nums []int

	lines, rderr := ReadLines(year, day)
	if rderr != nil {
		return nums, rderr
	}

	for _, line := range(lines) {
		n, perr := strconv.Atoi(string(line))
		if perr != nil {
			return nums, perr
		}
		nums = append(nums, n)
	}
	return nums, nil
}

