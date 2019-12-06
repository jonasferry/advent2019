package day02

import (
	"strconv"
	"strings"

	"github.com/jonasferry/advent2019/util"
)

func splitStr(in string) (out []int) {
	progStr := strings.Split(in, ",")
	prog := make([]int, len(progStr))
	for i := range prog {
		prog[i], _ = strconv.Atoi(progStr[i])
	}
	return prog
}

func execute(prog []int) (output []int) {
	pc := 0
	for pc < len(prog) {
		switch prog[pc] {
		case 1:
			prog[prog[pc+3]] = prog[prog[pc+1]] + prog[prog[pc+2]]
			pc = pc + 4
		case 2:
			prog[prog[pc+3]] = prog[prog[pc+1]] * prog[prog[pc+2]]
			pc = pc + 4
		case 99:
			return prog
		default:
			pc++
		}
	}
	return
}

func Run1() int {
	lines, err := util.ReadF("day02/input.txt")
	if err != nil {
		return -1
	}

	input := splitStr(lines[0])

	// Updated data according to problem
	input[1] = 12
	input[2] = 2

	execedProg := execute(input)
	return execedProg[0]

}
