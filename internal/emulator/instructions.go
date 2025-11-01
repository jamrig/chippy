package emulator

import "math/rand"

// Instruction represents a CPU instruction.
type Instruction struct {
	Name    string
	Is      func(o *Opcode) bool
	Execute func(c *CPU, o *Opcode)
}

// Instructions contains all of the available CPU instructions.
var Instructions = []Instruction{
	{
		Name: "[00E0] Clear Screen",
		Is:   func(o *Opcode) bool { return o.Raw == 0x00E0 },
		Execute: func(c *CPU, o *Opcode) {
			// c.Display.Clear()
		},
	},
	{
		Name: "[00EE] Return Subroutine",
		Is:   func(o *Opcode) bool { return o.Raw == 0x00EE },
		Execute: func(c *CPU, o *Opcode) {
			c.PC = c.Stack.Pop()
		},
	},
	{
		Name: "[1NNN] Jump",
		Is:   func(o *Opcode) bool { return o.F == 1 },
		Execute: func(c *CPU, o *Opcode) {
			c.PC = o.NNN
		},
	},
	{
		Name: "[2NNN] Call Subroutine",
		Is:   func(o *Opcode) bool { return o.F == 2 },
		Execute: func(c *CPU, o *Opcode) {
			c.Stack.Push(c.PC)
			c.PC = o.NNN
		},
	},
	{
		Name: "[3XNN] Skip If VX == NN",
		Is:   func(o *Opcode) bool { return o.F == 3 },
		Execute: func(c *CPU, o *Opcode) {
			if c.V[o.X] == o.NN {
				c.PC += 2
			}
		},
	},
	{
		Name: "[4XNN] Skip If VX != NN",
		Is:   func(o *Opcode) bool { return o.F == 4 },
		Execute: func(c *CPU, o *Opcode) {
			if c.V[o.X] != o.NN {
				c.PC += 2
			}
		},
	},
	{
		Name: "[5XY0] Skip If VX == VY",
		Is:   func(o *Opcode) bool { return o.F == 5 },
		Execute: func(c *CPU, o *Opcode) {
			if c.V[o.X] == c.V[o.Y] {
				c.PC += 2
			}
		},
	},
	{
		Name: "[6XNN] VX == NN",
		Is:   func(o *Opcode) bool { return o.F == 6 },
		Execute: func(c *CPU, o *Opcode) {
			c.V[o.X] = o.NN
		},
	},
	{
		Name: "[7XNN] Vx += NN (no carry)",
		Is:   func(o *Opcode) bool { return o.F == 7 },
		Execute: func(c *CPU, o *Opcode) {
			c.V[o.X] = c.V[o.X] + o.NN
		},
	},
	{
		Name: "[8XY0] Vx = Vy",
		Is:   func(o *Opcode) bool { return o.F == 8 && o.N == 0 },
		Execute: func(c *CPU, o *Opcode) {
			c.V[o.X] = c.V[o.Y]
		},
	},
	{
		Name: "[8XY1] Vx |= Vy",
		Is:   func(o *Opcode) bool { return o.F == 8 && o.N == 1 },
		Execute: func(c *CPU, o *Opcode) {
			c.V[o.X] |= c.V[o.Y]
		},
	},
	{
		Name: "[8XY2] Vx &= Vy",
		Is:   func(o *Opcode) bool { return o.F == 8 && o.N == 2 },
		Execute: func(c *CPU, o *Opcode) {
			c.V[o.X] &= c.V[o.Y]
		},
	},
	{
		Name: "[8XY3] Vx ^= Vy",
		Is:   func(o *Opcode) bool { return o.F == 8 && o.N == 3 },
		Execute: func(c *CPU, o *Opcode) {
			c.V[o.X] ^= c.V[o.Y]
		},
	},
	{
		Name: "[8XY4] Vx += Vy",
		Is:   func(o *Opcode) bool { return o.F == 8 && o.N == 4 },
		Execute: func(c *CPU, o *Opcode) {
			if int(c.V[o.X])+int(c.V[o.Y]) > 0xFFFF {
				c.V[15] = 1
			} else {
				c.V[15] = 0
			}

			c.V[o.X] += c.V[o.Y]
		},
	},
	{
		Name: "[8XY5] Vx -= Vy",
		Is:   func(o *Opcode) bool { return o.F == 8 && o.N == 5 },
		Execute: func(c *CPU, o *Opcode) {
			if int(c.V[o.X])-int(c.V[o.Y]) > 0x0 {
				c.V[15] = 1
			} else {
				c.V[15] = 0
			}

			c.V[o.X] -= c.V[o.Y]
		},
	},
	{
		Name: "[8XY6] Vx >>= 1",
		Is:   func(o *Opcode) bool { return o.F == 8 && o.N == 6 },
		Execute: func(c *CPU, o *Opcode) {
			VX := c.V[o.X]

			if c.Config.InstructionAssignBeforeShift {
				VX = c.V[o.Y]
			}

			c.V[o.X] = VX >> 1
			c.V[15] = GetBitAtPosition(VX, 0)
		},
	},
	{
		Name: "[8XY7] Vx = Vy - Vx",
		Is:   func(o *Opcode) bool { return o.F == 8 && o.N == 7 },
		Execute: func(c *CPU, o *Opcode) {
			if int(c.V[o.Y])-int(c.V[o.X]) > 0x0 {
				c.V[15] = 1
			} else {
				c.V[15] = 0
			}

			c.V[o.X] = c.V[o.Y] - c.V[o.X]
		},
	},
	{
		Name: "[8XYE] Vx <<= 1",
		Is:   func(o *Opcode) bool { return o.F == 8 && o.N == 0xE },
		Execute: func(c *CPU, o *Opcode) {
			VX := c.V[o.X]

			if c.Config.InstructionAssignBeforeShift {
				VX = c.V[o.Y]
			}

			c.V[o.X] = VX << 1
			c.V[15] = GetBitAtPosition(VX, 7)
		},
	},
	{
		Name: "[9XY0] Skip If Vx != Vy",
		Is:   func(o *Opcode) bool { return o.F == 9 },
		Execute: func(c *CPU, o *Opcode) {
			if c.V[o.X] != c.V[o.Y] {
				c.PC += 2
			}
		},
	},
	{
		Name: "[ANNN] Set Index",
		Is:   func(o *Opcode) bool { return o.F == 0xA },
		Execute: func(c *CPU, o *Opcode) {
			c.I = o.NNN
		},
	},
	{
		Name: "[BNNN] Jump With Offset",
		Is:   func(o *Opcode) bool { return o.F == 0xB },
		Execute: func(c *CPU, o *Opcode) {
			if c.Config.InstructionUseVxForOffset {
				c.PC = uint16(c.V[0]) + o.NNN
			} else {
				c.PC = uint16(c.V[o.X]) + o.NNN
			}
		},
	},
	{
		Name: "[CNNN] Rand",
		Is:   func(o *Opcode) bool { return o.F == 0xC },
		Execute: func(c *CPU, o *Opcode) {
			c.V[o.X] = byte(rand.Intn(256)) & o.NN
		},
	},
	{
		Name: "[DNNN] Display",
		Is:   func(o *Opcode) bool { return o.F == 0xC },
		Execute: func(c *CPU, o *Opcode) {
			c.V[o.X] = byte(rand.Intn(256)) & o.NN
		},
	},
	{
		Name: "[DNNN] Display",
		Is:   func(o *Opcode) bool { return o.F == 0xD },
		Execute: func(c *CPU, o *Opcode) {
			// x := int(c.V[o.X])
			// y := int(c.V[o.Y])
			n := int(o.N)
			data := make([]byte, 0, n)

			for i := 0; i < n; i++ {
				data = append(data, c.Memory.Read(c.I+uint16(i)))
			}

			// unset := c.Display.Write(x, y, data)

			// if unset {
			// 	c.V[15] = 1
			// } else {
			// 	c.V[15] = 0
			// }
		},
	},
	{
		Name: "[EX9E] Display",
		Is:   func(o *Opcode) bool { return o.F == 0xE && o.NN == 0x9E },
		Execute: func(c *CPU, o *Opcode) {
			// TODO: key
		},
	},
	{
		Name: "[EX9E] Skip If Key Pressed",
		Is:   func(o *Opcode) bool { return o.F == 0xE && o.NN == 0x9E },
		Execute: func(c *CPU, o *Opcode) {
			// TODO: key
		},
	},
	{
		Name: "[EXA1] Skip If Key Not Pressed",
		Is:   func(o *Opcode) bool { return o.F == 0xE && o.NN == 0xA1 },
		Execute: func(c *CPU, o *Opcode) {
			// TODO: key
		},
	},
	{
		Name: "[FX07] Vx = DelayTimer",
		Is:   func(o *Opcode) bool { return o.F == 0xF && o.NN == 0x07 },
		Execute: func(c *CPU, o *Opcode) {
			c.V[o.X] = byte(c.DelayTimer.GetValue())
		},
	},
	{
		Name: "[FX0A] Get Key",
		Is:   func(o *Opcode) bool { return o.F == 0xF && o.NN == 0x0A },
		Execute: func(c *CPU, o *Opcode) {
			// TODO: key
		},
	},
	{
		Name: "[FX15] DelayTimer = Vx",
		Is:   func(o *Opcode) bool { return o.F == 0xF && o.NN == 0x15 },
		Execute: func(c *CPU, o *Opcode) {
			c.DelayTimer.SetValue(int(c.V[o.X]))
		},
	},
	{
		Name: "[FX18] SoundTimer = Vx",
		Is:   func(o *Opcode) bool { return o.F == 0xF && o.NN == 0x18 },
		Execute: func(c *CPU, o *Opcode) {
			// TODO: sound
		},
	},
	{
		Name: "[FX1E] Add To Index",
		Is:   func(o *Opcode) bool { return o.F == 0xF && o.NN == 0x1E },
		Execute: func(c *CPU, o *Opcode) {
			if c.Config.InstructionOverflowAddIndex && int(c.I)+int(c.V[o.X]) > 0x0FFF {
				c.V[15] = 1
			}

			c.I += uint16(c.V[o.X])
		},
	},
	{
		Name: "[FX29] Set Index To Font Character",
		Is:   func(o *Opcode) bool { return o.F == 0xF && o.NN == 0x29 },
		Execute: func(c *CPU, o *Opcode) {
			// TODO: font
		},
	},
	{
		Name: "[FX33] Binary-coded Decimal Conversion",
		Is:   func(o *Opcode) bool { return o.F == 0xF && o.NN == 0x33 },
		Execute: func(c *CPU, o *Opcode) {
			VX := c.V[o.X]

			c.Memory.Write(c.I, []byte{
				byte(int(VX) / 100),
				byte((int(VX) % 100) / 10),
				byte(int(VX) % 10),
			})
		},
	},
	{
		Name: "[FX55] Store",
		Is:   func(o *Opcode) bool { return o.F == 0xF && o.NN == 0x55 },
		Execute: func(c *CPU, o *Opcode) {
			if c.Config.InstructionModifyIndexOnStoreAndLoad {
				for i := 0; i <= int(o.X); i++ {
					c.Memory.Write(c.I, []byte{c.V[i]})
					c.I++
				}
			} else {
				for i := 0; i <= int(o.X); i++ {
					c.Memory.Write(c.I+uint16(i), []byte{c.V[i]})
				}
			}
		},
	},
	{
		Name: "[FX65] Load",
		Is:   func(o *Opcode) bool { return o.F == 0xF && o.NN == 0x65 },
		Execute: func(c *CPU, o *Opcode) {
			if c.Config.InstructionModifyIndexOnStoreAndLoad {
				for i := 0; i <= int(o.X); i++ {
					c.V[i] = c.Memory.Read(c.I)
					c.I++
				}
			} else {
				for i := 0; i <= int(o.X); i++ {
					c.V[i] = c.Memory.Read(c.I + uint16(i))
				}
			}
		},
	},
}
