package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xvillaneau/adventofcode-golang/lib"
)

type Instr string

const (
	Forward Instr = "forward"
	Down    Instr = "down"
	Up      Instr = "up"
)

type Move struct {
	instr Instr
	val   int
}

func parse(data []string) []Move {
	moves := make([]Move, len(data))
	for i, ln := range data {
		fs := strings.Split(ln, " ")
		val, _ := strconv.Atoi(fs[1])
		mv := Move{Instr(fs[0]), val}
		moves[i] = mv
	}
	return moves
}

func run(moves []Move) (int, int, int) {
	pos, depth, aim := 0, 0, 0
	for _, mv := range moves {
		switch mv.instr {
		case Forward:
			pos += mv.val
			depth += mv.val * aim
		case Down:
			aim += mv.val
		case Up:
			aim -= mv.val
		}
	}
	return pos, depth, aim
}

func main() {
	fmt.Println("Advent of Code 2021, Day 2")

	data, err := lib.ReadLines(2021, 2)
	if err != nil {
		fmt.Println(err)
		return
	}

	moves := parse(data)
	pos, depth, aim := run(moves)
	fmt.Println("Part 1:", pos*aim)
	fmt.Println("Part 2:", pos*depth)
}
