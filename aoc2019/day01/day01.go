package main

import (
	"fmt"
	"github.com/xvillaneau/adventofcode-golang/lib"
)

func fuelReq(mass int) int {
	return (mass / 3) - 2
}

func totFuelReq(mass int) int {
	fuel := 0
	for {
		f := fuelReq(mass)
		if f <= 0 {
			break
		}
		fuel = fuel + f
		mass = f
	}
	return fuel
}

func main() {
	fmt.Println("AoC 2019 Day 1")

	mods, err := lib.ReadNumLn(2019, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fuel, totFuel := 0, 0
	for _, mass := range(mods) {
		fuel = fuel + fuelReq(mass)
		totFuel = totFuel + totFuelReq(mass)
	}
	fmt.Println("Part 1:", fuel)
	fmt.Println("Part 2:", totFuel)
}
