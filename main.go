package main

import "fmt"

import (
	"github.com/jonasferry/advent2019/day01"
	"github.com/jonasferry/advent2019/day02"
)

func main() {
	day := 2
	test := true

	switch day {
	case 1:
		fmt.Println("Day01#1: ", day01.Run1(test))
		fmt.Println("Day01#2: ", day01.Run2(test))
		break
	case 2:
		fmt.Println("Day02#1: ", day02.Run1(test))
//		fmt.Println("Day02#2: ", day02.Run2(test))
		break
	}
}