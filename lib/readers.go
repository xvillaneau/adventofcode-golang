package lib

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

func ReadBytes (year int, day int) ([]byte, error) {
	ydir := fmt.Sprintf("aoc%v", year)
	fname := fmt.Sprintf("day%02d.txt", day)
	return os.ReadFile(path.Join("data", ydir, fname))
}

func ReadFields (year int, day int, sep byte) ([]string, error) {
	var fields []string

	b, rderr := ReadBytes(year, day)
	if rderr != nil {
		return fields, rderr
	}

	for _, field := range(bytes.Split(b, []byte{sep})) {
		if len(field) != 0 {
			fields = append(fields, string(field))
		}
	}
	return fields, nil
}

func ReadLines (year int, day int) ([]string, error) {
	return ReadFields(year, day, byte('\n'))
}

func ReadNumbers (year int, day int, sep byte) ([]int, error) {
	var nums []int

	fields, rderr := ReadFields(year, day, sep)
	if rderr != nil {
		return nums, rderr
	}

	for _, field := range(fields) {
		n, perr := strconv.Atoi(strings.TrimSpace(string(field)))
		if perr != nil {
			return nums, perr
		}
		nums = append(nums, n)
	}
	return nums, nil
}

func ReadNumLn (year int, day int) ([]int, error) {
	return ReadNumbers(year, day, '\n')
}

func ReadNumComma (year int, day int) ([]int, error) {
	return ReadNumbers(year, day, ',')
}

