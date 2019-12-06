package day01

import (
	"math"
	"strconv"

	"github.com/jonasferry/advent2019/util"
)

func fuel(m int) int {
	// fuel = floor(mass / 3) - 2
	return int(math.Floor(float64(m)/3) - 2)
}

func fuel1(m int) int {
	// Fuel for modules
	return fuel(m)
}

func fuel2(m int) int {
	// Fuel for modules plus the fuel itself
	modfuel := fuel(m)
	totalfuel := modfuel
	fuelfuel := fuel(modfuel)

	if fuelfuel > 0 {
		for {
			totalfuel = totalfuel + fuelfuel
			if fuel(fuelfuel) <= 0 {
				break
			}
			fuelfuel = fuel(fuelfuel)
		}
	}

	return totalfuel
}

func Run1() int {
	// Read input file
	lines, err := util.ReadF("day01/input.txt")
	if err != nil {
		return -1
	}

	totalfuel := 0
	for i := 0; i < len(lines); i++ {
		f, _ := strconv.Atoi(lines[i])
		totalfuel = totalfuel + fuel1(f)
	}

	return totalfuel
}

func Run2() int {
	// Read input file
	lines, err := util.ReadF("day01/input.txt")
	if err != nil {
		return -1
	}

	totalfuel := 0
	for i := 0; i < len(lines); i++ {
		f, _ := strconv.Atoi(lines[i])
		totalfuel = totalfuel + fuel2(f)
	}

	return totalfuel

}
