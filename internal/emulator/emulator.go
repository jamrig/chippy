package emulator

import (
	"os"
	"time"
)

type Emulator struct {
	Config           *Config
	Memory           *Memory
	Stack            *Stack
	DelayTimer       *Timer
	InstructionTimer *Timer
	// PC is the program counter register.
	PC uint16
	// I is the index register.
	I uint16
	// V are the general purpose registers.
	V [16]uint8
	// CPU        *CPU
	// Display    *Display
	// Window     *Window
}

func New(fontFile, programFile string) (*Emulator, error) {
	font, err := LoadFile(fontFile)
	if err != nil {
		return nil, err
	}

	program, err := LoadFile(programFile)
	if err != nil {
		return nil, err
	}

	config := CHIP8Config

	// w := NewWindow()
	// d := NewDisplay(DisplayWidth, DisplayHeight, DisplayFrequency, w)
	e := &Emulator{
		Config:           config,
		Memory:           NewMemory(config.MemorySize),
		Stack:            NewStack(config.StackInitialSize),
		DelayTimer:       NewTimer(config.DelayTimerFrequency),
		InstructionTimer: NewTimer(config.InstructionTimerFrequency),
		PC:               config.MemoryProgramAddress,
		I:                0,
		V:                [16]byte{},
		// Display:    d,
		// Window:     w,
	}

	e.Memory.Write(config.MemoryFontAddress, font)
	e.Memory.Write(config.MemoryProgramAddress, program)

	return e, nil
}

func (e *Emulator) Start() {
	// e.Window.Init()
	// defer e.Window.Destroy()

	now := time.Now().UnixNano()
	delta := int64(0)

	for { // !e.Window.ShouldExit() {
		delta = time.Now().UnixNano() - now

		e.DelayTimer.Tick(delta)

		if e.InstructionTimer.Tick(delta) {
			e.ExecuteNextInstruction()
		}
		// e.Display.Tick(delta)

		now = time.Now().UnixNano()
		// TODO: use the delta to reduce the sleep
		time.Sleep(1 * time.Microsecond)
	}
}

func LoadFile(file string) ([]byte, error) {
	// TODO: wrap error

	return os.ReadFile(file)
}

func (e *Emulator) ExecuteNextInstruction() {
	// TODO: reset PC if at end
	e.ExecuteOpcode(e.FetchOpcode())
}

// Fetch the next instruction and increment the PC.
func (e *Emulator) FetchOpcode() uint16 {
	opcode := uint16(0)

	opcode += uint16(e.Memory.Read(e.PC)) << 0x08
	opcode += uint16(e.Memory.Read(e.PC + 1))
	e.PC += 2

	return opcode
}

// Process the opcode.
func (e *Emulator) ExecuteOpcode(opcode uint16) {
	F := (opcode & 0xF000) >> 0x0C
	X := byte((opcode & 0x0F00) >> 0x08)
	Y := byte((opcode & 0x00F0) >> 0x04)
	N := byte(opcode & 0x000F)
	NN := byte(opcode & 0x00FF)
	NNN := uint16(opcode & 0x0FFF)

	switch {
	case opcode == 0x00E0:
		e.Op00E0()
	case opcode == 0x00EE:
		e.Op00EE()
	case F == 1:
		e.Op1NNN(NNN)
	case F == 2:
		e.Op2NNN(NNN)
	case F == 3:
		e.Op3XNN(X, NN)
	case F == 4:
		e.Op4XNN(X, NN)
	case F == 5:
		e.Op5XY0(X, Y)
	case F == 6:
		e.Op6XNN(X, NN)
	case F == 7:
		e.Op7XNN(X, NN)
	case F == 8 && N == 0:
		e.Op8XY0(X, Y)
	case F == 8 && N == 1:
		e.Op8XY1(X, Y)
	case F == 8 && N == 2:
		e.Op8XY2(X, Y)
	case F == 8 && N == 3:
		e.Op8XY3(X, Y)
	case F == 8 && N == 4:
		e.Op8XY4(X, Y)
	case F == 8 && N == 5:
		e.Op8XY5(X, Y)
	case F == 8 && N == 6:
		e.Op8XY6(X, Y)
	case F == 8 && N == 7:
		e.Op8XY7(X, Y)
	case F == 8 && N == 0xE:
		e.Op8XYE(X, Y)
	case F == 9:
		e.Op9XY0(X, Y)
	case F == 0xA:
		e.OpANNN(NNN)
	case F == 0xB:
		e.OpBNNN(X, NNN)
	case F == 0xC:
		e.OpCXNN(X, NN)
	case F == 0xD:
		e.OpDXYN(X, Y, N)
	case F == 0xE && NN == 0x9E:
		e.OpEX9E(X)
	case F == 0xE && NN == 0xA1:
		e.OpEXA1(X)
	case F == 0xF && NN == 0x07:
		e.OpFX07(X)
	case F == 0xF && NN == 0x0A:
		e.OpFX0A(X)
	case F == 0xF && NN == 0x15:
		e.OpFX15(X)
	case F == 0xF && NN == 0x18:
		e.OpFX18(X)
	case F == 0xF && NN == 0x1E:
		e.OpFX1E(X)
	case F == 0xF && NN == 0x29:
		e.OpFX29(X)
	case F == 0xF && NN == 33:
		e.OpFX33(X)
	case F == 0xF && NN == 55:
		e.OpFX55(X)
	case F == 0xF && NN == 65:
		e.OpFX65(X)
	}
}
