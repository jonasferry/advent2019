package day02

import "github.com/jonasferry/advent2019/util"

import(
	"log"
	"strconv"
	"strings"
)

func splitStr(in string) (out []int) {
	progStr := strings.Split(in, ",")
	prog := make([]int, len(progStr))
	for i := range prog {
		prog[i], _ = strconv.Atoi(progStr[i])
	}
	return prog
}

func parseProg(input string) (instr [][]int) {
	util.Debug(false)
	prog := splitStr(input)
	log.Println("prog: ", prog)
	pc := 0
	instr = [][]int{}
	log.Println("len(prog): ", len(prog))
	i := 0
	for pc < len(prog) {
		log.Println("pc: ", pc)
		switch prog[pc] {
		case 1:
			log.Println("parseProg: ", prog[pc:pc+4], " ADD")
			instr = append(instr, prog[pc:pc+4])
			pc = pc+4
			break
		case 2:
			log.Println("parseProg: ", prog[pc:pc+4], " MULT")
			instr = append(instr, prog[pc:pc+4])
			pc = pc+4
			break
		case 99:
			log.Println("parseProg: ", prog[pc], " END")
			instr = append(instr, prog[pc:pc+1])
			pc++
			break
		default:
			log.Println("parseProg: ", prog[pc], " NOP")
			pc++
			break
		}
		i++	
	}
	return instr
}

func execute(input string, instrIn [][]int) (output []int) {
	util.Debug(false)
	prog := splitStr(input)
	output = prog
	log.Println("prog: ", prog)
	log.Println("output: ", output, "\n")
	for i := 0; i < len(instrIn); i++ {
		instr := instrIn[i]
		switch instr[0] {
		case 1:
			log.Println("progIn: ", output)
			log.Println("execIn: ", instr, " ADD")
			op1 := output[instr[1]]
			op2 := output[instr[2]]
			output[instr[3]] = op1 + op2
			log.Println("progOut: ", output)
			break
		case 2:
			log.Println("progIn: ", output)
			log.Println("execIn: ", instr, " MULT")
			op1 := output[instr[1]]
			op2 := output[instr[2]]
			output[instr[3]] = op1 * op2
			log.Println("progOut: ", output)
			break
		case 99:
			log.Println("execIn: ", instr, " END")
			return output
		default:
			log.Println("execIn: ", instr, " NOP")
			break
		}
	}
	return
}

func test1() {
	inputs := []string{
		"1,9,10,3,2,3,11,0,99,30,40,50",
		"1,0,0,0,99",
		"2,4,4,5,99,0",
		"1,1,1,4,99,5,6,0,99"}

		util.Debug(true)
		log.Println("len input: ", len(inputs))

	for i := range inputs {
		util.Debug(true)
		input := inputs[i]
		log.Println("input", i, ": ", splitStr(input))
		parsedProg := parseProg(input)
		execedProg := execute(input, parsedProg)
		util.Debug(true)
		log.Println(" exec: ", execedProg)
	}
}

func Run1(test bool) int {
	if test {
		test1()
		return 0
	} else {
		return 0
	}
}
