package main

import "fmt"

import (
	"github.com/jonasferry/advent2019/day01"
)

func main() {
	day := 1
	test := false

	switch day {
	case 1:
		fmt.Println("Day01#1: ", day01.Run1(test))
		fmt.Println("Day01#2: ", day01.Run2(test))
		break
	}
}