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
	if addr < 0 || addr >= m.Size {
		// TODO: log
		return
	}

	if (uint16(len(val)) + addr) > m.Size {
		// TODO: log
		return
	}

	for i, v := range val {
		m.Data[addr+uint16(i)] = v
	}
}

// Read bytes at a specific address.
func (m *Memory) Read(addr uint16) byte {
	if addr < 0 || addr >= m.Size {
		// TODO: log
		return 0
	}

	return m.Data[addr]
}
