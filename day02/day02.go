package day02

import (
	"log"
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
	util.Debug(false)
	log.Println("prog: ", prog)
	pc := 0
	for pc < len(prog) {
		switch prog[pc] {
		case 1:
			log.Println("progIn", pc, ": ", prog)
			log.Println("execIn", pc, ": ", prog[pc], " ADD")
			log.Println("ADD pc", pc,
				" op1", prog[pc+1],
				" op2", prog[pc+2],
				" at pos", prog[pc+3])
			prog[prog[pc+3]] = prog[prog[pc+1]] + prog[prog[pc+2]]
			log.Println("progOut: ", prog)
			pc = pc + 4
		case 2:
			log.Println("progIn", pc, ": ", prog)
			log.Println("execIn", pc, ": ", prog[pc], " MULT")
			prog[prog[pc+3]] = prog[prog[pc+1]] * prog[prog[pc+2]]
			log.Println("progOut: ", prog)
			pc = pc + 4
		case 99:
			log.Println("execIn", pc, ": ", prog[pc], " END")
			return prog
		default:
			log.Println("execIn", pc, ": ", prog[pc], " NOP")
			pc++
		}
	}
	return
}

func test1() int {
	inputs := []string{
		"1,9,10,3,2,3,11,0,99,30,40,50",
		"1,0,0,0,99",
		"2,4,4,5,99,0",
		"1,1,1,4,99,5,6,0,99"}

	util.Debug(false)
	log.Println("len input: ", len(inputs))

	execedProg := []int{}

	for i := range inputs {
		util.Debug(true)
		log.Println("\ninput", i, ": ", inputs[i])
		execedProg = execute(splitStr(inputs[i]))
		util.Debug(true)
		log.Println(" exec", i, ": ", execedProg)
	}
	return execedProg[0]
}

func Run1(test bool) int {
	if test {
		return test1()
	} else {
		lines, err := util.ReadF("day02/input.txt")
		if err != nil {
			return -1
		}

		input := splitStr(lines[0])

		// Updated data according to problem
		input[1] = 12
		input[2] = 2

		util.Debug(false)
		log.Println("\n", input)
		log.Println("input:", input)
		execedProg := execute(input)
		util.Debug(false)
		log.Println(" exec:", execedProg)
		return execedProg[0]
	}
}
