package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/clementi/mxcpu/cpu"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("program file required")
	}

	// Load program
	program, err := loadProgram(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// Run
	state := cpu.InitialState()
	if err = run(program, state); err != nil {
		log.Fatal(err)
	}

	// Print out final state
	fmt.Printf("Cycles    : %d\n", state.Cycles)
	fmt.Printf("INC       : %#x\n", state.Counter)
	fmt.Printf("PC        : %#x\n", state.ProgramCounter)
	fmt.Printf("ACC       : %#x\n", state.Accumulator)
	fmt.Printf("Registers : %v\n", state.Registers)
}

func loadProgram(path string) ([]byte, error) {
	buffer, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	programContents := string(buffer)
	programContents = strings.Trim(programContents, " \r\n\t")
	programParts := strings.Split(programContents, " ")

	program, err := toBytes(programParts)
	if err != nil {
		return nil, err
	}
	return program, nil
}

func toBytes(strings []string) ([]byte, error) {
	bytes := make([]byte, len(strings))

	for i := 0; i < len(strings); i++ {
		parsed, err := strconv.ParseUint(strings[i], 16, 8)
		if err != nil {
			return nil, err
		}
		bytes[i] = byte(parsed)
	}

	return bytes, nil
}

func run(program []byte, state *cpu.State) error {
	state.Cycles++

	switch opCode := program[state.ProgramCounter]; opCode {
	case 0x00:
		return nil
	case 0xB1:
		value := program[state.ProgramCounter+1]
		state.ProgramCounter = value
	case 0xB2:
		index := program[state.ProgramCounter+1]
		value := program[state.ProgramCounter+2]
		if state.Accumulator == state.Registers[index] {
			state.ProgramCounter = value
		} else {
			state.ProgramCounter += 3
		}
	case 0xB3:
		value := program[state.ProgramCounter+1]
		pcValue := program[state.ProgramCounter+2]
		if state.Accumulator == value {
			state.ProgramCounter = pcValue
		} else {
			state.ProgramCounter += 3
		}
	case 0xC0:
		index := program[state.ProgramCounter+1]
		state.Accumulator += state.Registers[index]
		state.ProgramCounter += 2
	case 0xC1:
		value := program[state.ProgramCounter+1]
		state.Accumulator += value
		state.ProgramCounter += 2
	case 0xC2:
		state.Counter++
		state.ProgramCounter++
	case 0xC3:
		state.Counter--
		state.ProgramCounter++
	case 0xC4:
		state.Counter = 0
		state.ProgramCounter++
	case 0xC5:
		state.Accumulator = state.Counter
		state.ProgramCounter++
	case 0xC6:
		state.Counter = state.Accumulator
		state.ProgramCounter++
	case 0xD0:
		index := program[state.ProgramCounter+1]
		value := state.Registers[index]
		state.Accumulator = value
		state.ProgramCounter += 2
	case 0xD1:
		value := program[state.ProgramCounter+1]
		state.Accumulator = value
		state.ProgramCounter += 2
	case 0xD2:
		index := program[state.ProgramCounter+1]
		state.Registers[index] = state.Accumulator
		state.ProgramCounter += 2
	default:
		return fmt.Errorf("unknown instruction %d", opCode)
	}
	return run(program, state)
}
