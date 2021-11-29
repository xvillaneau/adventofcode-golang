package main

import (
	"fmt"
	"github.com/xvillaneau/adventofcode-golang/aoc2019"
	"github.com/xvillaneau/adventofcode-golang/lib"
)

func main() {
	fmt.Println("AoC 2019 Day 2")

	prog, err := lib.ReadNumComma(2019, 2)
	if err != nil {
		fmt.Println(err)
		return
	}

	runprog := func (noun, verb int) int {
		comp := aoc2019.NewIntcode(prog)
		comp.Set(1, noun)
		comp.Set(2, verb)
		err = comp.RunFull()
		if err != nil {
			fmt.Println(err)
			return 0
		}
		res, _ := comp.Get(0)
		return res
	}

	fmt.Println("Part 1:", runprog(12, 2))
	for n:=0;n<100;n++ {
		for v:=0;v<100;v++ {
			res := runprog(n, v)
			if res == 19690720 {
				fmt.Println("Part 2:", n*100+v)
				return
			}
		}
	}
}

