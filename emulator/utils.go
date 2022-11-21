package emulator

// GetBitAtPosition gets the bit at a specific position.
// 0 is the lowest bit, 7 is the highest.
func GetBitAtPosition(b byte, pos int) byte {
	return (b & (0x01 << pos) >> pos)
}
