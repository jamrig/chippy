package emulator

// Memory represents the RAM.
type Memory struct {
	// Size is the size of the memory in bytes.
	Size uint16
	// Data is the raw byte data.
	Data []byte
}

// NewMemory returns a new Memory.
func NewMemory(size uint16) *Memory {
	m := &Memory{
		Size: size,
		Data: make([]byte, size),
	}

	return m
}

// Write bytes starting at a specific address.
func (m *Memory) Write(addr uint16, val []byte) {
	for i, v := range val {
		loc := addr + uint16(i)
		if loc >= m.Size {
			return
		}

		m.Data[loc] = v
	}
}

// Read bytes at a specific address.
func (m *Memory) Read(addr uint16) byte {
	if addr >= m.Size {
		return 0
	}

	return m.Data[addr]
}
