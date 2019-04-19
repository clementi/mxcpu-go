package cpu

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
