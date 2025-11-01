package emulator

import "math/rand"

type Instruction struct {
	Name    string
	Is      func(o *Opcode) bool
	Execute func(e *Emulator, o *Opcode)
}

var Instructions = []Instruction{
	{
		Name: "[00E0] Clear Screen",
		Is:   func(o *Opcode) bool { return o.Raw == 0x00E0 },
		Execute: func(e *Emulator, o *Opcode) {
			// e.Display.Clear()
		},
	},
	{
		Name: "[00EE] Return Subroutine",
		Is:   func(o *Opcode) bool { return o.Raw == 0x00EE },
		Execute: func(e *Emulator, o *Opcode) {
			e.PC = e.Stack.Pop()
		},
	},
	{
		Name: "[1NNN] Jump",
		Is:   func(o *Opcode) bool { return o.F == 1 },
		Execute: func(e *Emulator, o *Opcode) {
			e.PC = o.NNN
		},
	},
	{
		Name: "[2NNN] Call Subroutine",
		Is:   func(o *Opcode) bool { return o.F == 2 },
		Execute: func(e *Emulator, o *Opcode) {
			e.Stack.Push(e.PC)
			e.PC = o.NNN
		},
	},
	{
		Name: "[3XNN] Skip If VX == NN",
		Is:   func(o *Opcode) bool { return o.F == 3 },
		Execute: func(e *Emulator, o *Opcode) {
			if e.V[o.X] == o.NN {
				e.PC += 2
			}
		},
	},
	{
		Name: "[4XNN] Skip If VX != NN",
		Is:   func(o *Opcode) bool { return o.F == 4 },
		Execute: func(e *Emulator, o *Opcode) {
			if e.V[o.X] != o.NN {
				e.PC += 2
			}
		},
	},
	{
		Name: "[5XY0] Skip If VX == VY",
		Is:   func(o *Opcode) bool { return o.F == 5 },
		Execute: func(e *Emulator, o *Opcode) {
			if e.V[o.X] == e.V[o.Y] {
				e.PC += 2
			}
		},
	},
	{
		Name: "[6XNN] VX == NN",
		Is:   func(o *Opcode) bool { return o.F == 6 },
		Execute: func(e *Emulator, o *Opcode) {
			e.V[o.X] = o.NN
		},
	},
	{
		Name: "[7XNN] Vx += NN (no carry)",
		Is:   func(o *Opcode) bool { return o.F == 7 },
		Execute: func(e *Emulator, o *Opcode) {
			e.V[o.X] = e.V[o.X] + o.NN
		},
	},
	{
		Name: "[8XY0] Vx = Vy",
		Is:   func(o *Opcode) bool { return o.F == 8 && o.N == 0 },
		Execute: func(e *Emulator, o *Opcode) {
			e.V[o.X] = e.V[o.Y]
		},
	},
	{
		Name: "[8XY1] Vx |= Vy",
		Is:   func(o *Opcode) bool { return o.F == 8 && o.N == 1 },
		Execute: func(e *Emulator, o *Opcode) {
			e.V[o.X] |= e.V[o.Y]
		},
	},
	{
		Name: "[8XY2] Vx &= Vy",
		Is:   func(o *Opcode) bool { return o.F == 8 && o.N == 2 },
		Execute: func(e *Emulator, o *Opcode) {
			e.V[o.X] &= e.V[o.Y]
		},
	},
	{
		Name: "[8XY3] Vx ^= Vy",
		Is:   func(o *Opcode) bool { return o.F == 8 && o.N == 3 },
		Execute: func(e *Emulator, o *Opcode) {
			e.V[o.X] ^= e.V[o.Y]
		},
	},
	{
		Name: "[8XY4] Vx += Vy",
		Is:   func(o *Opcode) bool { return o.F == 8 && o.N == 4 },
		Execute: func(e *Emulator, o *Opcode) {
			if int(e.V[o.X])+int(e.V[o.Y]) > 0xFFFF {
				e.V[15] = 1
			} else {
				e.V[15] = 0
			}

			e.V[o.X] += e.V[o.Y]
		},
	},
	{
		Name: "[8XY5] Vx -= Vy",
		Is:   func(o *Opcode) bool { return o.F == 8 && o.N == 5 },
		Execute: func(e *Emulator, o *Opcode) {
			if int(e.V[o.X])-int(e.V[o.Y]) > 0x0 {
				e.V[15] = 1
			} else {
				e.V[15] = 0
			}

			e.V[o.X] -= e.V[o.Y]
		},
	},
	{
		Name: "[8XY6] Vx >>= 1",
		Is:   func(o *Opcode) bool { return o.F == 8 && o.N == 6 },
		Execute: func(e *Emulator, o *Opcode) {
			VX := e.V[o.X]

			if e.Config.InstructionAssignBeforeShift {
				VX = e.V[o.Y]
			}

			e.V[o.X] = VX >> 1
			e.V[15] = GetBitAtPosition(VX, 0)
		},
	},
	{
		Name: "[8XY7] Vx = Vy - Vx",
		Is:   func(o *Opcode) bool { return o.F == 8 && o.N == 7 },
		Execute: func(e *Emulator, o *Opcode) {
			if int(e.V[o.Y])-int(e.V[o.X]) > 0x0 {
				e.V[15] = 1
			} else {
				e.V[15] = 0
			}

			e.V[o.X] = e.V[o.Y] - e.V[o.X]
		},
	},
	{
		Name: "[8XYE] Vx <<= 1",
		Is:   func(o *Opcode) bool { return o.F == 8 && o.N == 0xE },
		Execute: func(e *Emulator, o *Opcode) {
			VX := e.V[o.X]

			if e.Config.InstructionAssignBeforeShift {
				VX = e.V[o.Y]
			}

			e.V[o.X] = VX << 1
			e.V[15] = GetBitAtPosition(VX, 7)
		},
	},
	{
		Name: "[9XY0] Skip If Vx != Vy",
		Is:   func(o *Opcode) bool { return o.F == 9 },
		Execute: func(e *Emulator, o *Opcode) {
			if e.V[o.X] != e.V[o.Y] {
				e.PC += 2
			}
		},
	},
	{
		Name: "[ANNN] Set Index",
		Is:   func(o *Opcode) bool { return o.F == 0xA },
		Execute: func(e *Emulator, o *Opcode) {
			e.I = o.NNN
		},
	},
	{
		Name: "[BNNN] Jump With Offset",
		Is:   func(o *Opcode) bool { return o.F == 0xB },
		Execute: func(e *Emulator, o *Opcode) {
			if e.Config.InstructionUseVxForOffset {
				e.PC = uint16(e.V[0]) + o.NNN
			} else {
				e.PC = uint16(e.V[o.X]) + o.NNN
			}
		},
	},
	{
		Name: "[CNNN] Rand",
		Is:   func(o *Opcode) bool { return o.F == 0xC },
		Execute: func(e *Emulator, o *Opcode) {
			e.V[o.X] = byte(rand.Intn(256)) & o.NN
		},
	},
	{
		Name: "[DNNN] Display",
		Is:   func(o *Opcode) bool { return o.F == 0xC },
		Execute: func(e *Emulator, o *Opcode) {
			e.V[o.X] = byte(rand.Intn(256)) & o.NN
		},
	},
	{
		Name: "[DNNN] Display",
		Is:   func(o *Opcode) bool { return o.F == 0xD },
		Execute: func(e *Emulator, o *Opcode) {
			// x := int(e.V[o.X])
			// y := int(e.V[o.Y])
			n := int(o.N)
			data := make([]byte, 0, n)

			for i := 0; i < n; i++ {
				data = append(data, e.Memory.Read(e.I+uint16(i)))
			}

			// unset := e.Display.Write(x, y, data)

			// if unset {
			// 	e.V[15] = 1
			// } else {
			// 	e.V[15] = 0
			// }
		},
	},
	{
		Name: "[EX9E] Display",
		Is:   func(o *Opcode) bool { return o.F == 0xE && o.NN == 0x9E },
		Execute: func(e *Emulator, o *Opcode) {
			// TODO: key
		},
	},
	{
		Name: "[EX9E] Skip If Key Pressed",
		Is:   func(o *Opcode) bool { return o.F == 0xE && o.NN == 0x9E },
		Execute: func(e *Emulator, o *Opcode) {
			// TODO: key
		},
	},
	{
		Name: "[EXA1] Skip If Key Not Pressed",
		Is:   func(o *Opcode) bool { return o.F == 0xE && o.NN == 0xA1 },
		Execute: func(e *Emulator, o *Opcode) {
			// TODO: key
		},
	},
	{
		Name: "[FX07] Vx = DelayTimer",
		Is:   func(o *Opcode) bool { return o.F == 0xF && o.NN == 0x07 },
		Execute: func(e *Emulator, o *Opcode) {
			e.V[o.X] = byte(e.DelayTimer.GetValue())
		},
	},
	{
		Name: "[FX0A] Get Key",
		Is:   func(o *Opcode) bool { return o.F == 0xF && o.NN == 0x0A },
		Execute: func(e *Emulator, o *Opcode) {
			// TODO: key
		},
	},
	{
		Name: "[FX15] DelayTimer = Vx",
		Is:   func(o *Opcode) bool { return o.F == 0xF && o.NN == 0x15 },
		Execute: func(e *Emulator, o *Opcode) {
			e.DelayTimer.SetValue(int(e.V[o.X]))
		},
	},
	{
		Name: "[FX18] SoundTimer = Vx",
		Is:   func(o *Opcode) bool { return o.F == 0xF && o.NN == 0x18 },
		Execute: func(e *Emulator, o *Opcode) {
			// TODO: sound
		},
	},
	{
		Name: "[FX1E] Add To Index",
		Is:   func(o *Opcode) bool { return o.F == 0xF && o.NN == 0x1E },
		Execute: func(e *Emulator, o *Opcode) {
			if e.Config.InstructionOverflowAddIndex && int(e.I)+int(e.V[o.X]) > 0x0FFF {
				e.V[15] = 1
			}

			e.I += uint16(e.V[o.X])
		},
	},
	{
		Name: "[FX29] Set Index To Font Character",
		Is:   func(o *Opcode) bool { return o.F == 0xF && o.NN == 0x29 },
		Execute: func(e *Emulator, o *Opcode) {
			// TODO: font
		},
	},
	{
		Name: "[FX33] Binary-coded Decimal Conversion",
		Is:   func(o *Opcode) bool { return o.F == 0xF && o.NN == 0x33 },
		Execute: func(e *Emulator, o *Opcode) {
			VX := e.V[o.X]

			e.Memory.Write(e.I, []byte{
				byte(int(VX) / 100),
				byte((int(VX) % 100) / 10),
				byte(int(VX) % 10),
			})
		},
	},
	{
		Name: "[FX55] Store",
		Is:   func(o *Opcode) bool { return o.F == 0xF && o.NN == 0x55 },
		Execute: func(e *Emulator, o *Opcode) {
			if e.Config.InstructionModifyIndexOnStoreAndLoad {
				for i := 0; i <= int(o.X); i++ {
					e.Memory.Write(e.I, []byte{e.V[i]})
					e.I++
				}
			} else {
				for i := 0; i <= int(o.X); i++ {
					e.Memory.Write(e.I+uint16(i), []byte{e.V[i]})
				}
			}
		},
	},
	{
		Name: "[FX65] Load",
		Is:   func(o *Opcode) bool { return o.F == 0xF && o.NN == 0x65 },
		Execute: func(e *Emulator, o *Opcode) {
			if e.Config.InstructionModifyIndexOnStoreAndLoad {
				for i := 0; i <= int(o.X); i++ {
					e.V[i] = e.Memory.Read(e.I)
					e.I++
				}
			} else {
				for i := 0; i <= int(o.X); i++ {
					e.V[i] = e.Memory.Read(e.I + uint16(i))
				}
			}
		},
	},
}
