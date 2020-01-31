package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(fmt.Sprintf("could not read file; got %v", err))
	}
	stack := bytes.Split(b, []byte(","))
	program := make([]int, 0, len(stack))

	for _, c := range stack {
		integer, err := strconv.Atoi(string(c))
		if err != nil {
			panic(fmt.Sprintf("could not convert %s to int; got %v", c, err))
		}
		program = append(program, integer)
	}

	var noun, verb int
	success := false
loop:
	for noun = 0; noun < len(program); noun++ {
		for verb = 0; verb < len(program); verb++ {
			copiedProgram := make([]int, len(program))
			copy(copiedProgram, program)
			copiedProgram[1] = noun
			copiedProgram[2] = verb
			result, err := runProgram(copiedProgram)
			if err != nil {
				panic(err)
			}
			if result == 19690720 {
				success = true
				break loop
			}
		}
	}
	if success {
		result := (100 * noun) + verb
		fmt.Println(result)
	}
}

const (
	OpAddition       = 1
	OpMultiplication = 2
	OpHalt           = 99
)

func runProgram(program []int) (int, error) {
	for i := 0; i < len(program); i += 4 {
		opcode := program[i]
		switch opcode {
		case OpAddition:
			program[program[i+3]] = program[program[i+1]] + program[program[i+2]]
		case OpMultiplication:
			program[program[i+3]] = program[program[i+1]] * program[program[i+2]]
		case OpHalt:
			return program[0], nil
		default:
			return 0, errors.New(fmt.Sprintf("ERR %d", opcode))
		}
	}
	return 0, errors.New("did not exit properly")
}
