package emulator

// Opcode is the parsed raw opcode.
type Opcode struct {
	// Raw is the raw opcode.
	Raw uint16
	// F is 0xF000.
	F byte
	// X is 0x0X00.
	X byte
	// Y is 0x00Y0.
	Y byte
	// N is 0x000N.
	N byte
	// NN is 0x00NN.
	NN byte
	// NNN is 0x0NNN.
	NNN uint16
}

func NewOpcode(raw uint16) *Opcode {
	return &Opcode{
		Raw: raw,
		F:   byte((raw & 0xF000) >> 0x0C),
		X:   byte((raw & 0x0F00) >> 0x08),
		Y:   byte((raw & 0x00F0) >> 0x04),
		N:   byte(raw & 0x000F),
		NN:  byte(raw & 0x00FF),
		NNN: uint16(raw & 0x0FFF),
	}
}
