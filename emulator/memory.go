package emulator

// Memory represents the RAM.
type Memory struct {
	// Size is the size of the memory in bytes.
	Size int
	// Data is the raw byte data.
	Data []byte
}

// NewMemory returns a new Memory.
func NewMemory(size int) *Memory {
	m := &Memory{
		Size: size,
		Data: make([]byte, size),
	}

	return m
}

// Write bytes starting at a specific address.
func (m *Memory) Write(addr int, val []byte) {
	if addr < 0 || addr >= m.Size {
		// TODO: log
		return
	}

	if (len(val) + addr) > m.Size {
		// TODO: log
		return
	}

	for i, v := range val {
		m.Data[addr+i] = v
	}
}
