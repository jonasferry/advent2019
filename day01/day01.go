package day01

import (
	"fmt"
	"math"
	"strconv"

	"github.com/jonasferry/advent2019/util"
)

func fuel(m int) int {
	// fuel = floor(mass / 3) - 2
	return int(math.Floor(float64(m)/3) - 2)
}

func fuel1(m int) int {
	// fuel = floor(mass / 3) - 2
	return fuel(m)
}

func fuel2(m int) int {
	// fuel = floor(mass / 3) - 2
	modfuel := fuel(m)
	totalfuel := modfuel
	fuelfuel := fuel(modfuel)

	if fuelfuel > 0 {
		for {
			//fmt.Println("totalfuel: ", totalfuel, "fuelfuel: ", fuelfuel)
			totalfuel = totalfuel + fuelfuel
			if fuel(fuelfuel) <= 0 {
				break
			}
			fuelfuel = fuel(fuelfuel)
		}
	}

	return totalfuel
}

func test1() {
	// Test input
	fmt.Println("\nTest Day01#1")
	fmt.Println("mass 12: ", fuel1(12))
	fmt.Println("mass 14: ", fuel1(14))
	fmt.Println("mass 1969: ", fuel1(1969))
	fmt.Println("mass 100756: ", fuel1(100756))
}

func test2() {
	// Test input
	fmt.Println("\nTest Day01#2")
	fmt.Println("mass 12: ", fuel2(12))
	fmt.Println("mass 14: ", fuel2(14))
	fmt.Println("mass 1969: ", fuel2(1969))
	fmt.Println("mass 100756: ", fuel2(100756))
}

func Run1(test bool) int {
	if test {
		test1()
		return 0

	} else {
		// Run real code
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
}

func Run2(test bool) int {
	if test {
		test2()
		return 0

	} else {
		// Run real code
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
}
