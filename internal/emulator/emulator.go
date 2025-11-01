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

// ExecuteNextInstruction will fetch and execute the next instruction, updating the PC when appropriate.
func (e *Emulator) ExecuteNextInstruction() {
	// TODO: reset PC if at end

	rawOpcode := uint16(0)
	rawOpcode += uint16(e.Memory.Read(e.PC)) << 0x08
	rawOpcode += uint16(e.Memory.Read(e.PC + 1))
	e.PC += 2

	opcode := NewOpcode(rawOpcode)

	for _, instr := range Instructions {
		if instr.Is(opcode) {
			instr.Execute(e, opcode)
			return
		}
	}
}
