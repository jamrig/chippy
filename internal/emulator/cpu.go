package emulator

type CPU struct {
	Config           *CPUConfig
	Memory           *Memory
	Display          *Display
	Stack            *Stack
	DelayTimer       *Timer
	InstructionTimer *Timer
	// PC is the program counter.
	PC uint16
	// I is the index register.
	I uint16
	// V are the general purpose registers.
	V [16]uint8
}

func NewCPU(config *CPUConfig, programAddress uint16, memory *Memory, display *Display) *CPU {
	c := &CPU{
		Config:           config,
		Memory:           memory,
		Display:          display,
		Stack:            NewStack(config.StackInitialSize),
		DelayTimer:       NewTimer(config.DelayTimerFrequency),
		InstructionTimer: NewTimer(config.InstructionTimerFrequency),
		PC:               programAddress,
		I:                0,
		V:                [16]byte{},
	}

	return c
}

// Tick will tick the contained timers and if ready will perform a full CPU cycle.
func (c *CPU) Tick(delta int64) {
	c.DelayTimer.Tick(delta)

	if c.InstructionTimer.Tick(delta) {
		opcode := c.Fetch()
		instr := opcode.Decode()
		if instr == nil {
			// TODO: handle nil
		} else {
			instr.Execute(c, opcode)
		}
	}
}

// Fetch gets the next opcode and updates the PC.
func (c *CPU) Fetch() *Opcode {
	// TODO: reset PC if at end

	rawOpcode := uint16(0)
	rawOpcode += uint16(c.Memory.Read(c.PC)) << 0x08
	rawOpcode += uint16(c.Memory.Read(c.PC + 1))
	c.PC += 2

	return NewOpcode(rawOpcode)
}
