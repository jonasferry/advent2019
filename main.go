package main

import "fmt"

import (
	"github.com/jonasferry/advent2019/day01"
	"github.com/jonasferry/advent2019/day02"
)

func main() {
	day := 0
	test := false
	answer := 0

	if day == 1 || day == 0 {
		answer = day01.Run1(test)
		if answer == 3318604 {
			fmt.Println("Day01#1: ", answer, "correct")
		} else {
			fmt.Println("Day01#1: ", answer, "wrong")
		}
		answer = day01.Run2(test)
		if answer == 4975039 {
			fmt.Println("Day01#2: ", answer, "correct")
		} else {
			fmt.Println("Day01#1: ", answer, "wrong")
		}
	} 

	if day == 2 || day == 0 { 
		answer = day02.Run1(test)
		if answer == 4330636 {
			fmt.Println("Day02#1: ", answer, "correct")
		} else {
			fmt.Println("Day02#1: ", answer, "wrong")
		}
//		fmt.Println("Day02#2: ", day02.Run2(test))
	}
}