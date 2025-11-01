package emulator

import "math/rand"

func (e *Emulator) Op00E0() {
	e.Display.Clear()
}

func (e *Emulator) Op00EE() {
	e.PC = e.Stack.Pop()
}

func (e *Emulator) Op1NNN(NNN uint16) {
	e.PC = NNN
}

func (e *Emulator) Op2NNN(NNN uint16) {
	e.Stack.Push(e.PC)
	e.PC = NNN
}

func (e *Emulator) Op3XNN(X byte, NN byte) {
	if e.V[X] == NN {
		e.PC += 2
	}
}

func (e *Emulator) Op4XNN(X byte, NN byte) {
	if e.V[X] != NN {
		e.PC += 2
	}
}

func (e *Emulator) Op5XY0(X byte, Y byte) {
	if e.V[X] == e.V[Y] {
		e.PC += 2
	}
}

func (e *Emulator) Op6XNN(X byte, NN byte) {
	e.V[X] = NN
}

func (e *Emulator) Op7XNN(X byte, NN byte) {
	e.V[X] = e.V[X] + NN
}

func (e *Emulator) Op8XY0(X byte, Y byte) {
	e.V[X] = e.V[Y]
}

func (e *Emulator) Op8XY1(X byte, Y byte) {
	e.V[X] = e.V[X] | e.V[Y]
}

func (e *Emulator) Op8XY2(X byte, Y byte) {
	e.V[X] = e.V[X] & e.V[Y]
}

func (e *Emulator) Op8XY3(X byte, Y byte) {
	e.V[X] = e.V[X] ^ e.V[Y]
}

func (e *Emulator) Op8XY4(X byte, Y byte) {
	if int(e.V[X])+int(e.V[Y]) > 0xFFFF {
		e.V[15] = 1
	} else {
		e.V[15] = 0
	}

	e.V[X] = e.V[X] + e.V[Y]
}

func (e *Emulator) Op8XY5(X byte, Y byte) {
	if int(e.V[X])-int(e.V[Y]) > 0x0 {
		e.V[15] = 1
	} else {
		e.V[15] = 0
	}

	e.V[X] = e.V[X] - e.V[Y]
}

func (e *Emulator) Op8XY6(X byte, Y byte) {
	VX := e.V[X]

	if e.Legacy {
		VX = e.V[Y]
	}

	e.V[X] = VX >> 1
	e.V[15] = GetBitAtPosition(VX, 0)
}

func (e *Emulator) Op8XY7(X byte, Y byte) {
	if int(e.V[Y])-int(e.V[X]) > 0x0 {
		e.V[15] = 1
	} else {
		e.V[15] = 0
	}

	e.V[X] = e.V[Y] - e.V[X]
}

func (e *Emulator) Op8XYE(X byte, Y byte) {
	VX := e.V[X]

	if e.Legacy {
		VX = e.V[Y]
	}

	e.V[X] = VX << 1
	e.V[15] = GetBitAtPosition(VX, 7)
}

func (e *Emulator) Op9XY0(X byte, Y byte) {
	if e.V[X] != e.V[Y] {
		e.PC += 2
	}
}

func (e *Emulator) OpANNN(NNN uint16) {
	e.I = NNN
}

func (e *Emulator) OpBNNN(X byte, NNN uint16) {
	if e.Legacy {
		e.PC = uint16(e.V[0]) + NNN
	} else {
		e.PC = uint16(e.V[X]) + NNN
	}
}

func (e *Emulator) OpCXNN(X byte, NN byte) {
	e.V[X] = byte(rand.Intn(256)) & NN
}

func (e *Emulator) OpDXYN(X byte, Y byte, N byte) {
	x := int(e.V[X])
	y := int(e.V[Y])
	n := int(N)
	data := make([]byte, 0, n)

	for i := 0; i < n; i++ {
		data = append(data, e.Memory.Read(e.I+uint16(i)))
	}

	unset := e.Display.Write(x, y, data)

	if unset {
		e.V[15] = 1
	} else {
		e.V[15] = 0
	}
}

func (e *Emulator) OpEX9E(X byte) {
	// key
}

func (e *Emulator) OpEXA1(X byte) {
	// key
}

func (e *Emulator) OpFX07(X byte) {
	e.V[X] = byte(e.DelayTimer.GetValue())
}

func (e *Emulator) OpFX0A(X byte) {
	// key
}

func (e *Emulator) OpFX15(X byte) {
	e.DelayTimer.SetValue(int(e.V[X]))
}

func (e *Emulator) OpFX18(X byte) {
	// No sound implemented.
}

func (e *Emulator) OpFX1E(X byte) {
	if int(e.I)+int(e.V[X]) > 0x0FFF {
		e.V[15] = 1
	}

	e.I = e.I + uint16(e.V[X])
}

func (e *Emulator) OpFX29(X byte) {
	// font
}

func (e *Emulator) OpFX33(X byte) {
	VX := e.V[X]

	e.Memory.Write(e.I, []byte{
		byte(int(VX) / 100),
		byte((int(VX) % 100) / 10),
		byte(int(VX) % 10),
	})
}

func (e *Emulator) OpFX55(X byte) {
	if e.Legacy {
		for i := 0; i <= int(X); i++ {
			e.Memory.Write(e.I, []byte{e.V[i]})
			e.I++
		}
	} else {
		b := make([]byte, 0, X+1)

		for i := 0; i <= int(X); i++ {
			b = append(b, e.V[i])
		}

		e.Memory.Write(e.I, b)
	}
}

func (e *Emulator) OpFX65(X byte) {
	if e.Legacy {
		for i := 0; i <= int(X); i++ {
			e.V[i] = e.Memory.Read(e.I)
			e.I++
		}
	} else {
		for i := 0; i <= int(X); i++ {
			e.V[i] = e.Memory.Read(e.I + uint16(i))
		}
	}
}
