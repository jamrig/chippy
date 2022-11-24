package emulator

import (
	"fmt"
)

// CPU represents the CPU.
type CPU struct {
	Memory         *Memory
	Display        *Display
	LastTimerValue int
	Timer          *Timer
	V              [15]uint8
	VF             byte
	PC             uint16
	I              uint16
}

// NewCPU returns a new CPU.
func NewCPU(memory *Memory, display *Display, freq int, pc uint16) *CPU {
	return &CPU{
		Memory:         memory,
		Display:        display,
		LastTimerValue: 0,
		Timer:          NewTimer(freq),
		V:              [15]byte{},
		VF:             0,
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
		c.OpScreenClear()
	case opcode == 0x00EE:
		c.OpSubroutineReturn()
	case F == 1:
		c.OpJump(NNN)
	case F == 2:
		c.OpSubroutineCall(NNN)
	case F == 3:
		c.OpRegisterSkipIfEqual()
	case F == 4:
		c.OpRegisterSkipIfNotEqual()
	case F == 5:
		c.OpRegistersSkipIfEqual()
	case F == 6:
		c.OpRegisterSet(X, NN)
	case F == 7:
		c.OpRegisterAdd(X, NN)
	case F == 8 && N == 0:
		c.OpRegistersSet(X, Y)
	case F == 8 && N == 1:
		c.OpRegistersBitwiseOR(X, Y)
	case F == 8 && N == 2:
		c.OpRegistersBitwiseAND(X, Y)
	case F == 8 && N == 3:
		c.OpRegistersBitwiseXOR(X, Y)
	case F == 8 && N == 4:
		c.OpRegistersAdd(X, Y)
	case F == 8 && N == 5:
		fmt.Println("Set VX to VX + VY (if > 255 set VF to 1 else 0)")
	case F == 8 && N == 7:
		fmt.Println("Set VX to VY - VX (if > 0 set VF to 1 else 0)")
	case F == 8 && N == 6:
		fmt.Println("Shift (ambiguous)")
	case F == 8 && N == 0xE:
		fmt.Println("Shift (ambiguous)")
	case F == 9:
		c.OpRegistersSkipIfNotEqual()
	case F == 0xA:
		c.OpRegisterSetIndex(NNN)
	case F == 0xB:
		fmt.Println("Jump with offset (ambiguous)")
	case F == 0xC:
		fmt.Println("Generate a random number, AND with NN, store in VX")
	case F == 0xD:
		c.OpScreenDraw(X, Y, N)
	case F == 0xE && NN == 0x9E:
		fmt.Println("Skip if key VX is down")
	case F == 0xE && NN == 0xA1:
		fmt.Println("Skip if key VX is up")
	case F == 0xF && NN == 0x07:
		fmt.Println("Set VX to delay timer value")
	case F == 0xF && NN == 0x15:
		fmt.Println("Set delay timer value to VX")
	case F == 0xF && NN == 0x18:
		fmt.Println("Set sound timer value to VX")
	case F == 0xF && NN == 0x1E:
		fmt.Println("Add VX to I and store in I (if > 0xFFF set VF to 1 else 0")
	case F == 0xF && NN == 0x0A:
		fmt.Println("Block until key, store in VX")
	case F == 0xF && NN == 0x29:
		fmt.Println("Font Character")
	case F == 0xF && NN == 33:
		fmt.Println("Binary-coded decimal conversion")
	case F == 0xF && NN == 55:
		fmt.Println("Store (ambiguous)")
	case F == 0xF && NN == 65:
		fmt.Println("Load (ambiguous)")
	}
}

// JUMP

func (c *CPU) OpJump(addr uint16) {
	c.PC = addr
}

// REGISTER

func (c *CPU) OpRegisterAdd(X byte, NN byte) {
	c.V[X] = c.V[X] + NN
}

func (c *CPU) OpRegisterSet(X byte, NN byte) {
	c.V[X] = NN
}

func (c *CPU) OpRegisterSetIndex(NNN uint16) {
	c.I = NNN
}

func (c *CPU) OpRegisterSkipIfEqual() {
	fmt.Println("OpRegisterSkipIfEqual")
}

func (c *CPU) OpRegisterSkipIfNotEqual() {
	fmt.Println("OpRegisterSkipIfNotEqual")
}

// REGISTERS

func (c *CPU) OpRegistersAdd(X byte, Y byte) {
	if int(c.V[X])+int(c.V[Y]) > 0xFFFF {
		c.VF = 1
	} else {
		c.VF = 0
	}

	c.V[X] = c.V[X] + c.V[Y]
}

func (c *CPU) OpRegistersBitwiseAND(X byte, Y byte) {
	c.V[X] = c.V[X] & c.V[Y]
}

func (c *CPU) OpRegistersBitwiseOR(X byte, Y byte) {
	c.V[X] = c.V[X] | c.V[Y]
}

func (c *CPU) OpRegistersBitwiseXOR(X byte, Y byte) {
	c.V[X] = c.V[X] ^ c.V[Y]
}

func (c *CPU) OpRegistersSet(X byte, Y byte) {
	c.V[X] = c.V[Y]
}

func (c *CPU) OpRegistersSkipIfEqual() {
	fmt.Println("OpRegistersSkipIfEqual")
}

func (c *CPU) OpRegistersSkipIfNotEqual() {
	fmt.Println("OpRegistersSkipIfNotEqual")
}

// SCREEN

func (c *CPU) OpScreenClear() {
	c.Display.Clear()
}

func (c *CPU) OpScreenDraw(X byte, Y byte, N byte) {
	x := int(c.V[X])
	y := int(c.V[Y])
	n := int(N)
	data := make([]byte, 0, n)

	for i := 0; i < n; i++ {
		data = append(data, c.Memory.Read(c.I+uint16(i)))
	}

	unset := c.Display.Write(x, y, data)

	if unset {
		c.VF = 1
	} else {
		c.VF = 0
	}
}

// SUBROUTINE

func (c *CPU) OpSubroutineCall(addr uint16) {
	fmt.Println("OpSubroutineCall")
}

func (c *CPU) OpSubroutineReturn() {
	fmt.Println("OpSubroutineReturn")
}
