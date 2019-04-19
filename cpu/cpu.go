package cpu

import "fmt"

// State holds the state of the machine
type State struct {
	Cycles         byte
	Counter        byte
	ProgramCounter byte
	Accumulator    byte
	Registers      []byte
}

// InitialState creates an empty State
func InitialState() *State {
	return &State{
		Cycles:         0,
		Counter:        0,
		ProgramCounter: 0,
		Accumulator:    0,
		Registers:      make([]byte, 16),
	}
}

// Run executes a program and updates state accordingly
func Run(program []byte, state *State) error {
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
	return Run(program, state)
}
