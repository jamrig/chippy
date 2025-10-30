package emulator

import "math/rand"

// CPU represents the CPU.
type CPU struct {
	Memory         *Memory
	Stack          *Stack
	Display        *Display
	DelayTimer     *Timer
	LastTimerValue int
	Timer          *Timer
	Legacy         bool
	V              [16]uint8
	PC             uint16
	I              uint16
}

// NewCPU returns a new CPU.
func NewCPU(memory *Memory, stack *Stack, display *Display, delayTimer *Timer, freq int, pc uint16, legacy bool) *CPU {
	return &CPU{
		Memory:         memory,
		Stack:          stack,
		Display:        display,
		DelayTimer:     delayTimer,
		LastTimerValue: 0,
		Timer:          NewTimer(freq),
		Legacy:         legacy,
		V:              [16]byte{},
		PC:             pc,
		I:              0,
	}
}

// Tick calls the timer tick and if it has changed performs a CPU cycle.
func (c *CPU) Tick(delta int64) {
	c.Timer.Tick(delta)

	if c.LastTimerValue != c.Timer.GetValue() {
		c.LastTimerValue = c.Timer.GetValue()

		// reset PC if at end

		opcode := c.Fetch()
		c.Process(opcode)
	}
}

// Fetch the next instruction and increment the PC.
func (c *CPU) Fetch() uint16 {
	opcode := uint16(0)

	opcode += uint16(c.Memory.Read(c.PC)) << 0x08
	opcode += uint16(c.Memory.Read(c.PC + 1))
	c.PC += 2

	return opcode
}

// Process the opcode.
func (c *CPU) Process(opcode uint16) {
	F := (opcode & 0xF000) >> 0x0C
	X := byte((opcode & 0x0F00) >> 0x08)
	Y := byte((opcode & 0x00F0) >> 0x04)
	N := byte(opcode & 0x000F)
	NN := byte(opcode & 0x00FF)
	NNN := uint16(opcode & 0x0FFF)

	switch {
	case opcode == 0x00E0:
		c.Op00E0()
	case opcode == 0x00EE:
		c.Op00EE()
	case F == 1:
		c.Op1NNN(NNN)
	case F == 2:
		c.Op2NNN(NNN)
	case F == 3:
		c.Op3XNN(X, NN)
	case F == 4:
		c.Op4XNN(X, NN)
	case F == 5:
		c.Op5XY0(X, Y)
	case F == 6:
		c.Op6XNN(X, NN)
	case F == 7:
		c.Op7XNN(X, NN)
	case F == 8 && N == 0:
		c.Op8XY0(X, Y)
	case F == 8 && N == 1:
		c.Op8XY1(X, Y)
	case F == 8 && N == 2:
		c.Op8XY2(X, Y)
	case F == 8 && N == 3:
		c.Op8XY3(X, Y)
	case F == 8 && N == 4:
		c.Op8XY4(X, Y)
	case F == 8 && N == 5:
		c.Op8XY5(X, Y)
	case F == 8 && N == 6:
		c.Op8XY6(X, Y)
	case F == 8 && N == 7:
		c.Op8XY7(X, Y)
	case F == 8 && N == 0xE:
		c.Op8XYE(X, Y)
	case F == 9:
		c.Op9XY0(X, Y)
	case F == 0xA:
		c.OpANNN(NNN)
	case F == 0xB:
		c.OpBNNN(X, NNN)
	case F == 0xC:
		c.OpCXNN(X, NN)
	case F == 0xD:
		c.OpDXYN(X, Y, N)
	case F == 0xE && NN == 0x9E:
		c.OpEX9E(X)
	case F == 0xE && NN == 0xA1:
		c.OpEXA1(X)
	case F == 0xF && NN == 0x07:
		c.OpFX07(X)
	case F == 0xF && NN == 0x0A:
		c.OpFX0A(X)
	case F == 0xF && NN == 0x15:
		c.OpFX15(X)
	case F == 0xF && NN == 0x18:
		c.OpFX18(X)
	case F == 0xF && NN == 0x1E:
		c.OpFX1E(X)
	case F == 0xF && NN == 0x29:
		c.OpFX29(X)
	case F == 0xF && NN == 33:
		c.OpFX33(X)
	case F == 0xF && NN == 55:
		c.OpFX55(X)
	case F == 0xF && NN == 65:
		c.OpFX65(X)
	}
}

func (c *CPU) Op00E0() {
	c.Display.Clear()
}

func (c *CPU) Op00EE() {
	c.PC = c.Stack.Pop()
}

func (c *CPU) Op1NNN(NNN uint16) {
	c.PC = NNN
}

func (c *CPU) Op2NNN(NNN uint16) {
	c.Stack.Push(c.PC)
	c.PC = NNN
}

func (c *CPU) Op3XNN(X byte, NN byte) {
	if c.V[X] == NN {
		c.PC += 2
	}
}

func (c *CPU) Op4XNN(X byte, NN byte) {
	if c.V[X] != NN {
		c.PC += 2
	}
}

