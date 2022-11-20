package emulator

import "os"

type Emulator struct {
	Memory *Memory
	Stack  *Stack
}

func New(fontFile, programFile string) (*Emulator, error) {
	e := &Emulator{
		Memory: NewMemory(MemorySize),
		Stack:  NewStack(StackInitialSize),
	}

	font, err := LoadFile(fontFile)
	if err != nil {
		return nil, err
	}

	program, err := LoadFile(programFile)
	if err != nil {
		return nil, err
	}

	e.Memory.Write(MemoryFontAddress, font)
	e.Memory.Write(MemoryProgramAddress, program)

	return e, nil
}

func LoadFile(file string) ([]byte, error) {
	// TODO: wrap error

	return os.ReadFile(file)
}
