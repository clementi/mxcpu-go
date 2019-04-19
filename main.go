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
	if err = cpu.Run(program, state); err != nil {
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
