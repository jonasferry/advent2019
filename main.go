package main

import (
	"fmt"

	"github.com/jonasferry/advent2019/day01"
	"github.com/jonasferry/advent2019/day02"
)

func main() {
	day := 0

	if day == 1 || day == 0 {
		fmt.Println("Day01a:", day01.Run1())
		fmt.Println("Day01b:", day01.Run2())
	}

	if day == 2 || day == 0 {
		fmt.Println("Day02a:", day02.Run1())
	}
}
