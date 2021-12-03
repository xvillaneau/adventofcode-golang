package main

import (
	"fmt"
	"github.com/xvillaneau/adventofcode-golang/lib"
)

func increases (depths []int, start int) int {
	sum := 0
	for i, d := range depths[start:] {
		if d > depths[i] {
			sum++
		}
	}
	return sum
}

func main () {
	fmt.Println("Advent of Code 2021, Day 1")

	depths, err := lib.ReadNumLn(2021, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Part 1:", increases(depths, 1))
	fmt.Println("Part 2:", increases(depths, 3))
}
