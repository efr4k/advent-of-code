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
	program[1] = 12
	program[2] = 2
	result, err := runProgram(program)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

const (
	OpcodeAdd = 1
	OpcodeMul = 2
	OpcodeDie = 99
)

func runProgram(program []int) (int, error) {
	for i := 0; i < len(program); i += 4 {
		opcode := program[i]
		switch opcode {
		case OpcodeAdd:
			//fmt.Print("ADD", program[i:i+4])
			program[program[i+3]] = program[program[i+1]] + program[program[i+2]]
			//fmt.Println("\t", program[i:i+4])
		case OpcodeMul:
			//fmt.Print("MUL", program[i:i+4])
			program[program[i+3]] = program[program[i+1]] * program[program[i+2]]
			//fmt.Println("\t", program[i:i+4])
		case OpcodeDie:
			//fmt.Println("DIE")
			return program[0], nil
		default:
			return 0, errors.New(fmt.Sprintf("ERR %d", opcode))
		}
	}
	return 0, errors.New("did not exit properly")
}