func (c *CPU) Op5XY0(X byte, Y byte) {
	if c.V[X] == c.V[Y] {
		c.PC += 2
	}
}

func (c *CPU) Op6XNN(X byte, NN byte) {
	c.V[X] = NN
}

func (c *CPU) Op7XNN(X byte, NN byte) {
	c.V[X] = c.V[X] + NN
}

func (c *CPU) Op8XY0(X byte, Y byte) {
	c.V[X] = c.V[Y]
}

func (c *CPU) Op8XY1(X byte, Y byte) {
	c.V[X] = c.V[X] | c.V[Y]
}

func (c *CPU) Op8XY2(X byte, Y byte) {
	c.V[X] = c.V[X] & c.V[Y]
}

func (c *CPU) Op8XY3(X byte, Y byte) {
	c.V[X] = c.V[X] ^ c.V[Y]
}

func (c *CPU) Op8XY4(X byte, Y byte) {
	if int(c.V[X])+int(c.V[Y]) > 0xFFFF {
		c.V[15] = 1
	} else {
		c.V[15] = 0
	}

	c.V[X] = c.V[X] + c.V[Y]
}

func (c *CPU) Op8XY5(X byte, Y byte) {
	if int(c.V[X])-int(c.V[Y]) > 0x0 {
		c.V[15] = 1
	} else {
		c.V[15] = 0
	}

	c.V[X] = c.V[X] - c.V[Y]
}

func (c *CPU) Op8XY6(X byte, Y byte) {
	VX := c.V[X]

	if c.Legacy {
		VX = c.V[Y]
	}

	c.V[X] = VX >> 1
	c.V[15] = GetBitAtPosition(VX, 0)
}

func (c *CPU) Op8XY7(X byte, Y byte) {
	if int(c.V[Y])-int(c.V[X]) > 0x0 {
		c.V[15] = 1
	} else {
		c.V[15] = 0
	}

	c.V[X] = c.V[Y] - c.V[X]
}

func (c *CPU) Op8XYE(X byte, Y byte) {
	VX := c.V[X]

	if c.Legacy {
		VX = c.V[Y]
	}

	c.V[X] = VX << 1
	c.V[15] = GetBitAtPosition(VX, 7)
}

func (c *CPU) Op9XY0(X byte, Y byte) {
	if c.V[X] != c.V[Y] {
		c.PC += 2
	}
}

func (c *CPU) OpANNN(NNN uint16) {
	c.I = NNN
}

func (c *CPU) OpBNNN(X byte, NNN uint16) {
	if c.Legacy {
		c.PC = uint16(c.V[0]) + NNN
	} else {
		c.PC = uint16(c.V[X]) + NNN
	}
}

func (c *CPU) OpCXNN(X byte, NN byte) {
	c.V[X] = byte(rand.Intn(256)) & NN
}

func (c *CPU) OpDXYN(X byte, Y byte, N byte) {
	x := int(c.V[X])
	y := int(c.V[Y])
	n := int(N)
	data := make([]byte, 0, n)

	for i := 0; i < n; i++ {
		data = append(data, c.Memory.Read(c.I+uint16(i)))
	}

	unset := c.Display.Write(x, y, data)

	if unset {
		c.V[15] = 1
	} else {
		c.V[15] = 0
	}
}

func (c *CPU) OpEX9E(X byte) {
	// key
}

func (c *CPU) OpEXA1(X byte) {
	// key
}

func (c *CPU) OpFX07(X byte) {
	c.V[X] = byte(c.DelayTimer.GetValue())
}

func (c *CPU) OpFX0A(X byte) {
	// key
}

func (c *CPU) OpFX15(X byte) {
	c.DelayTimer.SetValue(int(c.V[X]))
}

func (c *CPU) OpFX18(X byte) {
	// No sound implemented.
}

func (c *CPU) OpFX1E(X byte) {
	if int(c.I)+int(c.V[X]) > 0x0FFF {
		c.V[15] = 1
	}

	c.I = c.I + uint16(c.V[X])
}

func (c *CPU) OpFX29(X byte) {
	// font
}

func (c *CPU) OpFX33(X byte) {
	VX := c.V[X]

	c.Memory.Write(c.I, []byte{
		byte(int(VX) / 100),
		byte((int(VX) % 100) / 10),
		byte(int(VX) % 10),
	})
}

func (c *CPU) OpFX55(X byte) {
	if c.Legacy {
		for i := 0; i <= int(X); i++ {
			c.Memory.Write(c.I, []byte{c.V[i]})
			c.I++
		}
	} else {
		b := make([]byte, 0, X+1)

		for i := 0; i <= int(X); i++ {
			b = append(b, c.V[i])
		}

		c.Memory.Write(c.I, b)
	}
}

func (c *CPU) OpFX65(X byte) {
	if c.Legacy {
		for i := 0; i <= int(X); i++ {
			c.V[i] = c.Memory.Read(c.I)
			c.I++
		}
	} else {
		for i := 0; i <= int(X); i++ {
			c.V[i] = c.Memory.Read(c.I + uint16(i))
		}
	}
}
